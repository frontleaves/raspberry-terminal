package controller

import (
	"github.com/XiaoLFeng/go-gin-util/bresult"
	"github.com/gin-gonic/gin"
)

// ApiSystemInfoController
//
// # 获取系统信息
//
// 获取系统信息，用于获取系统的 CPU、内存、磁盘等信息。
//
// # 请求
//   - c: gin.Context Gin 上下文
func ApiSystemInfoController(c *gin.Context) {
	bresult.OkWithData(c, "获取系统信息成功", gin.H{
		"cpu":    "Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz",
		"memory": "8G",
		"disk":   "256G",
	})
}
