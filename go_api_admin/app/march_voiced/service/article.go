package service

import (
	"errors"
	"sync"

	"project/app/march_voiced/models"
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/utils"
	"project/utils/sensitiveWord"

	"go.uber.org/zap"
)

type Article struct{}

func (a *Article) ArticleSearch(p *dto.ArticleSearchPaginator, userId int) (data *bo.ArticleCollectByUserId, err error) {
	article := new(models.Article)
	data = new(bo.ArticleCollectByUserId)
	records1 := new([]bo.Article)
	records2 := new([]bo.Article)
	var goArticle bo.GoArticleMsg
	var keys []int
	data.Records = records1
	err = article.ArticleSearchContent(data, p, userId)

	articleMapList := make(map[int]*bo.Article, len(*data.Records))
	for _, i := range *data.Records {
		var articleBo bo.Article
		articleBo = i
		articleMapList[i.ID] = &articleBo
		keys = append(keys, i.ID)
	}
	// 数据拼接
	var wg sync.WaitGroup
	articleCh := make(chan *bo.GoArticleMsg, len(*data.Records))
	for _, v := range *data.Records {
		wg.Add(1)
		goArticle.ArticleId = uint(v.ID)
		goArticle.ArticleUserId = v.CreateBy
		goArticle.UserId = uint(userId)
		go goArticleMsg(&articleCh, &wg, goArticle)
	}
	wg.Wait()
	close(articleCh)

	//articleMap 排序 遍历
	for i := range articleCh {
		goArticle = *i
		articleMapList[int(goArticle.ArticleId)].ArticleTotal = goArticle.ArticleTotal
	}

	for _, i := range *data.Records {
		*records2 = append(*records2, *articleMapList[i.ID])
	}
	data.Records = records2

	data.Current = p.Current
	data.Size = p.Size
	data.Pages = utils.PagesCount(int(data.Current), int(p.Size))
	return
}

// GetCollectArticle 我的收藏列表页
func (a *Article) GetCollectArticle(p *dto.Paginator, userId int) (data *bo.ArticleCollectByUserId, err error) {
	article := new(models.Article)
	data = new(bo.ArticleCollectByUserId)
	records1 := new([]bo.Article)
	records2 := new([]bo.Article)
	var goArticle bo.GoArticleMsg
	var keys []int
	data.Records = records1
	err = article.ArticleCollectByUserId(data, p, userId)

	articleMapList := make(map[int]*bo.Article, len(*data.Records))
	for _, i := range *data.Records {
		var articleBo bo.Article
		articleBo = i
		articleMapList[i.ID] = &articleBo
		keys = append(keys, i.ID)
	}
	// 数据拼接
	var wg sync.WaitGroup
	articleCh := make(chan *bo.GoArticleMsg, len(*data.Records))
	for _, v := range *data.Records {
		wg.Add(1)
		goArticle.ArticleId = uint(v.ID)
		goArticle.ArticleUserId = v.CreateBy
		goArticle.UserId = uint(userId)
		go goArticleMsg(&articleCh, &wg, goArticle)
	}
	wg.Wait()
	close(articleCh)

	//articleMap 排序 遍历
	for i := range articleCh {
		goArticle = *i
		articleMapList[int(goArticle.ArticleId)].ArticleTotal = goArticle.ArticleTotal
	}

	for _, i := range *data.Records {
		*records2 = append(*records2, *articleMapList[i.ID])
	}
	data.Records = records2

	data.Current = p.Current
	data.Size = p.Size
	data.Pages = utils.PagesCount(int(data.Current), int(p.Size))
	return
}

// ArticleRecommend （后台）文章设置推荐
func (a *Article) ArticleRecommend(articleId int) (err error) {
	article := new(models.Article)
	err = article.ArticleRecommend(articleId)
	if err != nil {
		return
	}
	if *article.IsRecommend == 0 {
		*article.IsRecommend = 1
	} else {
		*article.IsRecommend = 0
	}
	err = article.UpdateArticle()
	return
}

