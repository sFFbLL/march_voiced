package apis

import (
	"project/app/march_voiced/models/dto"
	"project/app/march_voiced/service"
	"project/common/api"
	"project/utils/app"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

var co = new(service.ArticleComment)

// AddArticleComment	新增文章评论
// @Summary 新增文章评论
// @Description Author：YanSongWu 2021/02/17
// @Tags 评论：评论管理 Comment Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.AddArticleComment false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseArticleComment
// @Router /api/comment/article [post]
func AddArticleComment(c *gin.Context) {
	p := new(dto.AddArticleComment)
	// 上下文获取信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("GetUserMessage failed", zap.Error(err))
		return
	}
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("comment bind params failed", zap.String("username", user.Username))
		c.Error(err)
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	// 业务逻辑处理
	err = co.AddArticleComment(user.UserId, p)
	if err != nil {
		c.Error(err)
		zap.L().Error("add article comment failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeInsertOperationFail)
		return
	}

	// 返回响应
	app.ResponseSuccess(c, nil)
}

// DeleteArticleComment 删除文章详情页的评论
// @Summary 删除文章详情页的评论
// @Description Author：YanSongWu 2021/02/17
// @Tags 评论：详情页评论 comment Controller
// @Accept application/json
// @Produce application/json
// @Param object body int false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseArticleComment
// @Router /api/comment/article [delete]
func DeleteArticleComment(c *gin.Context) {
	p := new(dto.DeleteArticleComment)
	// 获取上下文信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		c.Error(err)
		zap.L().Error("DeleteArticleComment failed", zap.Error(err))
		return
	}
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("comment bind params failed", zap.String("username", user.Username), zap.Error(err))
		c.Error(err)
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseSuccess(c, app.CodeParamNotComplete)
		return
	}
	if err = co.DeleteArticleComment(user.UserId, p.ID); err != nil {
		c.Error(err)
		zap.L().Error("delete article comment failed", zap.String("Username", user.Username), zap.Error(err))
		return
	}
	app.ResponseSuccess(c, nil)
}

// GetArticleComment 查询文章详情页的评论
// @Summary 查询评论
// @Description Author：YanSongWu 2021/02/17
// @Tags 评论：评论管理 comment Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.GetArticleComment false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseArticleComment
// @Router /api/comment/article [get]
func GetArticleComment(c *gin.Context) {
	p := new(dto.GetArticleComment)
	// 获取上下文中信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		c.Error(err)
		zap.L().Error("GetArticleComment failed", zap.Error(err))
		return
	}
	if err := c.ShouldBindQuery(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("comment bind params failed", zap.String("Username", user.Username), zap.Error(err))
		c.Error(err)
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseSuccess(c, app.CodeParamNotComplete)
		return
	}
	// 业务逻辑处理
	res, err := co.GetArticleComment(p)
	if err != nil {
		c.Error(err)
		zap.L().Error("get article comment failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}
	// 返回响应
	app.ResponseSuccess(c, res)
}

// GetArticleChildComment 查询文章详情页一父评论的全部子评论
// @Summary 查询评论
// @Description Author：YanSongWu 2021/02/17
// @Tags 评论：评论管理 comment Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.GetArticleChildComment false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseArticleChildComment
// @Router /api/comment/children/article [get]
func GetArticleChildComment(c *gin.Context) {
	p := new(dto.GetArticleChildComment)
	// 获取上下文信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		c.Error(err)
		zap.L().Error("get article child comment failed", zap.Error(err))
		return
	}
	if err := c.ShouldBindQuery(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("comment bind params failed", zap.String("Username", user.Username), zap.Error(err))
		c.Error(err)
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 业务逻辑处理
	res, err := co.GetArticleChildComment(p)
	if err != nil {
		c.Error(err)
		zap.L().Error("get article child comment failed", zap.String("Username", user.Username), zap.Error(err))
		return
	}
	// 返回响应
	app.ResponseSuccess(c, res)
}
