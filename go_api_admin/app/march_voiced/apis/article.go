package apis

import (
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/app/march_voiced/service"
	"project/common/api"
	"project/utils"
	"project/utils/app"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var a = new(service.Article)

// InsertArticle 文章审核通过或驳回
// @Summary 文章审核通过或驳回
// @Description Author：Lbl 2021/02/18 获得身份令牌
// @Tags 文章：Article Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.InsertArticleDto false "添加参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSuccess
// @Router /api/article [post]
func ArticlePass() {

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

// ArticleDetail 文章详情
// @Summary 文章详情
// @Description Author：Lbl 2021/02/17 获得身份令牌
// @Tags 应用：文章管理 Article Controller
// @Accept application/json
// @Produce application/json
// @Param path false "修改参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._Article
// @Router /api/article/{id} [get]
func ArticleDetailHandler(c *gin.Context) {
	// 声明必要变量
	userMsg := new(api.UserMessage)
	article := new(bo.Article)
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
// @Success 200 {object} models._ResponseSuccess
// @Router /api/article/reprint [post]
func ReprintArticleHandler(c *gin.Context) {
	// 声明必要变量
	userMsg := new(api.UserMessage)
	var id []int
	var err error

	// 获取上下文用户信息
	userMsg, err = api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("ReprintArticleHandler Get userId failed", zap.Error(err))
		app.ResponseError(c, app.CodeNoUser)
		return
	}

	// 参数绑定
	err = c.ShouldBind(&id)
	if err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("ReprintArticleHandler params failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamIsInvalid)
		return
	}

	// 进入service层对数据操作
	err = a.ReprintArticle(id[0], userMsg.UserId)
	if err != nil {
		zap.L().Error("ReprintArticleHandler service failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeInsertOperationFail)
		return
	}

	// 成功返回状态
	app.ResponseSuccess(c, nil)
}

// TopArticleList 推荐文章列表
// @Summary 推荐文章列表
// @Description Author：Lbl 2021/02/17 获得身份令牌
// @Tags 应用：文章管理 Article Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.Paging false "添加参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._Article
// @Router /api/article/top [get]
func TopArticleListHandler(c *gin.Context) {
	// 声明必要变量
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
	paging.Current, err = utils.StringToInt(c.DefaultQuery("current", "1"))
	if err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("TopArticleListHandler params Current failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}
	paging.Size, err = utils.StringToInt(c.DefaultQuery("size", "10"))
	if err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("TopArticleListHandler params Size failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}

	// 进入service层对数据操作
	articleList, err = a.TopArticleList(paging, userMsg.UserId)
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
	res, err := a.MatchSensitiveWord(idI)
	if err != nil {
		zap.L().Error("match sensitiveWord failed", zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}
	// 返回响应
	app.ResponseSuccess(c, res)
}
