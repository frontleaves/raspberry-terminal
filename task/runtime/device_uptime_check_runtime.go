package runtime

import (
	"github.com/XiaoLFeng/go-gin-util/blog"
	"raspberry-terminal/config/c"
	"raspberry-terminal/model/entity"
	"time"
)

// DeviceUptimeCheck
//
// # 设备在线检查
//
// 设备在线检查，用于定时检查设备是否在线。
func (r *Runtime) DeviceUptimeCheck() {
	addFunc, err := r.cron.AddFunc("@every 30s", func() {
		blog.Tracef("CRON", "设备在线检查")
		var getAllDevice *[]entity.Device
		c.DB.Find(&getAllDevice)
		for _, device := range *getAllDevice {
			// 如果 Uptime 多于 30s 设定为下线
			if time.Now().Sub(device.Uptime).Seconds() > 30 {
				device.Login = false
				c.DB.Save(&device)
			}
		}
	})
	if err != nil {
		blog.Panicf("CRON", "添加定时任务失败：%v", err.Error())
	}
	r.funcID["DeviceUptimeCheck"] = addFunc
}
