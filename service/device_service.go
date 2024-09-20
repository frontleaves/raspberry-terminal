package service

import (
	"github.com/XiaoLFeng/go-gin-util/blog"
	"raspberry-terminal/config/c"
	"raspberry-terminal/model/entity"
	"raspberry-terminal/model/vo"
)

// DeviceGetNoRegister
//
// # 获取未注册设备
//
// 获取未登录设备，用于获取未登录设备列表；
// 设备在未注册时，链接 MQTT 服务后，将会在本服务内进行自动注册，并且显示出未注册设备列表；
// 设备需要进行注册后，才能进行后续操作；
// 本接口为获取未注册设备列表。
//
// # 请求
//   - getVO: vo.CustomPageVO 分页参数
//
// # 返回
//   - drives: *[]entity.Device 未注册设备列表
func DeviceGetNoRegister(getVO vo.CustomPageVO) (drives *[]entity.Device) {
	blog.Info("SERV", "DeviceGetNoRegister - 获取未注册设备")
	var getNoRegisterDrive *[]entity.Device
	if getVO.Search == "" {
		c.DB.Find(&getNoRegisterDrive, "authorized = ?", false).Limit(getVO.Limit).Offset(getVO.Page)
	} else {
		c.DB.Find(&getNoRegisterDrive, "authorized = ?", false).Where("device_name LIKE ?", "%"+getVO.Search+"%")
	}
	blog.Tracef("SERV", "\t> 取得数据：%v", getNoRegisterDrive)
	if len(*getNoRegisterDrive) == 0 {
		return new([]entity.Device)
	}
	return getNoRegisterDrive
}

// DeviceGet
//
// # 获取设备列表
//
// 获取设备列表，用于获取设备列表；
// 设备在注册后，将会在本服务内进行显示；
// 本接口为获取设备列表。
//
// # 请求
//   - getVO: vo.CustomPageVO 分页参数
//
// # 返回
//   - drives: *[]entity.Device 设备列表
func DeviceGet(getVO vo.CustomPageVO) (drives *[]entity.Device) {
	blog.Info("SERV", "DeviceGet - 获取设备列表")
	var getDrive *[]entity.Device
	if getVO.Search == "" {
		c.DB.Find(&getDrive, "authorized = ?", true).Limit(getVO.Limit).Offset(getVO.Page)
	} else {
		c.DB.Find(&getDrive, "authorized = ?", true).Where("device_name LIKE ?", "%"+getVO.Search+"%")
	}
	blog.Tracef("SERV", "\t> 取得数据：%v", getDrive)
	if len(*getDrive) == 0 {
		return new([]entity.Device)
	}
	return getDrive
}
