package model

import (
	"github.com/google/uuid"
	"time"
)

type Vendor struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"type:timestamptz;not null;index:ix_vendors_created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamptz;not null" json:"updated_at"`
	Name      string    `gorm:"type:varchar;not null;uniqueIndex:vendors_name_key" json:"name"`

	// One-to-many relationship with products
	Products []Product `gorm:"foreignKey:VendorID" json:"products,omitempty"`
}

func (Vendor) TableName() string {
	return "vendors"
}
