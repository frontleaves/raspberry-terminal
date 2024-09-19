package mqtt

import (
	"github.com/XiaoLFeng/go-gin-util/blog"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"raspberry-terminal/config/c"
)

// SetupMQTT
//
// # 初始化 MQTT 客户端
//
// 用于初始化 MQTT 客户端，订阅多个主题，进行数据处理操作。
func SetupMQTT() {
	c.MqttClient = mqtt.NewClient(mqtt.NewClientOptions().AddBroker(c.Broker).SetClientID(c.ClientId))
	// 创建 MQTT 客户端
	if token := c.MqttClient.Connect(); token.Wait() && token.Error() != nil {
		blog.Fatalf("MQTT", "连接失败：%v", token.Error())
	}
	blog.Info("MQTT", "连接成功")

	getDeviceActiveOperation()
	getDeviceAuth()
}
