package apis

import (
	"project/app/march_voiced/service"
	"project/common/api"
	"project/utils/app"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var u = new(service.User)

// SelectUserInfo 查询用户信息
// @Summary 查询用户信息
// @Description Author：YanSongWu 2021/02/19 获得身份令牌
// @Tags 应用：用户管理 User Controller
// @Accept application/json
// @Produce application/json
// @Param object path int false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSelectUserInfo
// @Router /api/user/:id [get]
func SelectUserInfo(c *gin.Context) {
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("GetUserMessage failed", zap.String("username:", user.Username), zap.Error(err))
		return
	}

	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		zap.L().Error("string to int failed", zap.Error(err))
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}

	// 返回响应
	res, err := u.SelectUserInfo(id, user.UserId)
	if err != nil {
		zap.L().Error("select user info failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}
	// 返回响应
	app.ResponseSuccess(c, res)
}
