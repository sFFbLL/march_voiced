package apis

import (
	"project/app/march_voiced/models/dto"
	"project/app/march_voiced/service"
	"project/common/api"
	"project/utils/app"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

var mc = new(service.MarchsoftComment)

// AddMarchsoftComment	新增文章评论
// @Summary 新增文章评论
// @Description Author：YanSongWu 2021/02/17
// @Tags 评论：评论管理 Comment Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.AddMarchsoftComment false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseMarchsoftComment
// @Router /api/comment/marchsoft [post]
func AddMarchsoftComment(c *gin.Context) {
	p := new(dto.AddMarchsoftComment)
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
	err = mc.AddMarchsoftComment(user.UserId, p)
	if err != nil {
		c.Error(err)
		zap.L().Error("add marchsoft comment failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeInsertOperationFail)
		return
	}

	// 返回响应
	app.ResponseSuccess(c, nil)
}

// DeleteMarchsoftComment 删除文章详情页的评论
// @Summary 删除文章详情页的评论
// @Description Author：YanSongWu 2021/02/17
// @Tags 评论：详情页评论 comment Controller
// @Accept application/json
// @Produce application/json
// @Param object body int false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseMarchsoftComment
// @Router /api/comment/marchsoft [delete]
func DeleteMarchsoftComment(c *gin.Context) {
	p := new(dto.DeleteMarchsoftComment)
	// 获取上下文信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		c.Error(err)
		zap.L().Error("DeleteMarchsoftComment failed", zap.Error(err))
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
	if err = mc.DeleteMarchsoftComment(user.UserId, p.ID); err != nil {
		c.Error(err)
		zap.L().Error("delete marchsoft comment failed", zap.String("Username", user.Username), zap.Error(err))
		return
	}
	app.ResponseSuccess(c, nil)
}

// GetMarchsoftComment 查询文章详情页的评论
// @Summary 查询评论
// @Description Author：YanSongWu 2021/02/17
// @Tags 评论：评论管理 comment Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.GetMarchsoftComment false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseMarchsoftComment
// @Router /api/comment/marchsoft [get]
func GetMarchsoftComment(c *gin.Context) {
	p := new(dto.GetMarchsoftComment)
	// 获取上下文中信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		c.Error(err)
		zap.L().Error("GetMarchsoftComment failed", zap.Error(err))
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
	res, err := mc.GetMarchsoftComment(p)
	if err != nil {
		c.Error(err)
		zap.L().Error("get marchsoft comment failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}
	// 返回响应
	app.ResponseSuccess(c, res)
}

// GetMarchsoftChildComment 查询文章详情页一父评论的全部子评论
// @Summary 查询评论
// @Description Author：YanSongWu 2021/02/17
// @Tags 评论：评论管理 comment Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.GetMarchsoftChildComment false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseMarchsoftChildComment
// @Router /api/comment/children/marchsoft [get]
func GetMarchsoftChildComment(c *gin.Context) {
	p := new(dto.GetMarchsoftChildComment)
	// 获取上下文信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		c.Error(err)
		zap.L().Error("get marchsoft child comment failed", zap.Error(err))
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
	res, err := mc.GetMarchsoftChildComment(p)
	if err != nil {
		c.Error(err)
		zap.L().Error("get marchsoft child comment failed", zap.String("Username", user.Username), zap.Error(err))
		return
	}
	// 返回响应
	app.ResponseSuccess(c, res)
}
