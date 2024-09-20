package mqtt

import (
	"encoding/json"
	"github.com/XiaoLFeng/go-gin-util/blog"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"raspberry-terminal/config/c"
	"raspberry-terminal/model/dto"
	"raspberry-terminal/model/entity"
	"regexp"
	"strings"
	"time"
)

// getDeviceAuth
//
// # 获取设备授权
//
// 对设备进行授权操作，获取订阅后的设备授权信息内容。
func getDeviceAuth() {
	for {
		// 订阅 MQTT 主题
		token := c.MqttClient.Subscribe(c.TopicAuth, 0, func(client mqtt.Client, msg mqtt.Message) {
			blog.Debugf("MQTT", "接收到消息：%v", string(msg.Payload()))
			// 对获取的数据进行 json 解码到结构体
			var unmarshal *dto.MqttAuthDTO
			err := json.Unmarshal(msg.Payload(), &unmarshal)
			if err == nil {
				// 使用设备信息获取数据库登记状态
				var device *entity.Device
				c.DB.Last(&device, "device_name = ?", unmarshal.Device)
				if device.UUID == uuid.Nil {
					// 注册设备
					device = &entity.Device{
						Type:           strings.Replace(regexp.MustCompile(`-(\w+)-`).FindStringSubmatch(unmarshal.Device)[1], "-", "", -1),
						DeviceName:     unmarshal.Device,
						DeviceUsername: unmarshal.Username,
						DevicePassword: unmarshal.Password,
						DeviceHost:     unmarshal.Host,
						DeviceMac:      unmarshal.Mac,
						Authorized:     false,
						CreatedAt:      time.Now(),
					}
					c.DB.Create(device)
					// 检查是否登录
					authReturnDTO := &dto.MqttAuthReturnDTO{
						Device:     unmarshal.Device,
						Authorized: false,
					}
					marshal, err := json.Marshal(authReturnDTO)
					if err == nil {
						c.MqttClient.Publish(c.TopicAuthReturn, 0, false, marshal)
					} else {
						blog.Warnf("MQTT", "消息序列化失败：%v", err.Error())
					}
				} else {
					// 检查是否登录
					authReturnDTO := &dto.MqttAuthReturnDTO{
						Device:     unmarshal.Device,
						Authorized: false,
					}
					// 检查是否授权
					authReturnDTO.Authorized = device.Authorized
					// 检查用户名与密码
					if device.Authorized {
						if unmarshal.Username == device.DeviceUsername && unmarshal.Password == device.DevicePassword {
							authReturnDTO.Authorized = true
						} else {
							authReturnDTO.Authorized = false
						}
					}
					device.Login = true
					device.Uptime = time.Now()
					c.DB.Save(&device)
					marshal, err := json.Marshal(authReturnDTO)
					if err == nil {
						c.MqttClient.Publish(c.TopicAuthReturn, 0, false, marshal)
					} else {
						blog.Warnf("MQTT", "消息序列化失败：%v", err.Error())
					}
				}
			} else {
				blog.Warnf("MQTT", "消息反序列化失败：%v", err.Error())
			}
		})

		if token.Wait() && token.Error() != nil {
			blog.Fatalf("MQTT", "订阅失败：%v", token.Error())
		}

		time.Sleep(time.Minute)
	}
}
