package dto

// MqttDeviceOperateDTO
//
// # Mqtt设备操作信息
//
// Mqtt设备操作信息，返回Mqtt设备操作信息。
//
// # 参数
//   - Device: string, 设备ID
//   - Type: string, 操作类型
//   - Data: interface{}, 操作数据
type MqttDeviceOperateDTO struct {
	Device string      `json:"device"`
	Value  interface{} `json:"value"`
}
