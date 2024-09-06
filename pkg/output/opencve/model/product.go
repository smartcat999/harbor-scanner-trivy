package model

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time  `gorm:"type:timestamptz;not null;index:ix_products_created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:timestamptz;not null" json:"updated_at"`
	Name      string     `gorm:"type:varchar;not null;index:ix_products_name" json:"name"`
	VendorID  *uuid.UUID `gorm:"type:uuid;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"vendor_id,omitempty"`

	// Define foreign key relationship with Vendor model
	Vendor *Vendor `gorm:"foreignKey:VendorID;references:ID" json:"vendor,omitempty"`
}

func (Product) TableName() string {
	return "products"
}
