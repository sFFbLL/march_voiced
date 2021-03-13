package router

import (
	"project/app/march_voiced/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, marchFavourAuthRouter)
}

// 需认证的路由代码
func marchFavourAuthRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/favour")
	{
		//三月圈点赞 by 李甲坤
		r.POST("march", apis.MarchFavour)
	}
}
