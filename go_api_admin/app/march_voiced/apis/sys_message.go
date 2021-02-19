package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"project/app/march_voiced/models/dto"
	"project/app/march_voiced/service"
	"project/common/api"
	"project/utils/app"
)

// GetSysMessage 系统消息
// @Summary 系统消息
// @Description Author：JiaKun Li 2021/02/19
// @Tags 消息：Message Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.Paginator false "添加参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseGetSysMessage
// @Router /api/message/system [get]
func GetSysMessage(c *gin.Context) {
	p := new(dto.Paginator)

	// 获取缓存信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("GetSysMessage GetUserMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeLoginExpire)
		return
	}

	// 获取参数 校验参数
	if err := c.ShouldBindQuery(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("GetSysMessage params failed", zap.String("Username", user.Username), zap.Error(err))
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
	s := new(service.SysMessage)
	getSysMessage, err := s.GetSysMessage(p, user.UserId)
	if err != nil {
		zap.L().Error("GetSysMessage service params failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	app.ResponseSuccess(c, getSysMessage)
}
