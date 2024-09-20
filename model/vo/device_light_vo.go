package vo

// DeviceLightVO
//
// # 设备灯光控制 VO
//
// 用于设备灯光控制时的参数绑定。
//
// # 请求
//   - Device: string 设备名称
//   - Value: bool 灯光状态
type DeviceLightVO struct {
	Device string `json:"device" form:"device" binding:"required"`
	Value  bool   `json:"value" form:"value"`
}
