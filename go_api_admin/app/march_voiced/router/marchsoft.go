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
		// （后台）查询申请三月圈的用户 by 李甲坤
		r.GET("march", apis.GetMarchApplyUser)
		//申请加入三月圈 by 李甲坤
		r.POST("march", apis.ApplyMarch)
		//（后台）三月圈审核通过或驳回 by 李甲坤
		r.PUT("march", apis.MarchPass)
		// （后台）发布文章审核通过或驳回 by 李甲坤
		r.PUT("article", apis.ArticlePass)
	}

	mr := v1.Group("/march")
	{
		mr.GET("/msg", apis.GetMarchMsg)
		mr.POST("", apis.InsertMarchSoft)
		mr.DELETE(":id", apis.DeleteMarchHandler)
		mr.GET("", apis.SelectMarchList)
		mr.GET("/user", apis.SelectMarchListById)
		mr.GET("/detail/:id", apis.MarchDetail)
	}
}
