package dto

import "time"

// MqttDevicePongDTO
//
// # MQTT 设备心跳 DTO
//
// 用于接收设备发送的心跳数据。
type MqttDevicePongDTO struct {
	Device string    `json:"device"`
	Pong   string    `json:"pong"`
	Now    time.Time `json:"now"`
}
