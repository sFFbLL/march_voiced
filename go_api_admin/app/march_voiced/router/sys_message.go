package router

import (
	"project/app/march_voiced/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, sysMessageAuthRouter)
}

// 需认证的路由代码
func sysMessageAuthRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/message")
	{
		// 查看系统消息 by 李甲坤
		r.GET("system", apis.GetSysMessage)
	}
}
