package dto

// MqttAuthDTO
//
// # Mqtt认证信息
//
// Mqtt认证信息，返回Mqtt认证信息。
//
// # 参数
//   - Type: string, 认证类型
//   - Device: string, 设备ID
//   - Username: string, 用户名
//   - Password: string, 密码
//   - Host: string, 主机
//   - Mac: string, Mac地址
type MqttAuthDTO struct {
	Type     string `json:"type"`
	Device   string `json:"device"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Mac      string `json:"mac"`
}
