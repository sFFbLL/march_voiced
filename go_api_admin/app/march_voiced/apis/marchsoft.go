package apis

import (
	"project/app/march_voiced/service"
	"project/common/api"
	"project/utils/app"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SelectDept 申请加入三月圈
// @Summary 申请加入三月圈
// @Description Author：JiaKun Li 2021/02/17
// @Tags 三月圈 marchsoft Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.ArticleFavourDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSuccess
// @Router /api/apply/march [post]
func ApplyMarch(c *gin.Context) {
	// 获取缓存信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("ApplyMarch GetUserMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeLoginExpire)
		return
	}

	//业务逻辑处理
	s := new(service.Marchsoft)
	err = s.ApplyMarch(user.UserId)
	if err != nil {
		zap.L().Error("ApplyMarch service params failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeUpdateOperationFail)
		return
	}

	app.ResponseSuccess(c, nil)
}

