package router

import (
	"project/app/march_voiced/apis"
	"project/utils/app"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, articleRouter)
	routerCheckRole = append(routerCheckRole, articleAuthRouter)
}

// 无需认证的路由代码
func articleRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/article")
	{
		r.GET("ping", func(c *gin.Context) {
			c.String(int(app.CodeSuccess), "ok")
		})
	}
}

// 需认证的路由代码
func articleAuthRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/article")
	{
		//r.GET("/:id")
		r.POST("", apis.InsertArticleHandler)
		r.PUT("", apis.UpdateArticleHandler)
		//r.DELETE("/:id")
		//r.GET("/")
		//r.POST("/superior")
	}
}
