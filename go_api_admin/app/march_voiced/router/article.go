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
		r.DELETE(":id", apis.DeleteArticleHandler)
		r.GET("top", apis.TopArticleListHandler)
		r.GET("word", apis.MatchSensitiveWord)
		r.GET("/user", apis.SelectArticleListByUserId)
		r.GET("/index", apis.SelectArticleListIndex)
		r.GET("/favour-collect/:id", apis.IsFavourCollectByArticleId)
		r.GET("/tag", apis.GetArticleTagList)
		r.POST("/reprint", apis.ReprintArticleHandler)
		//文章审核列表页 by 李甲坤
		r.GET("/admin", apis.ApplyArticleList)
		//文章设置推荐 by 李甲坤
		r.PUT("recommend/:id", apis.ArticleRecommend)
	}
	v1.GET("/base/searchArticle", apis.ArticleSearch)
}
