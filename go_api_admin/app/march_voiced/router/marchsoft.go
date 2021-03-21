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
		// 三月圈基本信息
		mr.GET("/msg", apis.GetMarchMsg)
		// 发布三月圈
		mr.POST("", apis.InsertMarchSoft)
		// 删除三月圈
		mr.DELETE(":id", apis.DeleteMarchHandler)
		// 三月圈列表页
		mr.GET("", apis.SelectMarchList)
		// 用户三月圈列表页
		mr.GET("/user", apis.SelectMarchListById)
		// 三月圈详情页
		mr.GET("/detail/:id", apis.MarchDetail)
	}
}
