package controller

import (
	"github.com/gorilla/websocket"
	"net/http"
)

// 创建一个 WebSocket 升级器
var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有连接
	},
}
