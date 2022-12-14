package router

import (
	"project/app/march_voiced/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, followMessageAuthRouter)
}

// 需认证的路由代码
func followMessageAuthRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/message")
	{
		//关注消息 by 李甲坤
		r.GET("follow-notice", apis.GetFollowMessage)
	}
}
