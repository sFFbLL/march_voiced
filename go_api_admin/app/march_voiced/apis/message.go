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

// SelectDept 查询是否有未读消息
// @Summary 查询是否有未读消息
// @Description Author：JiaKun Li 2021/02/17
// @Tags 消息 message Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.MessageUnreadDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseMessageUnRead
// @Router /api/message/user/unread-count [get]
func MessageUnreadCount(c *gin.Context) {
	p := new(dto.MessageUnreadDto)

	// 获取缓存信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("MessageUnread GetUserMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeLoginExpire)
		return
	}

	// 获取参数 校验参数
	if err := c.ShouldBindQuery(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("MessageUnread params failed", zap.String("Username", user.Username), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}

	//业务逻辑处理
	s := new(service.Message)
	count, err := s.MessageUnread(p, user.UserId)
	if err != nil {
		zap.L().Error("MessageUnread service params failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	app.ResponseSuccess(c, count)
}

// SelectDept 未读消息已读
// @Summary 未读消息已读
// @Description Author：JiaKun Li 2021/02/18
// @Tags 消息 message Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.MessageUnreadDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSuccess
// @Router /api/message/user [post]
func UnReadMessageRead(c *gin.Context)  {
	p := new(dto.UnReadMessageRead)

	// 获取缓存信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("UnReadMessageRead GetUserMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeLoginExpire)
		return
	}

	// 获取参数 校验参数
	if err := c.Bind(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("UnReadMessageRead params failed", zap.String("Username", user.Username), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}

	//业务逻辑处理
	s := new(service.Message)
	err = s.UnReadMessageRead(p, user.UserId)
	if err != nil {
		zap.L().Error("UnReadMessageRead service params failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeUpdateOperationFail)
		return
	}

	app.ResponseSuccess(c, nil)
}