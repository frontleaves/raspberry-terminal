package entity

import (
	"github.com/XiaoLFeng/go-general-utils/butil"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Device
//
// # 设备信息
//
// 设备信息实体，用于描述设备的基本信息。
//
// # 参数
//   - UUID: 设备唯一标识
//   - Type: 设备类型
//   - DeviceName: 设备名称
//   - DeviceUsername: 设备用户名
//   - DevicePassword: 设备密码
//   - DeviceHost: 设备主机
//   - DeviceMac: 设备MAC地址
//   - Authorized: 设备是否授权
//   - CreatedAt: 创建时间
type Device struct {
	UUID           uuid.UUID `json:"uuid" gorm:"type:varchar(36);primaryKey;not null"`
	Type           string    `json:"type" gorm:"type:varchar(32);not null"`
	DeviceName     string    `json:"device_name" gorm:"type:varchar(255);not null"`
	DeviceUsername string    `json:"device_username" gorm:"type:varchar(255);not null"`
	DevicePassword string    `json:"device_password" gorm:"type:varchar(255);not null"`
	DeviceHost     string    `json:"device_host" gorm:"type:varchar(255);not null"`
	DeviceMac      string    `json:"device_mac" gorm:"type:varchar(255);not null"`
	Authorized     bool      `json:"authorized" gorm:"type:boolean;not null;default:false"`
	CreatedAt      time.Time `json:"created_at" gorm:"type:timestamp;not null;default:current_timestamp"`
	Login          bool      `json:"login" gorm:"type:boolean;not null;default:false"`
	NowValue       string    `json:"now_value" gorm:"type:json"`
}

func (device *Device) BeforeCreate(_ *gorm.DB) error {
	if device.UUID == uuid.Nil {
		device.UUID = butil.GenerateUUID()
	}
	return nil
}

func (device *Device) BeforeUpdate(_ *gorm.DB) error {
	device.CreatedAt = time.Now()
	return nil
}
