package router

import (
	"project/app/march_voiced/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, articleCollectAuthRouter)
}

// 需认证的路由代码
func articleCollectAuthRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/collect")
	{
		r.GET("article", apis.GetCollectArticle)
		r.POST("article", apis.CollectArticle)
	}
}
