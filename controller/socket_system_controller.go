package controller

import (
	"encoding/json"
	"github.com/XiaoLFeng/go-general-utils/bcode"
	"github.com/XiaoLFeng/go-gin-util/berror"
	"github.com/XiaoLFeng/go-gin-util/blog"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"raspberry-terminal/model/dto"
	"raspberry-terminal/service"
	"time"
)

// SocketSystemController
//
// # WebSocket 系统信息
//
// 用于 WebSocket 实时获取系统信息。
//
// # 请求
//   - c: gin.Context Gin 上下文
func SocketSystemController(c *gin.Context) {
	// 升级协议为 WebSocket
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		_ = c.Error(berror.New(bcode.BadRequestInvalidInput, "升级为 WebSocket 失败"))
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	for {
		time.Sleep(1 * time.Second)
		cpuPercent, err := service.GetCpuPercent()
		if err != nil {
			break
		}
		ramPercent, err := service.GetRamPercent()
		if err != nil {
			break
		}
		diskPercent, err := service.GetDiskPercent()
		if err != nil {
			break
		}
		getData := dto.WsSystemInfoDTO{
			Cpu:            cpuPercent,
			CpuTemperature: "",
			Ram:            ramPercent,
			Disk:           diskPercent,
		}
		marshal, err := json.Marshal(getData)
		if err != nil {
			blog.Warnf("WS", "消息序列化失败，%v", err.Error())
		}
		// 将消息发送回客户端，实现实时同步
		err = conn.WriteMessage(websocket.TextMessage, marshal)
		if err != nil {
			break
		}
	}
}
