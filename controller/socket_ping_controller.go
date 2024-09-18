package controller

import (
	"encoding/json"
	"github.com/XiaoLFeng/go-general-utils/bcode"
	"github.com/XiaoLFeng/go-gin-util/berror"
	"github.com/XiaoLFeng/go-gin-util/blog"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"raspberry-terminal/model/dto"
)

// SocketPingController
//
// # WebSocket 心跳检测
//
// 用于 WebSocket 心跳检测，客户端发送 ping 消息，服务端返回 pong 消息。
//
// # 请求
//   - c: gin.Context Gin 上下文
func SocketPingController(c *gin.Context) {
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		_ = c.Error(berror.New(bcode.BadRequestInvalidInput, "升级为 WebSocket 失败"))
		return
	}

	// 关闭后释放资源
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	// 接收消息
	var getMsg *dto.WsPingDTO
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		err = json.Unmarshal(message, &getMsg)
		if err != nil {
			blog.Warnf("WS", "消息反序列化失败，%v", err.Error())
		}
		var pong = map[string]string{
			"msg": "pong",
		}
		blog.Debugf("WS", "接收到消息：%v", message)
		marshal, err := json.Marshal(pong)
		if err != nil {
			blog.Warnf("WS", "消息序列化失败，%v", err.Error())
		}
		err = conn.WriteMessage(messageType, marshal)
		if err != nil {
			blog.Warnf("WS", "发送消息失败，%v", err.Error())
		}
	}
}
