package service

import (
	"errors"
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/utils/sensitiveWord"
	"sort"
	"sync"

	"go.uber.org/zap"
)

type Article struct{}

// ArticlePass 文章审核
func (a *Article) ArticlePass(p *dto.ArticlePass, userId int) (err error) {
	article := new(models.Article)
	article.ID = int(p.Id)
	err = article.GetArticle()
	if err != nil {
		return
	}
	if *article.Status != 2 {
		return errors.New("文章不是发布未审核状态")
	}
	article.Status = p.Status
	err = article.UpdateArticle()
	go models.AddSysMessage(0, *p.Status, uint(userId), article.CreateBy)
	return
}

// 添加文章
func (a *Article) InsertArticle(articleDto *dto.InsertArticleDto, userId int) (err error) {
	// 实例化要添加的文章
	article := new(models.Article)
	article = &models.Article{
		CreateBy: uint(userId),
		UpdateBy: uint(userId),
		Title:    articleDto.Title,
		Content:  articleDto.Content,
		Image:    articleDto.Image,
		Kind:     articleDto.Kind,
		Tag:      articleDto.Tag,
		Status:   articleDto.Status,
		Type:     articleDto.Type,
	}

	// 调用dao方法
	err = article.InsertArticle()
	return
}

// 更新文章
func (a *Article) UpdateArticle(articleDto *dto.UpdateArticleDto, userId int) (err error) {
	// 实例化要修改的文章
	article := new(models.Article)
	// 获取该文章原内容
	article.ID = int(articleDto.ID)
	err = article.GetArticle()
	if err != nil {
		return
	}

	// 更新文章内容
	article = &models.Article{
		UpdateBy: uint(userId),
		Title:    articleDto.Title,
		Content:  articleDto.Content,
		Image:    articleDto.Image,
		Kind:     articleDto.Kind,
		Tag:      articleDto.Tag,
		Status:   articleDto.Status,
		Type:     articleDto.Type,
		BaseModel: models.BaseModel{
			ID: int(articleDto.ID),
		},
	}

	// 调用dao方法
	err = article.UpdateArticle()

	return
}

