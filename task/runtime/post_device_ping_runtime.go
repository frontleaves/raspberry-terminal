package runtime

import (
	"encoding/json"
	"github.com/XiaoLFeng/go-gin-util/blog"
	"raspberry-terminal/config/c"
	"time"
)

// PostDevicePing
//
// # 设备心跳
//
// 设备心跳，用于设备定时发送心跳信息。
func (r *Runtime) PostDevicePing() {
	addFunc, err := r.cron.AddFunc("@every 10s", func() {
		blog.Tracef("CRON", "向设备广播心跳")
		value := make(map[string]interface{})
		value["value"] = "ping"
		value["now"] = time.Now()
		marshal, err := json.Marshal(value)
		if err != nil {
			blog.Warnf("CRON", "消息序列化失败：%v", err.Error())
			return
		}
		subscribe := c.MqttClient.Subscribe(c.TopicPing, 0, nil)
		if subscribe.Wait() && subscribe.Error() != nil {
			blog.Warnf("CRON", "MQTT 订阅失败：%v", subscribe.Error())
		}
		token := c.MqttClient.Publish(c.TopicPing, 0, false, marshal)
		if token.Wait() && token.Error() != nil {
			blog.Warnf("CRON", "MQTT 发布失败：%v", token.Error())
		}
	})
	if err != nil {
		blog.Panicf("CRON", "添加定时任务失败：%v", err.Error())
	}
	r.funcID["PostDevicePing"] = addFunc
}
