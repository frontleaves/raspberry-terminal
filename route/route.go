package route

import (
	"embed"
	"github.com/XiaoLFeng/go-gin-util/bmiddle"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
	"raspberry-terminal/controller"
)

// Route
//
// # 路由
//
// 路由配置，用于配置路由。
//
// # 参数
//   - engine: *gin.Engine Gin 引擎
//   - staticFS: embed.FS 静态资源文件系统
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
		wsGroup.GET("/device", controller.SocketOnlineStatusController)
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

		// 设备
		deviceGroup := apiGroup.Group("/device")
		{
			deviceGroup.POST("/register", controller.ApiRegisterDeviceController)
			deviceListGroup := deviceGroup.Group("/list")
			{
				deviceListGroup.GET("/no-register", controller.ApiGetNoRegisterDeviceController)
				deviceListGroup.GET("/", controller.ApiGetDeviceController)
				deviceListGroup.GET("/online", controller.ApiGetOnlineDeviceController)
				deviceListGroup.GET("/offline", controller.ApiGetOfflineDeviceController)
			}
			deviceControlGroup := deviceGroup.Group("/control")
			{
				deviceControlGroup.POST("/light", controller.ApiDeviceLightController)
			}
		}
		// 系统
		systemGroup := apiGroup.Group("/system")
		{
			systemGroup.GET("/info", controller.ApiSystemInfoController)
		}
	}

	engine.NoRoute(bmiddle.NoRouteMiddleware())
}
