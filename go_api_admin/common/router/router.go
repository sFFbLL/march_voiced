package router

import (
	"net/http"
	"project/app/march_voiced/apis"
	"project/common/global"
	"project/utils/config"

	admin "project/app/admin/router"
	march "project/app/march_voiced/router"
	"project/common/middleware"
	"project/utils"

	//_ "project/docs"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Setup 路由设置
func Setup(cfg *config.Application) *gin.Engine {
	if cfg.Mode == string(utils.ModeProd) {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(middleware.Cors(), middleware.GinLogger(), middleware.Sentinel(200))
	r.Static(cfg.StaticFileUrl, cfg.StaticPath)
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 注册业务路由
	admin.InitAdminRouter(r)
	march.InitMarchRouter(r)

	r.GET("/wx/ping", func(c *gin.Context) {
		server := global.Wx.GetServer(c.Request, c.Writer)
		err := server.Serve()
		if err != nil {
			return
		}
		server.Send()
	})
	r.POST("/", apis.WxGetTicket)

	pprof.Register(r) // 注册pprof相关路由
	return r
}
