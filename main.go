package main

import (
	"embed"
	"github.com/XiaoLFeng/go-gin-util/bconfig"
	"github.com/XiaoLFeng/go-gin-util/blog"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"raspberry-terminal/config/startup"
	"raspberry-terminal/route"
	"raspberry-terminal/task/mqtt"
)

//go:embed "resources/dist/**"
var staticFS embed.FS

func main() {
	engine := gin.Default()
	// 配置文件初始化
	bconfig.LogConfiguration(".logs", "logger", true, logrus.TraceLevel)
	// 系统初始化
	startup.Initial()

	// 启动 MQTT 服务(内含定时任务)
	mqtt.SetupMQTT()

	// 路由
	route.Route(engine, staticFS)
	// 启动服务
	if engine.Run(":8080") != nil {
		blog.Warnf("STAT", "服务器启动失败")
	}
}
