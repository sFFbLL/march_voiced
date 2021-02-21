package router

import (
	"github.com/gin-gonic/gin"
	"project/app/march_voiced/apis"
	"project/utils/app"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, searchRouter)
	routerCheckRole = append(routerCheckRole, searchActileRouter)
}

// 无需认证的路由代码
func searchRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/base")
	{
		r.GET("/searchActile",apis.SearchActile)
		r.GET("/searchUser", apis.SearchUser)
		r.GET("/ping", func(c *gin.Context) {
			c.String(int(app.CodeSuccess), "ok")
		})
	}
}

// 需认证的路由代码
func searchActileRouter(v1 *gin.RouterGroup) {

	r := v1.Group("/base")
	{
		r.GET("/informationsss", apis.SearchUserInfo)
	}
}

