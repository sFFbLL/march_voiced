package router

import (
	"project/app/march_voiced/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, sysMessageAuthRouter)
}

// 需认证的路由代码
func sysMessageAuthRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/message")
	{
		r.GET("system", apis.GetSysMessage)
	}
}
