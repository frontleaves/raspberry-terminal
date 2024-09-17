package main

import (
	"embed"
	"github.com/XiaoLFeng/go-gin-util/blog"
	"github.com/gin-gonic/gin"
	"raspberry-terminal/route"
)

//go:embed "resources/dist/**"
var staticFS embed.FS

func main() {
	engine := gin.Default()

	// 路由
	route.Route(engine, staticFS)

	// 启动服务
	if engine.Run(":8080") != nil {
		blog.Warnf("STAT", "服务器启动失败")
	}
}
