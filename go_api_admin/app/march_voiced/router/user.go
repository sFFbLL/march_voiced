package router

import (
	"project/app/march_voiced/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, userRouter)
}

// 需认证的路由代码
func userRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/user")
	{
		r.GET("info/:id", apis.SelectUserInfo)
	}
}
