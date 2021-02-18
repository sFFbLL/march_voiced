package service

import (
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/utils"
	"sort"
	"sync"

	"go.uber.org/zap"
)

var wg sync.WaitGroup
var articleCh chan bo.Article

type Article struct{}

// 添加文章
func (a Article) InsertArticle(articleDto *dto.InsertArticleDto, userId int) (err error) {
	// 实例化要添加的文章
	article := new(models.Article)
	article.CreatBy = uint(userId)
	article.UpdateBy = uint(userId)
	article.Title = articleDto.Title
	article.Content = articleDto.Content
	article.Image = articleDto.Image
	article.Kind = articleDto.Kind
	article.Tag = articleDto.Tag
	article.Status = articleDto.Status
	article.Type = *articleDto.Type

	// 调用dao方法
	err = article.InsertArticle()
	return
}

// 更新文章
func (a Article) UpdateArticle(articleDto *dto.UpdateArticleDto, userId int) (err error) {
	// 实例化要修改的文章
	article := new(models.Article)
	article.UpdateBy = uint(userId)
	article.Title = articleDto.Title
	article.Content = articleDto.Content
	article.Image = articleDto.Image
	article.Kind = articleDto.Kind
	article.Tag = articleDto.Tag
	article.Status = articleDto.Status
	article.Type = articleDto.Type

	// 调用dao方法
	err = article.UpdateArticle(articleDto.ID)
	return
}

// 文章详情
func (a Article) ArticleDetail(id int, userId int) (articleMsg *bo.Article, err error) {
	//初始化
	articleMsg = new(bo.Article)
	article := new(models.Article)
	var userMsg bo.UserMsg
	var total [3]int64

	// 获取数据
	userMsg, total, err = article.ArticleDetail(id, userId)
	// 封装bo
	articleMsg.User = userMsg
	articleMsg.CollectTotal = total[0]
	articleMsg.CommentTotal = total[1]
	articleMsg.FavourTotal = total[2]
	articleMsg.ID = article.ID
	articleMsg.Title = article.Title
	articleMsg.Content = article.Content
	articleMsg.Image = article.Image
	articleMsg.Tag = article.Tag
	articleMsg.Kind = article.Kind
	articleMsg.Type = article.Type
	articleMsg.UpdateTime = article.UpdateTime
	articleMsg.CreatTime = article.CreateTime
	articleMsg.UpdateBy = article.UpdateBy
	articleMsg.CreatBy = article.CreatBy

	return
}

// 转载文章
func (a Article) ReprintArticle(id int, userId int) (err error) {
	// 实例化要添加的文章
	var status uint8 = 1
	article := new(models.Article)
	article.CreatBy = uint(userId)
	article.UpdateBy = uint(userId)
	article.Status = &status
	article.Type = uint(id)

	// 调用dao方法
	err = article.InsertArticle()
	return
}

// 推荐文章
func (a Article) TopArticleList(paging dto.Paging, userId int) (articleList *[]bo.Article, err error) {
	// 声明所需变量
	articleList = new([]bo.Article)
	article := new(models.Article)
	articleCh = make(chan bo.Article, 7)
	var keys []int

	// 取出要返回文章信息
	articleArray, err := article.TopArticleList(paging)
	if err != nil {
		zap.L().Error("TopArticleList Select ArticleList filed", zap.Error(err))
		close(articleCh)
		return
	}

	// 数据拼接
	for _, v := range *articleArray {
		wg.Add(1)
		go goArticleMsg(&articleCh, v.ID, userId)
	}
	wg.Wait()
	close(articleCh)

	//articleMap 排序 遍历
	articleMapList := make(map[int]bo.Article, len(*articleArray))
	for i := range articleCh {
		articleMapList[int(i.CreatTime)] = i
		keys = append(keys, int(i.CreatTime))
	}
	// 排序
	sort.Ints(keys)
	// 遍历
	for _, key := range keys {
		*articleList = append(*articleList, articleMapList[key])
	}

	return
}

func goArticleMsg(articleCh *chan bo.Article, articleId int, userId int) {
	articleMsg := new(bo.Article)
	article := new(models.Article)
	var userMsg bo.UserMsg
	var total [3]int64
	var err error

	// 获取数据
	userMsg, total, err = article.ArticleDetail(articleId, userId)
	if err != nil {
		zap.L().Error("goArticleMsg articleId："+utils.IntToString(articleId)+" failed", zap.Error(err))
		wg.Done()
		return
	}
	// 封装bo
	articleMsg.User = userMsg
	articleMsg.CollectTotal = total[0]
	articleMsg.CommentTotal = total[1]
	articleMsg.FavourTotal = total[2]
	articleMsg.ID = article.ID
	articleMsg.Title = article.Title
	articleMsg.Content = article.Content
	articleMsg.Image = article.Image
	articleMsg.Tag = article.Tag
	articleMsg.Kind = article.Kind
	articleMsg.Type = article.Type
	articleMsg.UpdateTime = article.UpdateTime
	articleMsg.CreatTime = article.CreateTime
	articleMsg.UpdateBy = article.UpdateBy
	articleMsg.CreatBy = article.CreatBy
	*articleCh <- *articleMsg

	wg.Done()
	return
}
