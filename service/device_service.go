package service

import (
	"encoding/json"
	"github.com/XiaoLFeng/go-general-utils/bcode"
	"github.com/XiaoLFeng/go-gin-util/berror"
	"github.com/XiaoLFeng/go-gin-util/blog"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"raspberry-terminal/config/c"
	"raspberry-terminal/model/dto"
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

// DeviceRegister
//
// # 设备注册
//
// 设备注册，用于设备注册；
// 设备在注册后，将会在本服务内进行显示。
//
// # 请求
//   - getVO: vo.DeviceRegisterVO 设备注册参数
//
// # 返回
//   - error: error 错误信息
func DeviceRegister(getVO vo.DeviceRegisterVO) error {
	blog.Info("SERV", "DeviceRegister - 设备注册")
	var device *entity.Device
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		tx.Last(&device, "device_name = ?", getVO.DeviceName)
		if device.UUID == uuid.Nil {
			return berror.New(bcode.NotFound, "设备未找到")
		}
		device.Authorized = getVO.Authorized
		tx.Save(&device)
		// 向 MQTT 服务发送授权信息
		authReturnDTO := &dto.MqttAuthReturnDTO{
			Device:     device.DeviceName,
			Authorized: device.Authorized,
		}
		marshal, err := json.Marshal(authReturnDTO)
		if err != nil {
			return berror.New(bcode.ServerInternalError, "消息序列化失败："+err.Error())
		}
		publish := c.MqttClient.Publish(c.TopicAuthReturn, 0, false, marshal)
		if publish.Wait() && publish.Error() != nil {
			blog.Warnf("SERV", "MQTT 发送失败：%v", publish.Error())
			return publish.Error()
		}
		return nil
	})
	return err
}

// DeviceLight
//
// # 设备灯光控制
//
// 设备灯光控制，用于设备灯光控制；
// 设备在注册后，将会在本服务内进行显示；
// 本接口为设备灯光控制。
//
// # 请求
//   - getVO: vo.DeviceLightVO 设备灯光控制参数
//
// # 返回
//   - error: error 错误信息
func DeviceLight(getVO vo.DeviceLightVO) error {
	blog.Info("SERV", "DeviceLight - 设备灯光控制")
	var device *entity.Device
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		tx.Last(&device, "device_name = ?", getVO.Device)
		if device.UUID == uuid.Nil {
			return berror.New(bcode.NotFound, "设备未找到")
		}
		// 检查是否授权
		if !device.Authorized {
			return berror.New(bcode.Unauthorized, "设备未授权")
		}
		marshal, err := json.Marshal(getVO)
		if err != nil {
			return berror.New(bcode.ServerInternalError, "消息序列化失败："+err.Error())
		}
		device.NowValue = string(marshal)
		tx.Save(&device)
		// 发送设备控制指令
		publish := c.MqttClient.Publish(c.TopicOperateReturn, 0, false, marshal)
		if publish.Wait() && publish.Error() != nil {
			blog.Warnf("SERV", "MQTT 发送失败：%v", publish.Error())
			return publish.Error()
		}
		return nil
	})
	return err
}

// DeviceGetOnline
//
// # 获取在线设备
//
// 获取在线设备，用于获取在线设备列表；
// 设备在注册后，将会在本服务内进行显示；
// 本接口为获取在线设备列表。
//
// # 请求
//   - getVO: vo.CustomPageVO 分页参数
//
// # 返回
//   - drives: *[]entity.Device 在线设备列表
func DeviceGetOnline(getVO vo.CustomPageVO) (devices *[]entity.Device) {
	blog.Info("SERV", "DeviceGetOnline - 获取在线设备")
	var getOnlineDevice *[]entity.Device
	if getVO.Search == "" {
		c.DB.Find(&getOnlineDevice, "login = ?", true).Limit(getVO.Limit).Offset(getVO.Page)
	} else {
		c.DB.Find(&getOnlineDevice, "login = ?", true).Where("device_name LIKE ?", "%"+getVO.Search+"%")
	}
	blog.Tracef("SERV", "\t> 取得数据：%v", getOnlineDevice)
	if len(*getOnlineDevice) == 0 {
		return new([]entity.Device)
	}
	return getOnlineDevice
}

// DeviceGetOffline
//
// # 获取离线设备
//
// 获取离线设备，用于获取离线设备列表；
// 设备在注册后，将会在本服务内进行显示；
// 本接口为获取离线设备列表。
//
// # 请求
//   - getVO: vo.CustomPageVO 分页参数
//
// # 返回
//   - drives: *[]entity.Device 离线设备列表
func DeviceGetOffline(getVO vo.CustomPageVO) (devices *[]entity.Device) {
	blog.Info("SERV", "DeviceGetOffline - 获取离线设备")
	var getOfflineDevice *[]entity.Device
	if getVO.Search == "" {
		c.DB.Find(&getOfflineDevice, "login = ?", false).Limit(getVO.Limit).Offset(getVO.Page)
	} else {
		c.DB.Find(&getOfflineDevice, "login = ?", false).Where("device_name LIKE ?", "%"+getVO.Search+"%")
	}
	blog.Tracef("SERV", "\t> 取得数据：%v", getOfflineDevice)
	if len(*getOfflineDevice) == 0 {
		return new([]entity.Device)
	}
	return getOfflineDevice
}