// ApplyArticleList （后台）文章审核列表页
func (a *Article) ApplyArticleList(p *dto.ApplyArticlePaginator, userId int) (applyArticleList *bo.ApplyArticleList, err error) {
	article := new(models.Article)
	applyArticleList = new(bo.ApplyArticleList)
	applyArticleListData := new([]bo.ApplyArticleListData)
	applyArticleList.Records = applyArticleListData
	err = article.GetApplyArticle(applyArticleList, p, userId)
	if err != nil {
		return
	}
	applyArticleList.Size = p.Size
	applyArticleList.Current = p.Current
	applyArticleList.Pages = utils.PagesCount(int(applyArticleList.Total), int(p.Size))
	return
}

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
	article.UpdateTime = utils.NowUnix()
	err = article.UpdateArticle()
	go models.AddSysMessage(0, *p.Status, uint(userId), article.CreateBy)
	if *p.Status == 1 {
		go models.AddMessage(0, 0, uint(article.ID), uint(userId), "")
	}
	return
}

// 添加文章
func (a *Article) InsertArticle(articleDto *dto.InsertArticleDto, userId int) (err error) {
	// 实例化要添加的文章
	article := &models.Article{
		CreateBy:  uint(userId),
		UpdateBy:  uint(userId),
		Title:     articleDto.Title,
		Content:   articleDto.Content,
		Describe:  articleDto.Describe,
		Image:     articleDto.Image,
		Kind:      articleDto.Kind,
		Tag:       articleDto.Tag,
		Status:    articleDto.Status,
		Type:      articleDto.Type,
		WordCount: articleDto.WordCount,
	}

	// 调用dao方法
	zap.L().Info("Call InsertArticle InsertArticle", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	err = article.InsertArticle()
	return
}

// 更新文章
func (a *Article) UpdateArticle(articleDto *dto.UpdateArticleDto, userId int) (err error) {
	// 实例化要修改的文章
	article := new(models.Article)
	var isRecommend uint8 = 0
	// 获取该文章原内容
	article.ID = int(articleDto.ID)
	zap.L().Info("Call UpdateArticle GetArticle", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	err = article.GetArticle()
	if err != nil {
		zap.L().Error("UpdateArticle GetArticle failed", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
		return
	}

	//判断是否有权利修改该文章
	if uint(userId) != article.CreateBy {
		err = errors.New("不能够修改他人文章！")
		zap.L().Error("UpdateArticle check CreateBy and userId failed", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
		return
	}

	// 更新文章内容
	article = &models.Article{
		UpdateBy:    uint(userId),
		Title:       articleDto.Title,
		Content:     articleDto.Content,
		Image:       articleDto.Image,
		IsRecommend: &isRecommend,
		Kind:        articleDto.Kind,
		Tag:         articleDto.Tag,
		Status:      articleDto.Status,
		Type:        articleDto.Type,
		WordCount:   articleDto.WordCount,
		BaseModel: models.BaseModel{
			ID: int(articleDto.ID),
		},
	}

	// 调用dao方法
	zap.L().Info("Call UpdateArticle UpdateArticle", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	err = article.UpdateArticle()

	return
}

// 删除文章
func (a *Article) DeleteArticle(articleId int, userId int) (err error) {
	// 实例化要修改的文章
	article := new(models.Article)
	// 获取该文章原内容
	article.ID = articleId
	zap.L().Info("Call DeleteArticle GetArticle", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	err = article.GetArticle()
	if err != nil {
		zap.L().Error("DeleteArticle GetArticle failed", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
		return
	}

	// 判断是否有权利删除该文章
	if uint(userId) != article.CreateBy {
		err = errors.New("不能够修改他人文章！")
		zap.L().Error("DeleteArticle check CreateBy and userId failed", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
		return
	}

	// 删除文章
	article.UpdateBy = uint(userId)
	article.IsDeleted = 1
	zap.L().Info("Call DeleteArticle DeleteArticle", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	err = article.DeleteArticle()

	return
}

// 文章详情
func (a *Article) ArticleDetail(id int, userId int) (articleMsg *bo.ArticleDetail, err error) {
	//初始化
	article := new(models.Article)
	articleCollect := new(models.ArticleCollect)
	articleComment := new(models.ArticleComment)
	articleFavour := new(models.ArticleFavour)
	follow := new(models.Follow)
	articleMsg = new(bo.ArticleDetail)

	// 获取数据
	article.ID = id
	zap.L().Info("Call ArticleDetail ArticleDetail", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	articleMsg, err = article.ArticleDetail()
	if err != nil {
		zap.L().Error("ArticleCollectCount ArticleDetail failed", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
		return
	}

	// 文章收藏数
	zap.L().Info("Call ArticleDetail ArticleCollectCount", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	articleCollect.ArticleId = uint(id)
	articleMsg.CollectTotal, err = articleCollect.ArticleCollectCount()
	if err != nil {
		zap.L().Error("ArticleCollectCount Get count failed", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	}

	// 文章评论数
	zap.L().Info("Call ArticleDetail ArticleCommentCount", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	articleComment.ArticleId = uint(id)
	articleMsg.CommentTotal, err = articleComment.ArticleCommentCount()
	if err != nil {
		zap.L().Error("ArticleCommentCount Get count failed", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	}

	// 文章喜欢数
	zap.L().Info("Call ArticleDetail ArticleFavourCount", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	articleFavour.ArticleId = uint(id)
	articleMsg.FavourTotal, err = articleFavour.ArticleFavourCount()
	if err != nil {
		zap.L().Error("ArticleFavourCount Get count failed", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	}

	// 是否关注
	if userId == int(articleMsg.CreateBy) {
		articleMsg.IsFollow = 3
	} else {
		zap.L().Info("Call ArticleDetail IsFollow", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
		articleMsg.IsFollow, err = follow.IsFollow(userId, int(articleMsg.CreateBy))
		if err != nil {
			zap.L().Error("ArticleDetail IsFollow failed", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
		}
	}

	// 文章是否收藏
	zap.L().Info("Call ArticleDetail IsCollectByArticleId", zap.String("ArticleId", utils.IntToString(id)), zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	articleCollect.ArticleId = uint(id)
	articleCollect.CreateBy = uint(userId)
	articleMsg.IsCollect, err = articleCollect.IsCollectByArticleId()
	if err != nil {
		zap.L().Error("ArticleDetail IsCollectByArticleId failed", zap.String("ArticleId", utils.IntToString(id)), zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
		return
	}

	// 文章是否喜欢
	zap.L().Info("Call ArticleDetail IsFavourByArticleId", zap.String("ArticleId", utils.IntToString(id)), zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	articleFavour.ArticleId = uint(id)
	articleFavour.CreateBy = uint(userId)
	articleMsg.IsFavour, err = articleFavour.IsFavourByArticleId()
	if err != nil {
		zap.L().Error("ArticleDetail IsFavourByArticleId failed", zap.String("ArticleId", utils.IntToString(id)), zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	}

	return
}

// 转载文章
func (a *Article) ReprintArticle(id int, userId int) (articleMsg *bo.ArticleDetailMsg, signal int, err error) {
	// 实例化要修改的文章
	article := new(models.Article)
	articleTag := new(models.ArticleTag)
	var t uint
	var s uint8 = 1
	var tag string

	// 获取该文章原内容
	zap.L().Info("Call ReprintArticle GetArticle", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	article.ID = id
	err = article.GetArticle()
	if err != nil {
		zap.L().Error("ReprintArticle GetArticle Failed", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
		return
	}

	// 判断文章是否发布
	if *article.Status != 1 {
		signal = 1
		return
	}

	// 获取文章tag
	zap.L().Info("Call ReprintArticle GetTagById", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	articleTag.ID = int(article.Tag)
	tag, err = articleTag.GetTagById()
	if err != nil {
		zap.L().Error("ReprintArticle GetTagById Failed ", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
		return
	}

	// 更新文章内容
	t = uint(article.ID)
	article.UpdateBy = uint(userId)
	article.CreateBy = uint(userId)
	article.UpdateTime = 0
	article.CreateTime = 0
	article.Type = &t
	article.Status = &s
	article.ID = 0

	// 调用dao方法
	zap.L().Info("Call ReprintArticle InsertArticle", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	err = article.InsertArticle()
	if err != nil {
		zap.L().Error("ReprintArticle InsertArticle failed", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
		return
	}
	// 返回数据
	articleMsg = &bo.ArticleDetailMsg{
		ID:         article.ID,
		Title:      article.Title,
		Content:    article.Content,
		Tag:        tag,
		Kind:       article.Kind,
		WordCount:  *article.WordCount,
		Status:     *article.Status,
		Type:       *article.Type,
		UpdateBy:   article.UpdateBy,
		CreateBy:   article.CreateBy,
		CreateTime: article.CreateTime,
		UpdateTime: article.UpdateTime,
	}

	zap.L().Info("Call ReprintArticle AddMessage", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	go models.AddMessage(0, 3, uint(id), uint(userId), "")
	return
}

// 推荐文章
func (a *Article) TopArticleList(paging dto.Paging, userId uint) (articleList *[]bo.Article, err error) {
	// 声明所需变量
	articleList = new([]bo.Article)
	article := new(models.Article)
	var goArticle bo.GoArticleMsg

	// 取出要返回文章信息
	zap.L().Info("Call TopArticleList ArticleList", zap.String("UserID", utils.UIntToString(userId)), zap.Error(err))
	articleArray, err := article.ArticleList(paging, 1)
	if err != nil {
		zap.L().Error("TopArticleList Select ArticleList filed", zap.String("UserID", utils.UIntToString(userId)), zap.Error(err))
		return
	}

	// 组装进map
	articleMapList := make(map[int]*bo.Article, len(*articleArray))
	for _, i := range *articleArray {
		var articleBo bo.Article
		articleBo = i
		articleMapList[i.ID] = &articleBo
	}

	// 数据拼接
	var wg sync.WaitGroup
	articleCh := make(chan *bo.GoArticleMsg, len(*articleArray))
	for _, v := range *articleArray {
		wg.Add(1)
		goArticle.ArticleId = uint(v.ID)
		goArticle.ArticleUserId = v.CreateBy
		goArticle.UserId = userId
		zap.L().Info("Call TopArticleList goArticleMsg", zap.String("ArticleId", utils.IntToString(v.ID)), zap.String("UserID", utils.UIntToString(userId)), zap.Error(err))
		go goArticleMsg(&articleCh, &wg, goArticle)
	}
	wg.Wait()
	close(articleCh)

	//articleMap 排序 遍历
	for i := range articleCh {
		goArticle = *i
		articleMapList[int(goArticle.ArticleId)].ArticleTotal = goArticle.ArticleTotal
	}

	// 遍历
	for _, v := range *articleArray {
		*articleList = append(*articleList, *articleMapList[v.ID])
	}
	return
}

// 主页文章
func (a *Article) SelectArticleListIndex(paging dto.Paging, userId uint) (articleList *[]bo.Article, err error) {
	// 声明所需变量
	articleList = new([]bo.Article)
	article := new(models.Article)
	var goArticle bo.GoArticleMsg

	// 取出要返回文章信息
	zap.L().Info("Call SelectArticleListIndex ArticleList", zap.String("UserID", utils.UIntToString(userId)), zap.Error(err))
	articleArray, err := article.ArticleList(paging, 0)
	if err != nil {
		zap.L().Error("SelectArticleListIndex Select ArticleList filed", zap.String("UserID", utils.UIntToString(userId)), zap.Error(err))
		return
	}

	// 组装进map
	articleMapList := make(map[int]*bo.Article, len(*articleArray))
	for _, i := range *articleArray {
		var articleBo bo.Article
		articleBo = i
		articleMapList[i.ID] = &articleBo
	}

	// 数据拼接
	var wg sync.WaitGroup
	articleCh := make(chan *bo.GoArticleMsg, len(*articleArray))
	for _, v := range *articleArray {
		wg.Add(1)
		goArticle.ArticleId = uint(v.ID)
		goArticle.ArticleUserId = v.CreateBy
		goArticle.UserId = userId
		zap.L().Info("Call TopArticleList goArticleMsg", zap.String("ArticleId", utils.IntToString(v.ID)), zap.String("UserID", utils.UIntToString(userId)), zap.Error(err))
		go goArticleMsg(&articleCh, &wg, goArticle)
	}
	wg.Wait()
	close(articleCh)

	//articleMap 排序 遍历
	for i := range articleCh {
		goArticle = *i
		articleMapList[int(goArticle.ArticleId)].ArticleTotal = goArticle.ArticleTotal
	}

	// 遍历
	for _, v := range *articleArray {
		*articleList = append(*articleList, *articleMapList[v.ID])
	}
	return
}

// 用户文章
func (a *Article) SelectArticleListByUserId(paging dto.SelectArticleByUser, userId uint) (articleList *[]bo.ArticleUser, err error) {
	// 声明所需变量
	articleList = new([]bo.ArticleUser)
	article := new(models.Article)
	var goArticle bo.GoArticleMsg

	// 取出要返回文章信息
	zap.L().Info("Call SelectArticleListByUserId SelectArticleListByUserId", zap.String("UserID", utils.UIntToString(userId)), zap.Error(err))
	articleArray, err := article.SelectArticleListByUserId(paging, userId)
	if err != nil {
		zap.L().Error("SelectArticleListByUserId Select ArticleList filed", zap.String("UserID", utils.UIntToString(userId)), zap.Error(err))
		return
	}
	// 当数据为空时直接返回
	if articleArray == nil {
		zap.L().Info("SelectArticleListByUserId Select ArticleList No Data", zap.String("UserID", utils.UIntToString(userId)), zap.Error(err))
		return
	}

	// 组装进map
	articleMapList := make(map[int]*bo.ArticleUser, len(*articleArray))
	for _, i := range *articleArray {
		articleBo := &bo.ArticleUser{
			ArticleMsg: i,
		}

		articleMapList[i.ID] = articleBo
	}

	// 数据拼接
	var wg sync.WaitGroup
	articleCh := make(chan *bo.GoArticleMsg, len(*articleArray))
	for _, v := range *articleArray {
		wg.Add(1)
		goArticle.ArticleId = uint(v.ID)
		goArticle.ArticleUserId = v.CreateBy
		goArticle.UserId = userId
		zap.L().Info("Call TopArticleList goArticleMsg", zap.String("ArticleId", utils.IntToString(v.ID)), zap.String("UserID", utils.UIntToString(userId)), zap.Error(err))
		go goArticleMsg(&articleCh, &wg, goArticle)
	}
	wg.Wait()
	close(articleCh)

	//articleMap 排序 遍历
	for i := range articleCh {
		goArticle = *i
		articleMapList[int(goArticle.ArticleId)].ArticleTotal = goArticle.ArticleTotal
	}

	// 遍历 排序
	for _, v := range *articleArray {
		*articleList = append(*articleList, *articleMapList[v.ID])
	}
	return
}

func goArticleMsg(articleCh *chan *bo.GoArticleMsg, wg *sync.WaitGroup, articleMsg bo.GoArticleMsg) {
	// 异常捕获
	defer func() {
		if err := recover(); err != nil {
			zap.L().Error("Call goArticleMsg defer recover", zap.String("ArticleId", utils.UIntToString(articleMsg.ArticleId)), zap.String("UserID", utils.UIntToString(articleMsg.UserId)), zap.String("error", string(utils.Stack())))
		}

		// 管道放数据
		*articleCh <- &articleMsg
		wg.Done()
		zap.L().Info("Call goArticleMsg Done", zap.String("ArticleId", utils.UIntToString(articleMsg.ArticleId)), zap.String("UserID", utils.UIntToString(articleMsg.UserId)))
	}()

	articleCollect := new(models.ArticleCollect)
	articleComment := new(models.ArticleComment)
	articleFavour := new(models.ArticleFavour)
	follow := new(models.Follow)
	var err error

	// 获取文章收藏数
	zap.L().Info("Call goArticleMsg ArticleCollectCount", zap.String("ArticleId", utils.UIntToString(articleMsg.ArticleId)), zap.String("UserID", utils.UIntToString(articleMsg.UserId)), zap.Error(err))
	articleCollect.ArticleId = articleMsg.ArticleId
	articleMsg.CollectTotal, err = articleCollect.ArticleCollectCount()
	if err != nil {
		zap.L().Error("goArticleMsg ArticleCollectCount failed", zap.String("ArticleId", utils.UIntToString(articleMsg.ArticleId)), zap.String("UserID", utils.UIntToString(articleMsg.UserId)), zap.Error(err))
	}

	// 获取文章评论数
	zap.L().Info("Call goArticleMsg ArticleCommentCount", zap.String("ArticleId", utils.UIntToString(articleMsg.ArticleId)), zap.String("UserID", utils.UIntToString(articleMsg.UserId)), zap.Error(err))
	articleComment.ArticleId = articleMsg.ArticleId
	articleMsg.CommentTotal, err = articleComment.ArticleCommentCount()
	if err != nil {
		zap.L().Error("goArticleMsg ArticleCommentCount failed", zap.String("ArticleId", utils.UIntToString(articleMsg.ArticleId)), zap.String("UserID", utils.UIntToString(articleMsg.UserId)), zap.Error(err))
	}

	// 获取文章喜欢数
	zap.L().Info("Call goArticleMsg ArticleFavourCount", zap.String("ArticleId", utils.UIntToString(articleMsg.ArticleId)), zap.String("UserID", utils.UIntToString(articleMsg.UserId)), zap.Error(err))
	articleFavour.ArticleId = articleMsg.ArticleId
	articleMsg.FavourTotal, err = articleFavour.ArticleFavourCount()
	if err != nil {
		zap.L().Error("goArticleMsg ArticleFavourCount failed", zap.String("ArticleId", utils.UIntToString(articleMsg.ArticleId)), zap.String("UserID", utils.UIntToString(articleMsg.UserId)), zap.Error(err))
	}

	// 是否关注
	if articleMsg.UserId == articleMsg.ArticleUserId {
		articleMsg.IsFollow = 3
	} else {
		zap.L().Info("Call goArticleMsg IsFollow", zap.String("ArticleId", utils.UIntToString(articleMsg.ArticleId)), zap.String("UserID", utils.UIntToString(articleMsg.UserId)), zap.Error(err))
		articleMsg.IsFollow, err = follow.IsFollow(int(articleMsg.UserId), int(articleMsg.ArticleUserId))
		if err != nil {
			zap.L().Error("goArticleMsg IsFollow failed", zap.String("ArticleId", utils.UIntToString(articleMsg.ArticleId)), zap.String("UserID", utils.UIntToString(articleMsg.UserId)), zap.Error(err))
			err = nil
		}
	}

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

	// 文章是否收藏
	zap.L().Info("Call IsFavourCollectByArticleId IsCollectByArticleId", zap.String("ArticleId", utils.IntToString(id)), zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	articleCollect.ArticleId = uint(id)
	articleCollect.CreateBy = uint(userId)
	articleMsg.IsCollect, err = articleCollect.IsCollectByArticleId()
	if err != nil {
		zap.L().Error("goArticleMsg IsFollow failed", zap.String("ArticleId", utils.IntToString(id)), zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
		return
	}

	// 文章是否喜欢
	zap.L().Info("Call IsFavourCollectByArticleId IsFavourByArticleId", zap.String("ArticleId", utils.IntToString(id)), zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	articleFavour.ArticleId = uint(id)
	articleFavour.CreateBy = uint(userId)
	articleMsg.IsFavour, err = articleFavour.IsFavourByArticleId()

	return
}

// 文章标签
func (a *Article) GetArticleTagList(userId int) (articleTagList *[]bo.ArticleTagList, err error) {
	//初始化
	articleTag := new(models.ArticleTag)
	articleTagList = new([]bo.ArticleTagList)

	// 文章TagList
	zap.L().Info("Call GetArticleTagList GetTagList", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	articleTagList, err = articleTag.GetTagList()

	return
}
