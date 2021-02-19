package service

import (
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/utils/sensitiveWord"
	"sort"
	"sync"

	"go.uber.org/zap"
)

type Article struct{}

// 添加文章
func (a *Article) InsertArticle(articleDto *dto.InsertArticleDto, userId int) (err error) {
	// 实例化要添加的文章
	article := new(models.Article)
	article.CreateBy = uint(userId)
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
func (a *Article) UpdateArticle(articleDto *dto.UpdateArticleDto, userId int) (err error) {
	// 实例化要修改的文章
	article := new(models.Article)
	// 获取该文章原内容
	err = article.GetArticle()
	if err != nil {
		return
	}

	// 更新文章内容
	article.UpdateBy = uint(userId)
	article.Title = articleDto.Title
	article.Content = articleDto.Content
	article.Image = articleDto.Image
	article.Kind = articleDto.Kind
	article.Tag = articleDto.Tag
	article.Status = articleDto.Status
	article.Type = articleDto.Type
	article.ID = int(articleDto.ID)

	// 调用dao方法
	err = article.UpdateArticle()
	return
}

// 文章详情
func (a *Article) ArticleDetail(id int, userId int) (articleMsg *bo.Article, err error) {
	//初始化
	article := new(models.Article)
	articleCollect := new(models.ArticleCollect)
	articleComment := new(models.ArticleComment)
	articleFavour := new(models.ArticleFavour)
	follow := new(models.Follow)
	articleMsg = new(bo.Article)
	var userMsg bo.UserMsg

	// 获取数据
	article.ID = id
	userMsg, err = article.ArticleDetail()
	if err != nil {
		return
	}

	// 封装bo
	articleMsg.UserMsg = userMsg
	// 文章信息
	articleMsg.ID = article.ID
	articleMsg.Title = article.Title
	articleMsg.Content = article.Content
	articleMsg.Image = article.Image
	articleMsg.Tag = article.Tag
	articleMsg.Kind = article.Kind
	articleMsg.Type = article.Type
	articleMsg.UpdateTime = article.UpdateTime
	articleMsg.CreateTime = article.CreateTime
	articleMsg.UpdateBy = article.UpdateBy
	articleMsg.CreateBy = article.CreateBy
	// 文章收藏数
	articleCollect.ArticleId = uint(id)
	articleMsg.CollectTotal, err = articleCollect.ArticleCollectCount()
	if err != nil {
		zap.L().Error("ArticleCollectCount Get count failed", zap.Error(err))
		err = nil
	}
	// 文章评论数
	articleComment.ArticleId = uint(id)
	articleMsg.CommentTotal, err = articleComment.ArticleCommentCount()
	if err != nil {
		zap.L().Error("ArticleCommentCount Get count failed", zap.Error(err))
		err = nil
	}
	// 文章喜欢数
	articleFavour.ArticleId = uint(id)
	articleMsg.FavourTotal, err = articleFavour.ArticleFavourCount()
	if err != nil {
		zap.L().Error("ArticleFavourCount Get count failed", zap.Error(err))
		err = nil
	}
	// 是否关注
	if userId == userMsg.UserID {
		articleMsg.IsFollow = 0
	} else {
		articleMsg.IsFollow, err = follow.IsFollow(userId, userMsg.UserID)
		if err != nil {
			zap.L().Error("IsFollow failed", zap.Error(err))
			err = nil
		}
	}

	return
}

// 转载文章
func (a *Article) ReprintArticle(id int, userId int) (err error) {
	// 实例化要添加的文章
	var status uint8 = 1
	article := new(models.Article)
	article.CreateBy = uint(userId)
	article.UpdateBy = uint(userId)
	article.Status = &status
	article.Type = uint(id)

	// 调用dao方法
	err = article.InsertArticle()
	return
}

// 推荐文章
func (a *Article) TopArticleList(paging dto.Paging, userId int) (articleList *[]bo.Article, err error) {
	// 声明所需变量
	articleList = new([]bo.Article)
	article := new(models.Article)
	var keys []int

	// 取出要返回文章信息
	articleArray, err := article.TopArticleList(paging)
	if err != nil {
		zap.L().Error("TopArticleList Select ArticleList filed", zap.Error(err))
		return
	}

	// 数据拼接
	var wg sync.WaitGroup
	articleCh := make(chan *bo.Article, len(*articleArray))
	for _, v := range *articleArray {
		wg.Add(1)
		go goArticleMsg(&articleCh, &wg, userId, v)
	}
	wg.Wait()
	close(articleCh)

	//articleMap 排序 遍历
	articleMapList := make(map[int]*bo.Article, len(*articleArray))
	for i := range articleCh {
		articleMapList[int(i.CreateTime)] = i
		keys = append(keys, int(i.CreateTime))
	}
	// 排序
	sort.Ints(keys)
	// 遍历
	for _, key := range keys {
		*articleList = append(*articleList, *articleMapList[key])
	}
	return
}

func goArticleMsg(articleCh *chan *bo.Article, wg *sync.WaitGroup, userId int, articleMsg bo.Article) {
	articleCollect := new(models.ArticleCollect)
	articleComment := new(models.ArticleComment)
	articleFavour := new(models.ArticleFavour)
	follow := new(models.Follow)
	var err error

	// 获取数据
	articleCollect.ArticleId = uint(articleMsg.ID)
	articleMsg.CollectTotal, err = articleCollect.ArticleCollectCount()
	if err != nil {
		zap.L().Error("goArticleMsg ArticleCollectCount failed", zap.Error(err))
	}
	articleComment.ArticleId = uint(articleMsg.ID)
	articleMsg.CommentTotal, err = articleComment.ArticleCommentCount()
	if err != nil {
		zap.L().Error("goArticleMsg ArticleCommentCount failed", zap.Error(err))
	}
	articleFavour.ArticleId = uint(articleMsg.ID)
	articleMsg.FavourTotal, err = articleFavour.ArticleFavourCount()
	if err != nil {
		zap.L().Error("goArticleMsg ArticleFavourCount failed", zap.Error(err))
	}
	// 是否关注
	if userId == articleMsg.UserID {
		articleMsg.IsFollow = 0
	} else {
		articleMsg.IsFollow, err = follow.IsFollow(userId, articleMsg.UserID)
		if err != nil {
			zap.L().Error("goArticleMsg IsFollow failed", zap.Error(err))
			err = nil
		}
	}
	// 管道放数据
	*articleCh <- &articleMsg
	wg.Done()
	return
}

// MatchSensitiveWord  文章敏感词匹配
func (a *Article) MatchSensitiveWord(id int) (res *bo.WordRes, err error) {
	article := new(models.Article)
	// 获取文章
	content, err := article.GetArticleContent(id)
	if err != nil {
		return nil, err
	}
	// 匹配敏感词库
	data, err := sensitiveWord.Check(content)
	if err != nil {
		return nil, err
	}

	res = new(bo.WordRes)
	res.Sensitive = *data
	return
}
