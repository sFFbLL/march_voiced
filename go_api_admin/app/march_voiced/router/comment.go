package router

import (
	"project/app/march_voiced/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, commentRouter)
}

// 需要认证的路由代码
func commentRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/comment")
	{
		r.POST("article", apis.AddArticleComment)
		r.DELETE("article", apis.DeleteArticleComment)
		r.GET("article", apis.GetArticleComment)

		r.POST("marchsoft", apis.AddMarchsoftComment)
		r.DELETE("marchsoft", apis.DeleteMarchsoftComment)
		r.GET("marchsoft", apis.GetMarchsoftComment)

		r1 := r.Group("/children")
		{
			r1.GET("article", apis.GetArticleChildComment)
			r1.GET("marchsoft", apis.GetMarchsoftChildComment)

		}
	}
}
