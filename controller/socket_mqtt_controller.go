package controller

import (
	"encoding/json"
	"github.com/XiaoLFeng/go-general-utils/bcode"
	"github.com/XiaoLFeng/go-gin-util/berror"
	"github.com/XiaoLFeng/go-gin-util/blog"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"time"
)

// 保存所有连接的 WebSocket 客户端
var wsConn *websocket.Conn
var mqttClient mqtt.Client

func init() {
	// 初始化 MQTT 客户端
	broker := "tcp://localhost:1883" // 替换为你的 MQTT 服务器地址
	topic := "test/topic"

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID("go_mqtt")

	// 创建 MQTT 客户端
	mqttClient = mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		blog.Fatalf("WS", "MQTT 连接失败：%v", token.Error())
	}

	// 订阅 MQTT 主题
	if token := mqttClient.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		blog.Infof("WS", "接收到消息：%v", string(msg.Payload()))
		// 广播到所有 WebSocket 客户端
		getMessage(msg.Payload())
	}); token.Wait() && token.Error() != nil {
		blog.Fatalf("WS", "MQTT 订阅失败：%v", token.Error())
	}
}

func getMessage(message []byte) {
	var unmarshal map[string]interface{}
	err := json.Unmarshal(message, &unmarshal)
	if err != nil {
		blog.Warnf("WS", "消息反序列化失败：%v", err.Error())
	}
	if unmarshal == nil || unmarshal["type"] == nil {
		return
	}
	switch unmarshal["type"].(string) {
	case "auth":
		// 认证消息
		authMessage := map[string]string{
			"action": "auth",
			"value":  "true",
		}
		marshal, err := json.Marshal(authMessage)
		if err != nil {
			blog.Warnf("WS", "消息序列化失败：%v", err.Error())
		}
		mqttClient.Publish("test/topic", 0, false, marshal)
	case "light":
		if unmarshal["value"] == nil {
			return
		}
		if unmarshal["value"].(string) == "on" {
			send([]byte("LED 灯已打开"))
		} else {
			send([]byte("LED 灯已关闭"))
		}
	}
}

// 广播 MQTT 消息到所有 WebSocket 客户端
func send(message []byte) {
	err := wsConn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		blog.Warnf("WS", "发送消息失败：%v", err.Error())
		err := wsConn.Close()
		if err != nil {
			return
		}
	}
}

// SocketMqttController handles WebSocket connections and forwards MQTT messages to clients
func SocketMqttController(c *gin.Context) {
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

	// 保存客户端
	wsConn = conn

	for {
		time.Sleep(1 * time.Second)
		// 读取客户端消息（保持连接）
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}
