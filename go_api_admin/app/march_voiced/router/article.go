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
		// 文章详情页
		r.GET("detail/:id", apis.ArticleDetailHandler)
		// 添加文章
		r.POST("", apis.InsertArticleHandler)
		// 修改文章
		r.PUT("", apis.UpdateArticleHandler)
		// 删除文章
		r.DELETE(":id", apis.DeleteArticleHandler)
		// 推荐文章列表页
		r.GET("top", apis.TopArticleListHandler)
		// 用户文章列表
		r.GET("/user", apis.SelectArticleListByUserId)
		// 未推荐文章列表
		r.GET("/index", apis.SelectArticleListIndex)
		// 文章是否点赞收藏
		r.GET("/favour-collect/:id", apis.IsFavourCollectByArticleId)
		// 获取文章标签列表
		r.GET("/tag", apis.GetArticleTagList)
		// 文章转发
		r.POST("/reprint", apis.ReprintArticleHandler)
		// 文章敏感词匹配
		r.GET("word", apis.MatchSensitiveWord)
		//文章审核列表页 by 李甲坤
		r.GET("/admin", apis.ApplyArticleList)
		//文章设置推荐 by 李甲坤
		r.PUT("recommend/:id", apis.ArticleRecommend)
	}
	v1.GET("/base/searchArticle", apis.ArticleSearch)
}
