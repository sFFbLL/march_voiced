package router

import (
	"project/app/march_voiced/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, articleAuthRouter)
}

// 需认证的路由代码
func articleAuthRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/article")
	{
		r.GET("detail/:id", apis.ArticleDetailHandler)
		r.POST("", apis.InsertArticleHandler)
		r.PUT("", apis.UpdateArticleHandler)
		r.POST("reprint", apis.ReprintArticleHandler)
		r.GET("top", apis.TopArticleListHandler)
		r.GET("word", apis.MatchSensitiveWord)
	}

}
