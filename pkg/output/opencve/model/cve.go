package model

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"time"
)

type CVE struct {
	ID        uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time    `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt time.Time    `gorm:"type:timestamptz;not null" json:"updated_at"`
	CveID     string       `gorm:"type:varchar;not null" json:"cve_id"`
	Json      pgtype.JSONB `gorm:"type:jsonb" json:"json,omitempty"`
	Vendors   pgtype.JSONB `gorm:"type:jsonb" json:"vendors,omitempty"`
	Cwes      pgtype.JSONB `gorm:"type:jsonb" json:"cwes,omitempty"`
	Summary   string       `gorm:"type:varchar;not null" json:"summary"`
	Cvss2     float64      `gorm:"type:double precision" json:"cvss2,omitempty"`
	Cvss3     float64      `gorm:"type:double precision" json:"cvss3,omitempty"`
}

func (CVE) TableName() string {
	return "cves"
}
