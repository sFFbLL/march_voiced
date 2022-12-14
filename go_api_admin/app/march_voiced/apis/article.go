package apis

import (
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/app/march_voiced/service"
	"project/common/api"
	"project/utils"
	"project/utils/app"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// ArticleSearch 文章搜索
// @Summary 文章搜索
// @Description Author：JiaKun Li 2021/02/20
// @Tags 文章：Article Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.ArticleSearchPaginator false "添加参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseGetCollectArticle
// @Router /api/base/searchArticle [get]
func ArticleSearch(c *gin.Context) {
	p := new(dto.ArticleSearchPaginator)

	// 获取缓存信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("ArticleSearch GetUserMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeLoginExpire)
		return
	}

	// 获取参数 校验参数
	if err := c.ShouldBindQuery(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("ArticleSearch params failed", zap.String("Username", user.Username), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}

	//业务逻辑处理
	s := new(service.Article)
	data, err := s.ArticleSearch(p, user.UserId)
	if err != nil {
		zap.L().Error("ArticleSearch service params failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	app.ResponseSuccess(c, data)
}

// ArticleRecommend （后台）文章设置推荐
// @Summary （后台）文章设置推荐
// @Description Author：JiaKun Li 2021/02/20
// @Tags 文章：Article Controller
// @Accept application/json
// @Produce application/json
// @Param id path int false "添加参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSuccess
// @Router /api/article/recommend/:id [put]
func ArticleRecommend(c *gin.Context) {
	id := c.Param("id")

	// 获取缓存信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("ArticleRecommend GetUserMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeLoginExpire)
		return
	}

	articleId, err := strconv.Atoi(id)
	if err != nil {
		zap.L().Error("ArticleRecommend params failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamTypeBindError)
	}

	if articleId <= 0 {
		app.ResponseError(c, app.CodeParamIsInvalid)
	}

	//业务逻辑处理
	s := new(service.Article)
	err = s.ArticleRecommend(articleId)
	if err == gorm.ErrRecordNotFound {
		zap.L().Error("ArticleRecommend gorm.ErrRecordNotFound", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseErrorWithMsg(c, app.CodeSelectOperationFail, "该文章未发布")
	}
	if err != nil {
		zap.L().Error("ArticleRecommend service params failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	app.ResponseSuccess(c, nil)
}

// ApplyArticleList （后台）文章审核列表页
// @Summary （后台）文章审核列表页
// @Description Author：JiaKun Li 2021/02/20
// @Tags 文章：Article Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.ApplyArticlePaginator false "添加参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseApplyArticleList
// @Router /api/article/admin [get]
func ApplyArticleList(c *gin.Context) {
	p := new(dto.ApplyArticlePaginator)

	// 获取缓存信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("ApplyArticleList GetUserMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeLoginExpire)
		return
	}

	// 获取参数 校验参数
	if err := c.ShouldBindQuery(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("ApplyArticleList params failed", zap.String("Username", user.Username), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}

	if p.Size == 0 && p.Current == 0 {
		p.Size = 10
		p.Current = 1
	}

	//业务逻辑处理
	s := new(service.Article)
	applyArticleList, err := s.ApplyArticleList(p, user.UserId)
	if err != nil {
		zap.L().Error("ApplyArticleList service params failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	app.ResponseSuccess(c, applyArticleList)
}

// ArticlePass 文章审核通过或驳回
// @Summary 文章审核通过或驳回
// @Description Author：JiaKun Li 2021/02/18
// @Tags 文章：Article Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.ArticlePass false "添加参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSuccess
// @Router /api/apply/article [put]
func ArticlePass(c *gin.Context) {
	p := new(dto.ArticlePass)

	// 获取缓存信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("ArticlePass GetUserMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeLoginExpire)
		return
	}

	// 获取参数 校验参数
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("ArticlePass params failed", zap.String("Username", user.Username), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}

	//业务逻辑处理
	s := new(service.Article)
	err = s.ArticlePass(p, user.UserId)
	if err != nil {
		if err.Error() == "文章不是发布未审核状态" {
			zap.L().Error("ArticlePass service params failed", zap.String("Username", user.Username), zap.Error(err))
			app.ResponseError(c, app.CodeArticleIsNotApply)
			return
		}
		zap.L().Error("ArticlePass service params failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeUpdateOperationFail)
		return
	}

	app.ResponseSuccess(c, nil)
}

// InsertArticle 添加文章
// @Summary 添加文章
// @Description Author：Lbl 2021/02/17 获得身份令牌
// @Tags 应用：文章管理 Article Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.InsertArticleDto false "添加参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSuccess
// @Router /api/article [post]
func InsertArticleHandler(c *gin.Context) {
	// 声明必要变量
	a := new(service.Article)
	article := new(dto.InsertArticleDto)
	userMsg := new(api.UserMessage)
	var err error

	// 获取上下文用户信息
	userMsg, err = api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("InsertArticleHandler Get userId failed", zap.Error(err))
		app.ResponseError(c, app.CodeNoUser)
		return
	}

	// 参数绑定
	err = c.ShouldBindJSON(article)
	if err != nil {
		zap.L().Error("InsertArticleHandler ShouldBindJSON Params Failed", zap.String("Username", userMsg.Username), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 判断参数是否合法
	if *article.Status != 0 && *article.Status != 2 {
		zap.L().Error("InsertArticleHandler article.Status Params Failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamIsInvalid)
		return
	}
	// 进入service层对数据操作
	err = a.InsertArticle(article, userMsg.UserId)
	if err != nil {
		zap.L().Error("InsertArticleHandler service failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeInsertOperationFail)
		return
	}

	// 成功返回状态
	app.ResponseSuccess(c, nil)
}

// UpdateArticle 修改文章
// @Summary 修改文章
// @Description Author：Lbl 2021/02/17 获得身份令牌
// @Tags 应用：文章管理 Article Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.UpdateArticleDto false "修改参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSuccess
// @Router /api/article [put]
func UpdateArticleHandler(c *gin.Context) {
	// 声明必要变量
	a := new(service.Article)
	article := new(dto.UpdateArticleDto)
	userMsg := new(api.UserMessage)
	var err error

	// 获取上下文中的用户信息
	userMsg, err = api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("UpdateArticleHandler Get userId failed", zap.Error(err))
		app.ResponseError(c, app.CodeNoUser)
		return
	}

	// 参数绑定
	err = c.ShouldBindJSON(article)
	if err != nil {
		zap.L().Error("UpdateArticleHandler ShouldBindJSON Params Failed", zap.String("Username", userMsg.Username), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 进入service层对数据操作
	err = a.UpdateArticle(article, userMsg.UserId)
	if err != nil {
		zap.L().Error("UpdateArticleHandler service failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeUpdateOperationFail)
		return
	}

	// 成功返回状态
	app.ResponseSuccess(c, nil)
}

// DeleteArticle 删除文章
// @Summary 删除文章
// @Description Author：Lbl 2021/02/17 获得身份令牌
// @Tags 应用：文章管理 Article Controller
// @Accept application/json
// @Produce application/json
// @Param id path int false "修改参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSuccess
// @Router /api/article/{id} [delete]
func DeleteArticleHandler(c *gin.Context) {
	// 声明必要变量
	a := new(service.Article)
	userMsg := new(api.UserMessage)
	var err error
	var id int

	// 获取上下文中的用户信息
	userMsg, err = api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("DeleteArticleHandler Get userId failed", zap.Error(err))
		app.ResponseError(c, app.CodeNoUser)
		return
	}

	// 参数绑定

	// 获取参数
	id, err = utils.StringToInt(c.Param("id"))
	if err != nil {
		zap.L().Error("ArticleDetailHandler bind Query Failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 进入service层对数据操作
	err = a.DeleteArticle(id, userMsg.UserId)
	if err != nil {
		zap.L().Error("DeleteArticleHandle service failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeDeleteOperationFail)
		return
	}

	// 成功返回状态
	app.ResponseSuccess(c, nil)
}

// ArticleDetail 文章详情
// @Summary 文章详情
// @Description Author：Lbl 2021/02/17 获得身份令牌
// @Tags 应用：文章管理 Article Controller
// @Accept application/json
// @Produce application/json
// @Param id path int false "修改参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseArticleDetail
// @Router /api/article/{id} [get]
func ArticleDetailHandler(c *gin.Context) {
	// 声明必要变量
	a := new(service.Article)
	userMsg := new(api.UserMessage)
	article := new(bo.ArticleDetail)
	var id int
	var err error

	// 获取上下文中的用户信息
	userMsg, err = api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("ArticleDetailHandler Get userId failed", zap.Error(err))
		app.ResponseError(c, app.CodeNoUser)
		return
	}

	// 获取参数
	id, err = utils.StringToInt(c.Param("id"))
	if err != nil {
		zap.L().Error("ArticleDetailHandler bind Query Failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 进入service层对数据操作
	article, err = a.ArticleDetail(id, userMsg.UserId)
	if err != nil {
		zap.L().Error("ArticleDetailHandler service failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	// 成功返回状态
	app.ResponseSuccess(c, article)
}

// ReprintArticle 文章转载
// @Summary 文章转载
// @Description Author：Lbl 2021/02/17 获得身份令牌
// @Tags 应用：文章管理 Article Controller
// @Accept application/json
// @Produce application/json
// @Param object body []int false "添加参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseArticleReprint
// @Router /api/article/reprint [post]
func ReprintArticleHandler(c *gin.Context) {
	// 声明必要变量
	a := new(service.Article)
	userMsg := new(api.UserMessage)
	articleMsg := new(bo.ArticleDetailMsg)
	artRep := new(dto.ArticleReprint)
	var signal int
	var err error

	// 获取上下文用户信息
	userMsg, err = api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("ReprintArticleHandler Get userId failed", zap.Error(err))
		app.ResponseError(c, app.CodeNoUser)
		return
	}

	// 参数绑定
	err = c.ShouldBindJSON(artRep)
	if err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("ReprintArticleHandler params failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamIsInvalid)
		return
	}

	// 进入service层对数据操作
	articleMsg, signal, err = a.ReprintArticle(int(artRep.ID), userMsg.UserId)
	if err != nil {
		zap.L().Error("ReprintArticleHandler service failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeInsertOperationFail)
		return
	}
	if signal == 1 {
		zap.L().Error("ReprintArticleHandler the article dont's public ", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseErrorWithMsg(c, app.CodeOperationFail, "该文章没有审核通过，不能被转载")
		return
	}

	// 成功返回状态
	app.ResponseSuccess(c, articleMsg)
}

// TopArticleList 推荐文章列表
// @Summary 推荐文章列表
// @Description Author：Lbl 2021/02/17 获得身份令牌
// @Tags 应用：文章管理 Article Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.Paging false "添加参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseArticleList
// @Router /api/article/top [get]
func TopArticleListHandler(c *gin.Context) {
	// 声明必要变量
	a := new(service.Article)
	userMsg := new(api.UserMessage)
	articleList := new([]bo.Article)
	var paging dto.Paging
	var err error

	// 获取上下文用户信息
	userMsg, err = api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("TopArticleListHandler Get userId failed", zap.Error(err))
		app.ResponseError(c, app.CodeNoUser)
		return
	}

	// 参数绑定
	paging.Current, err = utils.StringToUint(c.DefaultQuery("current", "1"))
	if err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("TopArticleListHandler params Current failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}
	paging.Size, err = utils.StringToUint(c.DefaultQuery("size", "10"))
	if err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("TopArticleListHandler params Size failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}

	// 进入service层对数据操作
	articleList, err = a.TopArticleList(paging, uint(userMsg.UserId))
	if err != nil {
		zap.L().Error("TopArticleListHandler service failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	// 成功返回状态
	app.ResponseSuccess(c, articleList)
}

// MatchSensitiveWord 敏感词匹配
// @Summary 敏感词匹配
// @Description Author：YanSongWu 2021/02/18 获得身份令牌
// @Tags 应用：文章管理 Article Controller
// @Accept application/json
// @Produce application/json
// @Param object query int false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseMatchSensitiveWord
// @Router /api/article/word [get]
func MatchSensitiveWord(c *gin.Context) {
	// 获取上下文信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("GetUserMessage failed", zap.Error(err))
		return
	}

	// 接收参数
	id := c.Query("id")
	idI, err := strconv.Atoi(id)
	if err != nil {
		zap.L().Error("MatchSensitiveWord failed", zap.String("username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamIsInvalid)
		return
	}

	a := new(service.Article)
	res, err := a.MatchSensitiveWord(idI)
	if err != nil {
		zap.L().Error("match sensitiveWord failed", zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}
	// 返回响应
	app.ResponseSuccess(c, res)
}

// SelectArticleListIndex 文章列表
// @Summary 文章列表
// @Description Author：Lbl 2021/02/17 获得身份令牌
// @Tags 应用：文章管理 Article Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.Paging false "添加参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseArticleList
// @Router /api/article/index [get]
func SelectArticleListIndex(c *gin.Context) {
	// 声明必要变量
	a := new(service.Article)
	userMsg := new(api.UserMessage)
	articleList := new([]bo.Article)
	var paging dto.Paging
	var err error

	// 获取上下文用户信息
	userMsg, err = api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("TopArticleListHandler Get userId failed", zap.Error(err))
		app.ResponseError(c, app.CodeNoUser)
		return
	}

	// 参数绑定
	paging.Current, err = utils.StringToUint(c.DefaultQuery("current", "1"))
	if err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("TopArticleListHandler params Current failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}
	paging.Size, err = utils.StringToUint(c.DefaultQuery("size", "10"))
	if err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("TopArticleListHandler params Size failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}

	// 进入service层对数据操作
	articleList, err = a.SelectArticleListIndex(paging, uint(userMsg.UserId))
	if err != nil {
		zap.L().Error("TopArticleListHandler service failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	// 成功返回状态
	app.ResponseSuccess(c, articleList)
}

// SelectArticleListByUserId 用户文章列表
// @Summary 用户文章列表
// @Description Author：Lbl 2021/02/17 获得身份令牌
// @Tags 应用：文章管理 Article Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.SelectArticleByUser false "添加参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseArticleListByUserId
// @Router /api/article/user [get]
func SelectArticleListByUserId(c *gin.Context) {
	// 声明必要变量
	a := new(service.Article)
	userMsg := new(api.UserMessage)
	articleList := new([]bo.ArticleUser)
	var paging dto.SelectArticleByUser
	var err error

	// 获取上下文用户信息
	userMsg, err = api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("SelectArticleListByUserId Get userMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeNoUser)
		return
	}

	// 参数绑定
	err = c.ShouldBindQuery(&paging)
	if err != nil {
		zap.L().Error("SelectArticleListByUserId ShouldBindQuery Params Failed", zap.String("Username", userMsg.Username), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 参数矫正
	if paging.Current == 0 && paging.Size == 0 {
		paging.Current = 1
		paging.Current = 10
	}
	if paging.Current == 0 {
		paging.Current = 1
	}

	// 查看权限
	if paging.ID == 0 {
		paging.ID = uint(userMsg.UserId)
	}

	// 进入service层对数据操作
	articleList, err = a.SelectArticleListByUserId(paging, uint(userMsg.UserId))
	if err != nil {
		zap.L().Error("SelectArticleListByUserId service failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	// 成功返回状态
	app.ResponseSuccess(c, articleList)
}

// IsFavourCollectByArticleId 文章是否点赞收藏
// @Summary 文章是否点赞收藏
//@Description Author：Lbl 2021/02/17 获得身份令牌
// @Tags 应用：文章管理 Article Controller
// @Accept application/json
// @Produce application/json
// @Param id path int false "修改参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseIsFavourCollectByArticleId
// @Router /api/article/favour-collect/{id} [get]
func IsFavourCollectByArticleId(c *gin.Context) {
	// 声明必要变量
	a := new(service.Article)
	userMsg := new(api.UserMessage)
	article := new(bo.IsFavourCollectByArticleId)
	var id int
	var err error

	// 获取上下文中的用户信息
	userMsg, err = api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("IsFavourCollectByArticleId Get userMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeNoUser)
		return
	}

	// 获取参数
	id, err = utils.StringToInt(c.Param("id"))
	if err != nil {
		zap.L().Error("IsFavourCollectByArticleId bind Query Failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 进入service层对数据操作
	article, err = a.IsFavourCollectByArticleId(id, userMsg.UserId)
	if err != nil {
		zap.L().Error("IsFavourCollectByArticleId service failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	// 成功返回状态
	app.ResponseSuccess(c, article)
}

// GetArticleTagList 文章标签
// @Summary 文章标签
//@Description Author：Lbl 2021/02/17 获得身份令牌
// @Tags 应用：文章管理 Article Controller
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseArticleTagList
// @Router /api/article/tag [get]
func GetArticleTagList(c *gin.Context) {
	// 声明必要变量
	a := new(service.Article)
	userMsg := new(api.UserMessage)
	articleTagList := new([]bo.ArticleTagList)
	var err error

	// 获取上下文中的用户信息
	userMsg, err = api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("GetArticleTagList Get userMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeNoUser)
		return
	}

	// 进入service层对数据操作
	articleTagList, err = a.GetArticleTagList(userMsg.UserId)
	if err != nil {
		zap.L().Error("GetArticleTagList service failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	// 成功返回状态
	app.ResponseSuccess(c, articleTagList)
}
