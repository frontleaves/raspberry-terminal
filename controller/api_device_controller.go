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
