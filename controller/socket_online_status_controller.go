package controller

import (
	"encoding/json"
	"github.com/XiaoLFeng/go-general-utils/bcode"
	"github.com/XiaoLFeng/go-gin-util/berror"
	"github.com/XiaoLFeng/go-gin-util/blog"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"raspberry-terminal/model/vo"
	"raspberry-terminal/service"
	"time"
)

func SocketOnlineStatusController(c *gin.Context) {
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

	for {
		time.Sleep(5 * time.Second)
		getDevices := service.DeviceGet(vo.CustomPageVO{
			Page:  1,
			Limit: 1000,
		})
		marshal, err := json.Marshal(getDevices)
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
