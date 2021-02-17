package router

import (
	"project/common/middleware"

	"github.com/gin-gonic/gin"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole   = make([]func(v1 *gin.RouterGroup), 0)
)

//InitMarchRouter 后台模块路由
func InitMarchRouter(r *gin.Engine) *gin.Engine {

	// 无需认证的路由
	marchNoCheckRoleRouter(r)
	// 需要认证的路由
	marchCheckRoleRouter(r)

	return r
}

func marchNoCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/api")
	// 空接口防止v1定义无使用报错
	v1.GET("/marchadag", nil)

	for _, f := range routerNoCheckRole {
		f(v1)
	}
}

func marchCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/api", middleware.JWTAuthMiddleware())
	// 空接口防止v1定义无使用报错
	v1.GET("/marchadaasdg", nil)

	for _, f := range routerCheckRole {
		f(v1)
	}
}
