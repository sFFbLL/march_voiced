package apis

import (
	"project/app/march_voiced/models/dto"
	"project/app/march_voiced/service"
	"project/common/api"
	"project/utils/app"
	"strconv"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

var fo = new(service.Follow)

// GetFollowList 查询关注列表
// @Summary 查询关注列表
// @Description Author：YanSongWu 2021/02/18
// @Tags 用户：关注管理 Follow Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.GetFollowList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseGetFollowList
// @Router /api/follow [get]
func GetFollowList(c *gin.Context) {
	p := new(dto.GetFollowList)

	// 获取上下文消息
	user, err := api.GetUserMessage(c)
	if err != nil {
		c.Error(err)
		zap.L().Error("GetUserMessage failed", zap.Error(err))
		return
	}

	if err := c.ShouldBindQuery(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("follow bind params failed", zap.String("Username", user.Username), zap.Error(err))
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
	if p.Id == 0 {
		p.Id = user.UserId
	}
	res, err := fo.GetFollowList(p)
	if err != nil {
		zap.L().Error("get follow list failed", zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}
	// 返回响应
	app.ResponseSuccess(c, res)
}

// GetFansList 查询粉丝列表
// @Summary 查询关注列表
// @Description Author：YanSongWu 2021/02/18
// @Tags 用户：关注管理 Follow Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.GetFollowList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseGetFansList
// @Router /api/follow/fans [get]
func GetFansList(c *gin.Context) {
	p := new(dto.GetFollowList)
	// 获取上下文消息
	user, err := api.GetUserMessage(c)
	if err != nil {
		c.Error(err)
		zap.L().Error("GetUserMessage failed", zap.Error(err))
		return
	}
	if err := c.ShouldBindQuery(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("fans bind params failed", zap.String("Username", user.Username), zap.Error(err))
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
	if p.Id == 0 {
		p.Id = user.UserId
	}
	res, err := fo.GetFansList(p)
	if err != nil {
		zap.L().Error("get fans list failed", zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}
	// 返回响应
	app.ResponseSuccess(c, res)
}

// GetStatus 查看关注状态
// @Summary 查询关注状态
// @Description Author：YanSongWu 2021/02/18 获得身份令牌
// @Tags 用户：关注管理 Follow Controller
// @Accept application/json
// @Produce application/json
// @Param id query string false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSuccess
// @Router /api/follow/judge [get]
func GetStatus(c *gin.Context) {
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("GetUserMessage failed", zap.Error(err))
	}
	// 1.获取参数，设置默认值，校验参数
	id := c.Query("id")
	idI, err := strconv.Atoi(id)
	if err != nil {
		zap.L().Error("string to int failed", zap.Error(err))
		app.ResponseError(c, app.CodeParamIsInvalid)
		return
	}
	err = fo.GetFollowStatus(idI, user.UserId)
	if err != nil {
		zap.L().Error("get follow status failed", zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}
	app.ResponseSuccess(c, nil)
	return
}

// UpdateStatus 修改关注状态
// @Summary 修改关注状态
// @Description Author：Cgl 2021/02/18 获得身份令牌
// @Tags 用户：关注管理 Follow Controller
// @Accept application/json
// @Produce application/json
// @Param id query string false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSuccess
// @Router /api/follow/center [put]
func UpdateStatus(c *gin.Context) {
	// 获取上下文信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("GetUserMessage failed", zap.Error(err))
		return
	}
	// 绑定参数
	id := c.Query("id")
	idI, err := strconv.Atoi(id)
	if err != nil {
		zap.L().Error("string to int failed", zap.String("username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamTypeBindError)
	}
	if idI == 0 {
		zap.L().Error("UpdateStatus failed", zap.String("username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamIsInvalid)
	}
	err = fo.UpdateStatus(idI, user.UserId)
	if err != nil {
		zap.L().Error("UpdateStatus failed", zap.String("username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeUpdateOperationFail)
		return
	}
	app.ResponseSuccess(c, nil)
}
