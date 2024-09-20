package entity

import (
	"github.com/XiaoLFeng/go-general-utils/butil"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Log
//
// # 日志信息
//
// 日志信息实体，用于描述日志的基本信息。
//
// # 参数
//   - UUID: 日志唯一标识
//   - Operate: 操作
//   - Content: 内容
//   - CreatedAt: 创建时间
type Log struct {
	LogUUID   uuid.UUID `json:"log_uuid" gorm:"type:varchar(36);primaryKey;not null"`
	Operate   string    `json:"operate" gorm:"type:varchar(64);not null"`
	UUID      uuid.UUID `json:"uuid" gorm:"type:varchar(36);not null"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;not null;default:current_timestamp"`
}

func (log *Log) BeforeCreate(_ *gorm.DB) error {
	if log.LogUUID == uuid.Nil {
		log.LogUUID = butil.GenerateUUID()
	}
	if log.CreatedAt.IsZero() {
		log.CreatedAt = time.Now()
	}
	return nil
}
