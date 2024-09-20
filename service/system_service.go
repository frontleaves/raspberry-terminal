package service

import (
	"bytes"
	"github.com/XiaoLFeng/go-gin-util/blog"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

// GetCpuPercent
//
// # 获取CPU使用率
//
// 获取CPU使用率，返回值为 float64 类型的百分比。
//
// # 返回
//   - float64 CPU使用率
//   - error 异常信息
func GetCpuPercent() (float64, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		blog.Errorf("SERV", "获取 CPU 温度异常：%v", err.Error())
		return 0, err
	}
	return percent[0], nil
}

// GetCpuTemp
//
// # 获取CPU温度
//
// 获取CPU温度，返回值为 string 类型的温度值。
//
// # 返回
//   - string CPU温度
//   - error 异常信息
func GetCpuTemp() (string, error) {
	cmd := exec.Command("cat", "/sys/class/thermal/thermal_zone0/temp")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		blog.Errorf("SERV", "获取 CPU 温度异常：%v", err.Error())
		return "", err
	}
	blog.Debugf("SERV", "CPU温度：%v", out.String())
	tempStr := strings.Replace(out.String(), "\n", "", -1)
	temp, err := strconv.Atoi(tempStr)
	if err != nil {
		blog.Errorf("SERV", "CPU温度转换异常：%v", err.Error())
		return "", err
	}
	temp = temp / 1000
	blog.Debugf("SERV", "CPU温度：%v", temp)
	return strconv.Itoa(temp), nil
}

// GetRamPercent
//
// # 获取内存使用率
//
// 获取内存使用率，返回值为 float64 类型的百分比。
//
// # 返回
//   - float64 内存使用率
//   - error 异常信息
func GetRamPercent() (float64, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		blog.Errorf("SERV", "获取内存使用率异常：%v", err.Error())
		return 0, err
	}
	return memInfo.UsedPercent, nil
}

// GetDiskPercent
//
// # 获取磁盘使用率
//
// 获取磁盘使用率，返回值为 float64 类型的百分比。
//
// # 返回
//   - float64 磁盘使用率
//   - error 异常信息
func GetDiskPercent() (float64, error) {
	parts, err := disk.Partitions(true)
	if err != nil {
		blog.Errorf("SERV", "获取磁盘使用率异常：%v", err.Error())
		return 0, err
	}
	diskInfo, err := disk.Usage(parts[0].Mountpoint)
	if err != nil {
		blog.Errorf("SERV", "获取磁盘使用率异常：%v", err.Error())
		return 0, err
	}
	return diskInfo.UsedPercent, nil
}
