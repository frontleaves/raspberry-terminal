package runtime

import (
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
		c.MqttClient.Publish(c.TopicPing, 0, false, value)
	})
	if err != nil {
		blog.Panicf("CRON", "添加定时任务失败：%v", err.Error())
	}
	r.funcID["PostDevicePing"] = addFunc
}
