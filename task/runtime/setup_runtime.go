package runtime

import (
	"github.com/XiaoLFeng/go-gin-util/blog"
	"github.com/robfig/cron/v3"
)

type Runtime struct {
	cron   *cron.Cron
	funcID map[string]cron.EntryID
}

func New(c *cron.Cron) *Runtime {
	return &Runtime{
		cron:   c,
		funcID: make(map[string]cron.EntryID),
	}
}

// SetupRuntime
//
// # 运行时初始化
//
// 运行时初始化，用于初始化定时任务。
func SetupRuntime() {
	runtime := New(cron.New())

	runtime.PostDevicePing()
	runtime.DeviceUptimeCheck()

	runtime.cron.Start()
	blog.Infof("CRON", "运行时启动成功")
}
