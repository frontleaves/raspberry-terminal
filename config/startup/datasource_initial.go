package startup

import (
	"github.com/XiaoLFeng/go-gin-util/blog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"raspberry-terminal/config/c"
	"raspberry-terminal/model/entity"
)

// datasource
//
// # 数据源初始化
//
// 数据源初始化，用于初始化数据库连接。
func (i *InitStr) datasource() {
	blog.Infof("INIT", "数据库初始化中...")
	db, err := gorm.Open(sqlite.Open("raspberry-terminal.db"), &gorm.Config{})
	if err != nil {
		blog.Panicf("INIT", "数据库连接失败：%v", err.Error())
	}
	err = db.AutoMigrate(&entity.Device{}, &entity.Log{})
	if err != nil {
		blog.Panicf("INIT", "数据表创建(检查)失败： %v", err.Error())
	}
	c.DB = db
}
