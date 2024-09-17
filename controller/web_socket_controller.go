package controller

import (
	"github.com/XiaoLFeng/go-general-utils/bcode"
	"github.com/XiaoLFeng/go-gin-util/berror"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

// 创建一个 WebSocket 升级器
var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有连接
	},
}

func WsController(c *gin.Context) {
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

	// 启动一个 goroutine 定期发送消息
	go func() {
		for {
			time.Sleep(5 * time.Second) // 每5秒发送一次消息
			message := "Server time: " + time.Now().Format(time.RFC3339)
			if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				return
			}
		}
	}()

	for {
		// 读取客户端消息
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// 将消息发送回客户端，实现实时同步
		if err := conn.WriteMessage(messageType, message); err != nil {
			break
		}
	}
}
