package dto

// WsSystemInfoDTO
//
// # WebSocket 系统信息实体
//
// 用于 WebSocket 系统信息实体，用于传输系统信息。
//
// # 属性
//   - Cpu: string CPU 信息
//   - CpuTemperature: string CPU 温度
//   - Ram: string 内存信息
//   - Disk: string 磁盘信息
type WsSystemInfoDTO struct {
	Cpu            float64 `json:"cpu"`
	CpuTemperature string  `json:"cpu_temperature"`
	Ram            float64 `json:"ram"`
	Disk           float64 `json:"disk"`
}
