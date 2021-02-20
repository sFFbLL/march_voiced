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

// GetMarchApplyUser 查询申请三月圈审核中用户
// @Summary 查询申请三月圈审核中用户
// @Description Author：JiaKun Li 2021/02/17
// @Tags 三月圈 marchsoft Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.ApplyMarchPaginator false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseApplyMarchUser
// @Router /api/apply/march [get]
func GetMarchApplyUser(c *gin.Context) {
	p := new(dto.ApplyMarchPaginator)

	// 获取缓存信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("GetMarchApplyUser GetUserMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeLoginExpire)
		return
	}

	// 获取参数 校验参数
	if err := c.ShouldBindQuery(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("GetMarchApplyUser params failed", zap.String("Username", user.Username), zap.Error(err))
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
	s := new(service.Marchsoft)
	applyMarchUser, err := s.GetMarchApplyUser(p, user.UserId)
	if err != nil {
		zap.L().Error("GetMarchApplyUser service params failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	app.ResponseSuccess(c, applyMarchUser)
}

// ApplyMarch 申请加入三月圈
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

// MarchPass （后台）三月圈审核通过或驳回
// @Summary （后台）三月圈审核通过或驳回
// @Description Author：JiaKun Li 2021/02/18
// @Tags 三月圈：Marchsoft Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.MarchPass false "添加参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSuccess
// @Router /api/apply/march [put]
func MarchPass(c *gin.Context) {
	p := new(dto.MarchPass)

	// 获取缓存信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("MarchPass GetUserMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeLoginExpire)
		return
	}

	// 获取参数 校验参数
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("MarchPass params failed", zap.String("Username", user.Username), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}

	//业务逻辑处理
	s := new(service.Marchsoft)
	err = s.MarchPass(p, user.UserId)
	if err != nil {
		zap.L().Error("MarchPass service params failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeUpdateOperationFail)
		return
	}

	app.ResponseSuccess(c, nil)
}