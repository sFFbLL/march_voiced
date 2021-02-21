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
		// MessageUnreadCount 查询是否有未读消息
		r.GET("user/unread-count", apis.MessageUnreadCount)
		// UnReadMessageRead 未读消息已读
		r.POST("user", apis.UnReadMessageRead)
		// GetMessage 关注页面(动态)列表页
		r.GET("follow", apis.GetMessage)
		r.GET("user", apis.GetMessageMe)
	}
}
