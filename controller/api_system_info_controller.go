package controller

import (
	"github.com/XiaoLFeng/go-gin-util/bresult"
	"github.com/gin-gonic/gin"
)

func ApiSystemInfoController(c *gin.Context) {
	bresult.OkWithData(c, "获取系统信息成功", gin.H{
		"cpu":    "Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz",
		"memory": "8G",
		"disk":   "256G",
	})
}
