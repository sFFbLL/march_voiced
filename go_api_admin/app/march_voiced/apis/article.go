package apis

import (
	"project/app/march_voiced/models/dto"
	"project/app/march_voiced/service"
	"project/common/api"
	"project/utils/app"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

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
	err = service.InsertArticle(article, userMsg.UserId)
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
	err = service.UpdateArticle(article, userMsg.UserId)
	if err != nil {
		zap.L().Error("UpdateArticleHandler service failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeUpdateOperationFail)
		return
	}

	// 成功返回状态
	app.ResponseSuccess(c, nil)
}
