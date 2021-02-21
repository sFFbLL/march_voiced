package router

import (
	"github.com/gin-gonic/gin"
	"project/app/march_voiced/apis"
)
func init() {
	routerNoCheckRole = append(routerNoCheckRole, usersRouter)
	routerCheckRole = append(routerCheckRole, userAuthRouter)
}
//func init() {
//
//	routerNoCheckRole = append(routerNoCheckRole, userRouter)
//	routerCheckRole = append(routerCheckRole, userAuthRouter)
//}

// 无需认证的路由代码
func usersRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/user")
	{
		r.GET("ll", apis.Aaa)

		r.POST("login", apis.LoginHandler)
		r.POST("creatNewUser",apis.CreatSysUser)
		r.GET("isExist",apis.SearchUsername)

	}
}

// 需认证的路由代码
func userAuthRouter(v1 *gin.RouterGroup) {

	r := v1.Group("/user")
	{
		//r.GET("/information", apis.SearchUserInfo)
		r.PUT("/modInformation", apis.ModInformation)
		r.GET("/information", apis.SearchUserInfo)
	}
}
