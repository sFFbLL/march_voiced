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

// SelectDept 三月圈点赞
// @Summary 三月圈点赞
// @Description Author：JiaKun Li 2021/02/17
// @Tags 三月圈 marchsoft Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.ArticleFavourDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSuccess
// @Router /api/favour/march [post]
func MarchFavour(c *gin.Context) {
	p := new(dto.MarchFavourDto)

	// 获取缓存信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("MarchFavour GetUserMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeLoginExpire)
		return
	}

	// 获取参数 校验参数
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("MarchFavour params failed", zap.String("Username", user.Username), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}

	//业务逻辑处理
	s := new(service.MarchFavour)
	err = s.AddMarchFavour(p, user.UserId)
	if err != nil {
		zap.L().Error("MarchFavour service params failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeUpdateOperationFail)
		return
	}

	app.ResponseSuccess(c, nil)
}
