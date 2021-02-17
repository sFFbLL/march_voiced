package router

import (
	"github.com/gin-gonic/gin"
	"project/app/march_voiced/apis"
	"project/utils/app"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, userRouter)
	routerCheckRole = append(routerCheckRole, userAuthRouter)
}

// 无需认证的路由代码
func userRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/collect")
	{
		r.GET("ping", func(c *gin.Context) {
			c.String(int(app.CodeSuccess), "ok")
		})
	}
}

// 需认证的路由代码
func userAuthRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/collect")
	{
		r.POST("article", apis.CollectArticle)
	}
}
