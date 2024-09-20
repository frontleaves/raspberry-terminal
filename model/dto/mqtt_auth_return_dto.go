package dto

// MqttAuthReturnDTO
//
// # Mqtt设备授权返回信息
//
// Mqtt设备授权返回信息，返回Mqtt设备授权返回信息。
//
// # 参数
//   - Device: string, 设备ID
//   - Authorized: bool, 是否授权
type MqttAuthReturnDTO struct {
	Device     string `json:"device"`
	Authorized bool   `json:"authorized"`
}
