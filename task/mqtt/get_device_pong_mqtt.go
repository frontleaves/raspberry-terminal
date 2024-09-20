package mqtt

import (
	"encoding/json"
	"github.com/XiaoLFeng/go-gin-util/blog"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"raspberry-terminal/config/c"
	"raspberry-terminal/model/dto"
	"raspberry-terminal/model/entity"
	"time"
)

// getDevicePong
//
// # 获取设备心跳
//
// 对设备进行心跳操作，获取订阅后的设备心跳信息内容。
func getDevicePong() {
	// 订阅 MQTT 主题
	token := c.MqttClient.Subscribe(c.TopicAuth, 0, func(client mqtt.Client, msg mqtt.Message) {
		blog.Debugf("MQTT", "接收到消息：%v", string(msg.Payload()))
		// 对获取的数据进行 json 解码到结构体
		var unmarshal *dto.MqttDevicePongDTO
		err := json.Unmarshal(msg.Payload(), &unmarshal)
		if err != nil {
			blog.Warnf("MQTT", "消息解码失败：%v", err.Error())
			return
		}
		var device *entity.Device
		c.DB.Last(&device, "device_name = ?", unmarshal.Device)
		if !device.Authorized {
			blog.Tracef("MQTT", "设备 %s 未授权", unmarshal.Device)
			return
		}
		// 检查误差时间是否超过 3 秒
		if time.Now().After(unmarshal.Now.Add(time.Second * 3)) {
			blog.Warnf("MQTT", "设备 %s 心跳超时：%v", unmarshal.Device, unmarshal.Now)
			return
		}
		// 更新设备心跳时间
		device.Login = true
		device.Uptime = time.Now()
		c.DB.Save(&device)
	})

	if token.Wait() && token.Error() != nil {
		blog.Fatalf("MQTT", "MQTT 订阅失败：%v", token.Error())
	}

	time.Sleep(time.Second)
}