// 删除文章
func (a *Article) DeleteArticle(articleId int, userId int) (err error) {
	// 实例化要修改的文章
	article := new(models.Article)

	// 调用dao方法
	article.ID = articleId
	article.UpdateBy = uint(userId)
	article.IsDeleted = 1
	err = article.DeleteArticle()

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
	var userMsg bo.Article

	// 获取数据
	article.ID = id
	userMsg, err = article.ArticleDetail()
	if err != nil {
		return
	}
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
func (a *Article) ReprintArticle(id int, userId int) (articleMsg *bo.ArticleMsg, signal int, err error) {
	// 实例化要修改的文章
	article := new(models.Article)
	var t uint = 1
	var s uint8 = 1

	// 获取该文章原内容
	article.ID = id
	err = article.GetArticle()
	if err != nil {
		return
	}
	if *article.Status != 1 {
		signal = 1
		return
	}

	// 更新文章内容
	article.UpdateBy = uint(userId)
	article.CreateBy = uint(userId)
	article.UpdateTime = 0
	article.CreateTime = 0
	article.Type = &t
	article.Status = &s
	article.ID = 0

	// 调用dao方法
	err = article.InsertArticle()
	if err != nil {
		return
	}
	// 返回数据
	articleMsg = &bo.ArticleMsg{
		ID:         article.ID,
		Title:      article.Title,
		Content:    article.Content,
		Tag:        article.Tag,
		Kind:       article.Kind,
		WordCount:  article.WordCount,
		Type:       *article.Type,
		UpdateBy:   article.UpdateBy,
		CreateBy:   article.CreateBy,
		CreateTime: article.CreateTime,
		UpdateTime: article.UpdateTime,
	}

	return
}

// 推荐文章
func (a *Article) TopArticleList(paging dto.Paging, userId uint) (articleList *[]bo.Article, err error) {
	// 声明所需变量
	articleList = new([]bo.Article)
	article := new(models.Article)
	var goArticle bo.GoArticleMsg
	var keys []int

	// 取出要返回文章信息
	articleArray, err := article.ArticleList(paging, 1)
	if err != nil {
		zap.L().Error("TopArticleList Select ArticleList filed", zap.Error(err))
		return
	}

	// 组装进map
	articleMapList := make(map[int]*bo.Article, len(*articleArray))
	for _, i := range *articleArray {
		var articleBo bo.Article
		articleBo = i
		articleMapList[int(i.CreateTime)] = &articleBo
		keys = append(keys, int(i.CreateTime))
	}

	// 数据拼接
	var wg sync.WaitGroup
	articleCh := make(chan *bo.GoArticleMsg, len(*articleArray))
	for _, v := range *articleArray {
		wg.Add(1)
		goArticle.ArticleId = uint(v.ID)
		goArticle.ArticleUserId = v.CreateBy
		goArticle.UserId = userId
		goArticle.CreateTime = v.CreateTime
		go goArticleMsg(&articleCh, &wg, goArticle)
	}
	wg.Wait()
	close(articleCh)

	//articleMap 排序 遍历
	for i := range articleCh {
		goArticle = *i
		articleMapList[int(goArticle.CreateTime)].ArticleTotal = goArticle.ArticleTotal
	}
	// 排序
	sort.Ints(keys)
	// 遍历
	for _, key := range keys {
		*articleList = append(*articleList, *articleMapList[key])
	}
	return
}

// 主页文章
func (a *Article) SelectArticleListIndex(paging dto.Paging, userId uint) (articleList *[]bo.Article, err error) {
	// 声明所需变量
	articleList = new([]bo.Article)
	article := new(models.Article)
	var goArticle bo.GoArticleMsg
	var keys []int

	// 取出要返回文章信息
	articleArray, err := article.ArticleList(paging, 0)
	if err != nil {
		zap.L().Error("SelectArticleListIndex Select ArticleList filed", zap.Error(err))
		return
	}

	// 组装进map
	articleMapList := make(map[int]*bo.Article, len(*articleArray))
	for _, i := range *articleArray {
		var articleBo bo.Article
		articleBo = i
		articleMapList[int(i.CreateTime)] = &articleBo
		keys = append(keys, int(i.CreateTime))
	}

	// 数据拼接
	var wg sync.WaitGroup
	articleCh := make(chan *bo.GoArticleMsg, len(*articleArray))
	for _, v := range *articleArray {
		wg.Add(1)
		goArticle.ArticleId = uint(v.ID)
		goArticle.ArticleUserId = v.CreateBy
		goArticle.UserId = userId
		goArticle.CreateTime = v.CreateTime
		go goArticleMsg(&articleCh, &wg, goArticle)
	}
	wg.Wait()
	close(articleCh)

	//articleMap 排序 遍历
	for i := range articleCh {
		goArticle = *i
		articleMapList[int(goArticle.CreateTime)].ArticleTotal = goArticle.ArticleTotal
	}
	// 排序
	sort.Ints(keys)
	// 遍历
	for _, key := range keys {
		*articleList = append(*articleList, *articleMapList[key])
	}
	return
}

// 用户文章
func (a *Article) SelectArticleListByUserId(paging dto.SelectArticleByUser, userId uint) (articleList *[]bo.ArticleUser, err error) {
	// 声明所需变量
	articleList = new([]bo.ArticleUser)
	article := new(models.Article)
	var goArticle bo.GoArticleMsg
	var keys []int

	// 取出要返回文章信息
	articleArray, err := article.SelectArticleListByUserId(paging)
	if err != nil {
		zap.L().Error("SelectArticleListByUserId Select ArticleList filed", zap.Error(err))
		return
	}

	// 组装进map
	articleMapList := make(map[int]*bo.ArticleUser, len(*articleArray))
	for _, i := range *articleArray {
		articleBo := &bo.ArticleUser{
			ArticleMsg: bo.ArticleMsg{
				ID:         i.ID,
				Title:      i.Title,
				Content:    i.Content,
				Image:      i.Image,
				Tag:        i.Tag,
				Kind:       i.Kind,
				WordCount:  i.WordCount,
				Type:       *i.Type,
				UpdateBy:   i.UpdateBy,
				CreateBy:   i.CreateBy,
				CreateTime: i.CreateTime,
				UpdateTime: i.UpdateTime,
			},
		}

		articleMapList[int(i.CreateTime)] = articleBo
		keys = append(keys, int(i.CreateTime))
	}

	// 数据拼接
	var wg sync.WaitGroup
	articleCh := make(chan *bo.GoArticleMsg, len(*articleArray))
	for _, v := range *articleArray {
		wg.Add(1)
		goArticle.ArticleId = uint(v.ID)
		goArticle.ArticleUserId = v.CreateBy
		goArticle.UserId = userId
		goArticle.CreateTime = v.CreateTime
		go goArticleMsg(&articleCh, &wg, goArticle)
	}
	wg.Wait()
	close(articleCh)

	//articleMap 排序 遍历
	for i := range articleCh {
		goArticle = *i
		articleMapList[int(goArticle.CreateTime)].ArticleTotal = goArticle.ArticleTotal
	}
	// 排序
	sort.Ints(keys)
	// 遍历
	for _, key := range keys {
		*articleList = append(*articleList, *articleMapList[key])
	}
	return
}

func goArticleMsg(articleCh *chan *bo.GoArticleMsg, wg *sync.WaitGroup, articleMsg bo.GoArticleMsg) {
	articleCollect := new(models.ArticleCollect)
	articleComment := new(models.ArticleComment)
	articleFavour := new(models.ArticleFavour)
	follow := new(models.Follow)
	var err error

	// 获取数据
	articleCollect.ArticleId = articleMsg.ArticleId
	articleMsg.CollectTotal, err = articleCollect.ArticleCollectCount()
	if err != nil {
		zap.L().Error("goArticleMsg ArticleCollectCount failed", zap.Error(err))
	}
	articleComment.ArticleId = articleMsg.ArticleId
	articleMsg.CommentTotal, err = articleComment.ArticleCommentCount()
	if err != nil {
		zap.L().Error("goArticleMsg ArticleCommentCount failed", zap.Error(err))
	}
	articleFavour.ArticleId = articleMsg.ArticleId
	articleMsg.FavourTotal, err = articleFavour.ArticleFavourCount()
	if err != nil {
		zap.L().Error("goArticleMsg ArticleFavourCount failed", zap.Error(err))
	}
	// 是否关注
	if articleMsg.UserId == articleMsg.ArticleUserId {
		articleMsg.IsFollow = 0
	} else {
		articleMsg.IsFollow, err = follow.IsFollow(int(articleMsg.UserId), int(articleMsg.ArticleUserId))
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
	article.ID = id
	err = article.GetArticle()
	if err != nil {
		return nil, err
	}
	// 匹配敏感词库s
	data, err := sensitiveWord.Check(article.Content)
	if err != nil {
		return nil, err
	}

	res = new(bo.WordRes)
	res.Sensitive = *data
	return
}

// 文章收藏关注
func (a *Article) IsFavourCollectByArticleId(id int, userId int) (articleMsg *bo.IsFavourCollectByArticleId, err error) {
	//初始化
	articleCollect := new(models.ArticleCollect)
	articleFavour := new(models.ArticleFavour)
	articleMsg = new(bo.IsFavourCollectByArticleId)

	// 文章是收藏
	articleCollect.ArticleId = uint(id)
	articleCollect.CreateBy = uint(userId)
	articleMsg.IsCollect, err = articleCollect.IsCollectByArticleId()
	if err != nil {
		return
	}

	// 文章是喜欢
	articleFavour.ArticleId = uint(id)
	articleFavour.CreateBy = uint(userId)
	articleMsg.IsFavour, err = articleFavour.IsFavourByArticleId()

	return
}

// 文章标签
func (a *Article) GetArticleTagList() (articleTagList *[]bo.ArticleTagList, err error) {
	//初始化
	articleTag := new(models.ArticleTag)
	articleTagList = new([]bo.ArticleTagList)

	// 文章TagList
	articleTagList, err = articleTag.GetTagList()
	if err != nil {
		return
	}

	return
}
