package mqtt

import (
	"encoding/json"
	"github.com/XiaoLFeng/go-gin-util/blog"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"raspberry-terminal/config/c"
	"raspberry-terminal/model/dto"
	"raspberry-terminal/model/entity"
	"time"
)

// getDeviceActiveOperation
//
// # 获取设备主动操控信息
//
// 获取设备主动操控信息，返回设备主动操控信息。
func getDeviceActiveOperation() {
	// 订阅 MQTT 主题
	token := c.MqttClient.Subscribe(c.TopicOperate, 0, func(client mqtt.Client, msg mqtt.Message) {
		blog.Debugf("MQTT", "接收到消息：%v", string(msg.Payload()))
		// 对获取的数据进行 json 解码到结构体
		var unmarshal *dto.MqttDeviceOperateDTO
		err := json.Unmarshal(msg.Payload(), &unmarshal)
		if err != nil {
			blog.Warnf("MQTT", "消息反序列化失败：%v", err.Error())
		}
		// 验证设备是否授权
		var device *entity.Device
		c.DB.Last(&device, "device_name = ?", unmarshal.Device)
		if device.UUID == uuid.Nil || !device.Authorized {
			blog.Tracef("MQTT", "设备未授权：%v", unmarshal.Device)
			return
		}
		// 保存设备主动操控信息
		marshal, err := json.Marshal(unmarshal.Data)
		if err != nil {
			blog.Warnf("MQTT", "消息序列化失败：%v", err.Error())
		}
		device.Login = true
		device.NowValue = string(marshal)
		c.DB.Save(&device)
		// 保存日志信息
		c.DB.Create(&entity.Log{
			Operate: "DRIVING",
			LogUUID: device.UUID,
			Content: device.DeviceUsername + ":" + unmarshal.Device,
		})
	})

	if token.Wait() && token.Error() != nil {
		blog.Fatalf("MQTT", "MQTT 订阅失败：%v", token.Error())
	}

	time.Sleep(time.Second)
}
