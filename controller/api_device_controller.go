package controller

import (
	"github.com/XiaoLFeng/go-general-utils/bcode"
	"github.com/XiaoLFeng/go-gin-util/berror"
	"github.com/XiaoLFeng/go-gin-util/blog"
	"github.com/XiaoLFeng/go-gin-util/bresult"
	"github.com/gin-gonic/gin"
	"raspberry-terminal/model/vo"
	"raspberry-terminal/service"
)

// ApiGetNoRegisterDeviceController
//
// # 获取未注册设备
//
// 获取未登录设备，用于获取未登录设备列表；
// 设备在未注册时，链接 MQTT 服务后，将会在本服务内进行自动注册，并且显示出未注册设备列表；
// 设备需要进行注册后，才能进行后续操作；
// 本接口为获取未注册设备列表。
//
// # 请求
//   - c: gin.Context Gin 上下文
func ApiGetNoRegisterDeviceController(c *gin.Context) {
	blog.Info("CONT", "ApiGetNoRegisterDeviceController - 获取未注册设备")
	var getVO vo.CustomPageVO
	if err := c.ShouldBindQuery(&getVO); err != nil {
		_ = c.Error(berror.New(bcode.BadRequestInvalidParameter, "参数绑定错误："+err.Error()))
		return
	}
	getDrives := service.DeviceGetNoRegister(getVO)
	bresult.OkWithData(c, "获取未注册设备成功", getDrives)
}

// ApiGetDeviceController
//
// # 获取设备列表
//
// 获取设备列表，用于获取设备列表；
// 设备在注册后，将会在本服务内进行显示；
// 本接口为获取设备列表。
//
// # 请求
//   - c: gin.Context Gin 上下文
func ApiGetDeviceController(c *gin.Context) {
	blog.Info("CONT", "ApiGetDeviceController - 获取设备列表")
	var getVO vo.CustomPageVO
	if err := c.ShouldBindQuery(&getVO); err != nil {
		_ = c.Error(berror.New(bcode.BadRequestInvalidParameter, "参数绑定错误："+err.Error()))
		return
	}
	getDrives := service.DeviceGet(getVO)
	bresult.OkWithData(c, "获取设备成功", getDrives)
}

// ApiGetOnlineDeviceController
//
// # 获取在线设备
//
// 获取在线设备，用于获取在线设备列表；
// 设备在注册后，将会在本服务内进行显示；
// 本接口为获取在线设备。
//
// # 请求
//   - c: gin.Context Gin 上下文
func ApiGetOnlineDeviceController(c *gin.Context) {
	blog.Info("CONT", "ApiGetOnlineDeviceController - 获取在线设备")
	var getVO vo.CustomPageVO
	if err := c.ShouldBindQuery(&getVO); err != nil {
		_ = c.Error(berror.New(bcode.BadRequestInvalidParameter, "参数绑定错误："+err.Error()))
		return
	}
	getDrives := service.DeviceGetOnline(getVO)
	bresult.OkWithData(c, "获取在线设备成功", getDrives)
}

func ApiGetOfflineDeviceController(c *gin.Context) {
	blog.Info("CONT", "ApiGetOfflineDeviceController - 获取离线设备")
	var getVO vo.CustomPageVO
	if err := c.ShouldBindQuery(&getVO); err != nil {
		_ = c.Error(berror.New(bcode.BadRequestInvalidParameter, "参数绑定错误："+err.Error()))
		return
	}
	getDrives := service.DeviceGetOffline(getVO)
	bresult.OkWithData(c, "获取离线设备成功", getDrives)
}

// ApiRegisterDeviceController
//
// # 注册设备
//
// 注册设备，用于注册设备；
// 设备在未注册时，链接 MQTT 服务后，将会在本服务内进行自动注册；
// 本接口为注册设备。
//
// # 请求
//   - c: gin.Context Gin 上下文
func ApiRegisterDeviceController(c *gin.Context) {
	blog.Info("CONT", "ApiRegisterDeviceController - 注册设备")
	var registerVO vo.DeviceRegisterVO
	if err := c.ShouldBindJSON(&registerVO); err != nil {
		_ = c.Error(berror.New(bcode.BadRequestInvalidParameter, "参数绑定错误："+err.Error()))
		return
	}
	err := service.DeviceRegister(registerVO)
	if err != nil {
		_ = c.Error(err)
		return
	}
	bresult.Ok(c, "注册设备成功")
}

// ApiDeviceLightController
//
// # 设备灯光控制
//
// 设备灯光控制，用于设备灯光控制；
// 设备在注册后，将会在本服务内进行显示；
// 本接口为设备灯光控制。
//
// # 请求
//   - c: gin.Context Gin 上下文
func ApiDeviceLightController(c *gin.Context) {
	blog.Info("CONT", "ApiDeviceLightController - 设备灯光控制")
	var lightVO vo.DeviceLightVO
	if err := c.ShouldBindJSON(&lightVO); err != nil {
		_ = c.Error(berror.New(bcode.BadRequestInvalidParameter, "参数绑定错误："+err.Error()))
		return
	}
	err := service.DeviceLight(lightVO)
	if err != nil {
		_ = c.Error(err)
		return
	}
	bresult.Ok(c, "设备灯光成功")
}
