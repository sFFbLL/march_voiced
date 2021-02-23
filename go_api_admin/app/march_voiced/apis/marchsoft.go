package apis

import (
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/app/march_voiced/service"
	"project/common/api"
	"project/utils"
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

// GetMarchMsg 三月圈条数与成员数与是否为三月
// @Summary 三月圈条数与成员数与是否为三月
// @Description Author：lbl 2021/02/20
// @Tags 三月圈 marchsoft Controller
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSoftInfo
// @Router /api/march/msg [get]
func GetMarchMsg(c *gin.Context) {
	marchMsg := new(bo.MarchSoftInfo)
	s := new(service.Marchsoft)
	var err error

	// 获取缓存信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("GetMarchMsg GetUserMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeLoginExpire)
		return
	}

	//业务逻辑处理
	marchMsg, err = s.GetMarchMsg(user.Username, user.UserId)
	if err != nil {
		zap.L().Error("GetMarchApplyUser service failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	app.ResponseSuccess(c, marchMsg)
}

// InsertMarchSoft 发布三月圈
// @Summary 发布三月圈
// @Description Author：lbl 2021/02/20
// @Tags 三月圈 marchsoft Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.InsertMarchSoft false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSuccess
// @Router /api/march [post]
func InsertMarchSoft(c *gin.Context) {
	march := new(dto.InsertMarchSoft)
	s := new(service.Marchsoft)
	var err error

	// 获取缓存信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("InsertMarchSoft GetUserMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeLoginExpire)
		return
	}

	// 获取参数 校验参数
	if err := c.ShouldBindJSON(march); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("InsertMarchSoft params failed", zap.String("Username", user.Username), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}

	//业务逻辑处理
	err = s.InsertMarchSoft(march, uint(user.UserId))
	if err != nil {
		zap.L().Error("InsertMarchSoft service failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeInsertOperationFail)
		return
	}

	app.ResponseSuccess(c, nil)
}

// DeleteMarchHandler 删除三月圈
// @Summary 删除三月圈
// @Description Author：Lbl 2021/02/17 获得身份令牌
// @Tags 应用：文章管理 Article Controller
// @Accept application/json
// @Produce application/json
// @Param id path int false "修改参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSuccess
// @Router /api/march/{id} [delete]
func DeleteMarchHandler(c *gin.Context) {
	// 声明必要变量
	a := new(service.Marchsoft)
	userMsg := new(api.UserMessage)
	var err error
	var id int

	// 获取上下文中的用户信息
	userMsg, err = api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("DeleteMarchHandler Get userId failed", zap.Error(err))
		app.ResponseError(c, app.CodeNoUser)
		return
	}

	// 获取参数
	id, err = utils.StringToInt(c.Param("id"))
	if err != nil {
		zap.L().Error("DeleteMarchHandler bind Query Failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 进入service层对数据操作
	err = a.DeleteMarchSoft(id, userMsg.UserId)
	if err != nil {
		zap.L().Error("DeleteMarchHandler service failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeDeleteOperationFail)
		return
	}

	// 成功返回状态
	app.ResponseSuccess(c, nil)
}

// SelectMarchList 三月圈文章列表
// @Summary 三月圈文章列表
// @Description Author：lbl 2021/02/20
// @Tags 三月圈 marchsoft Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.Paging false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseMarch
// @Router /api/march [get]
func SelectMarchList(c *gin.Context) {
	userMsg := new(api.UserMessage)
	MarchList := new([]bo.March)
	m := new(service.Marchsoft)
	var paging dto.Paging
	var err error

	// 获取上下文用户信息
	userMsg, err = api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("SelectMarchList Get userId failed", zap.Error(err))
		app.ResponseError(c, app.CodeNoUser)
		return
	}

	// 参数绑定
	paging.Current, err = utils.StringToUint(c.DefaultQuery("current", "1"))
	if err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("SelectMarchList params Current failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}
	paging.Size, err = utils.StringToUint(c.DefaultQuery("size", "10"))
	if err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("SelectMarchList params Size failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}

	// 进入service层对数据操作
	MarchList, err = m.SelectMarchList(paging, uint(userMsg.UserId))
	if err != nil {
		zap.L().Error("SelectMarchList service failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	// 成功返回状态
	app.ResponseSuccess(c, MarchList)

}

// SelectMarchListById 用户三月圈文章列表
// @Summary 用户三月圈文章列表
// @Description Author：lbl 2021/02/20
// @Tags 三月圈 marchsoft Controller
// @Accept application/json
// @Produce application/json
// @Param object query dto.SelectMarchListById false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseMarch
// @Router /api/march [get]
func SelectMarchListById(c *gin.Context) {
	// 声明必要变量
	m := new(service.Marchsoft)
	userMsg := new(api.UserMessage)
	MarchList := new([]bo.MarchByUserId)
	var paging dto.SelectMarchListById
	var err error

	// 获取上下文用户信息
	userMsg, err = api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("SelectMarchListById Get userMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeNoUser)
		return
	}
	err = c.ShouldBindQuery(&paging)
	if err != nil {
		zap.L().Error("SelectMarchListById ShouldBindQuery Params Failed", zap.String("Username", userMsg.Username), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	if paging.Current == 0 {
		paging.Current = 1
	}
	if paging.Size == 0 {
		paging.Current = 10
	}

	// 进入service层对数据操作
	MarchList, err = m.SelectMarchListById(paging, uint(userMsg.UserId))
	if err != nil {
		zap.L().Error("SelectMarchListById service failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	// 成功返回状态
	app.ResponseSuccess(c, MarchList)
}

// MarchDetail 三月圈文章详情页
// @Summary 三月圈文章详情页
// @Description Author：Lbl 2021/02/17 获得身份令牌
// @Tags 三月圈 marchsoft Controller
// @Accept application/json
// @Produce application/json
// @Param id path int false "修改参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseMarch
// @Router /api/march/{id} [get]
func MarchDetail(c *gin.Context) {
	// 声明必要变量
	m := new(service.Marchsoft)
	userMsg := new(api.UserMessage)
	march := new(bo.March)
	var id int
	var err error

	// 获取上下文中的用户信息
	userMsg, err = api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("MarchDetail Get userId failed", zap.Error(err))
		app.ResponseError(c, app.CodeNoUser)
		return
	}

	// 获取参数
	id, err = utils.StringToInt(c.Param("id"))
	if err != nil {
		zap.L().Error("MarchDetail bind Query Failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 进入service层对数据操作
	march, err = m.MarchDetail(id, userMsg.UserId)
	if err != nil {
		zap.L().Error("MarchDetail service failed", zap.String("Username", userMsg.Username), zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}

	// 成功返回状态
	app.ResponseSuccess(c, march)
}
