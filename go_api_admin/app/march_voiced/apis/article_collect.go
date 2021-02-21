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

// SelectDept 我的收藏列表页
// @Summary 我的收藏列表页
// @Description Author：JiaKun Li 2021/02/16
// @Tags 文章 article Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.CollectArticleDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSuccess
// @Router /api/collect/article [get]
func GetCollectArticle(c *gin.Context) {

}

// SelectDept 文章收藏
// @Summary 文章收藏
// @Description Author：JiaKun Li 2021/02/16
// @Tags 文章 collect Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.CollectArticleDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSuccess
// @Router /api/collect/article [post]
func CollectArticle(c *gin.Context) {
	p := new(dto.CollectArticleDto)

	// 获取缓存信息
	user, err := api.GetUserMessage(c)
	if err != nil {
		zap.L().Error("CollectArticle GetUserMsg failed", zap.Error(err))
		app.ResponseError(c, app.CodeLoginExpire)
		return
	}

	// 获取参数 校验参数
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("CollectArticle params failed", zap.String("Username", user.Username), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamTypeBindError)
		return
	}

	//业务逻辑处理
	s := new(service.ArticleCollect)
	err = s.AddArticleCollect(p, user.UserId)
	if err != nil {
		zap.L().Error("AddArticleCollect service params failed", zap.String("Username", user.Username), zap.Error(err))
		app.ResponseError(c, app.CodeUpdateOperationFail)
		return
	}

	app.ResponseSuccess(c, nil)
}
