package router

import (
	"project/app/march_voiced/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, followRouter)
}

// 需要认证的路由代码
func followRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/follow")
	{
		r.GET("", apis.GetFollowList)
		r.GET("fans", apis.GetFansList)
		r.GET("judge", apis.GetStatus)
		r.POST("", apis.UpdateStatus)
	}
}
