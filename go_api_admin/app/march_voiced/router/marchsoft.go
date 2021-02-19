package router

import (
	"project/app/march_voiced/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, marchRouter)
}

// 需认证的路由代码
func marchRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/apply")
	{
		r.GET("march", apis.GetMarchApplyUser)
		r.POST("march", apis.ApplyMarch)
		r.PUT("march", apis.MarchPass)
		r.PUT("article", apis.ArticlePass)
	}
}
