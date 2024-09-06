package model

import (
	"github.com/google/uuid"
	"time"
)

type CVETag struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamptz;not null" json:"updated_at"`
	Tags      []string  `gorm:"type:jsonb" json:"tags"` // 将 jsonb 字段映射为 Go 的字符串切片
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id"`
	CVEID     uuid.UUID `gorm:"type:uuid" json:"cve_id"`
}

// TableName 显式定义表名为 "cves_tags"
func (CVETag) TableName() string {
	return "cves_tags"
}
