package router

import (
	"project/app/march_voiced/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, messageAuthRouter)
}

// 需认证的路由代码
func messageAuthRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/message")
	{
		r.GET("user/unread-count", apis.MessageUnreadCount)
		r.POST("user", apis.UnReadMessageRead)
	}
}
