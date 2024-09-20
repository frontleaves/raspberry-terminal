package vo

// DeviceRegisterVO
//
// # 设备注册 VO
//
// 用于设备注册时的参数绑定。
//
// # 请求
//   - DeviceName: string 设备名称
type DeviceRegisterVO struct {
	DeviceName string `json:"device_name" form:"device_name" binding:"required"`
	Authorized bool   `json:"authorized" form:"authorized"`
}
