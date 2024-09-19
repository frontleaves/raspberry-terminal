package route

import (
	"embed"
	"github.com/XiaoLFeng/go-gin-util/bmiddle"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
	"raspberry-terminal/controller"
)

func Route(engine *gin.Engine, staticFS embed.FS) {
	// 引入 Handler
	engine.Use(bmiddle.CrossDomainClearingMiddleware())

	// 静态资源
	st, _ := fs.Sub(staticFS, "resources/dist/assets")
	engine.StaticFS("/assets", http.FS(st))

	// 首页重定向
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/home")
	})

	// WebSocket
	wsGroup := engine.Group("/ws")
	{
		wsGroup.GET("/ping", controller.SocketPingController)
		wsGroup.GET("/system", controller.SocketSystemController)
		wsGroup.GET("/mqtt", controller.SocketMqttController)
	}

	// 首页
	homeGroup := engine.Group("/home")
	{
		homeGroup.GET("/*filepath", func(c *gin.Context) {
			file, _ := staticFS.ReadFile("resources/dist/index.html")
			c.Data(http.StatusOK, "text/html", file)
		})
	}

	// API
	apiGroup := engine.Group("/api/v1")
	{
		apiGroup.Use(bmiddle.ReturnResultMiddleware())
		systemGroup := apiGroup.Group("/system")
		{
			systemGroup.GET("/info", controller.ApiSystemInfoController)
		}
	}

	engine.NoRoute(bmiddle.NoRouteMiddleware())
}
