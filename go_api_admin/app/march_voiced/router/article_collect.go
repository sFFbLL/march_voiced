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
		//收藏列表页 by 李甲坤
		r.GET("article", apis.GetCollectArticle)
		//收藏 by 李甲坤
		r.POST("article", apis.CollectArticle)
	}
}
