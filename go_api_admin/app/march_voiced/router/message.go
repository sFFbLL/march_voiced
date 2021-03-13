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
		// MessageUnreadCount 查询是否有未读消息 by 李甲坤
		r.GET("user/unread-count", apis.MessageUnreadCount)
		// UnReadMessageRead 未读消息已读 by 李甲坤
		r.PUT("user", apis.UnReadMessageRead)
		// GetMessage 关注页面(动态)列表页 by 李甲坤
		r.GET("follow", apis.GetMessage)
		// 互动消息 by 李甲坤
		r.GET("user", apis.GetMessageMe)
	}
}
