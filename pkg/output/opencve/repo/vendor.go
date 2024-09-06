package repo

import (
	"github.com/aquasecurity/harbor-scanner-trivy/pkg/output/opencve/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VendorRepository interface {
	FindByID(id uuid.UUID) (*model.Vendor, error)
	FindAll() ([]model.Vendor, error)
	Create(vendor *model.Vendor) error
	Update(vendor *model.Vendor) error
	Delete(id uuid.UUID) error
}

type vendorRepositoryImpl struct {
	db *gorm.DB
}

func NewVendorRepository(db *gorm.DB) VendorRepository {
	return &vendorRepositoryImpl{db: db}
}

func (r *vendorRepositoryImpl) FindByID(id uuid.UUID) (*model.Vendor, error) {
	var vendor model.Vendor
	if err := r.db.First(&vendor, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &vendor, nil
}

func (r *vendorRepositoryImpl) FindAll() ([]model.Vendor, error) {
	var vendors []model.Vendor
	if err := r.db.Find(&vendors).Error; err != nil {
		return nil, err
	}
	return vendors, nil
}

func (r *vendorRepositoryImpl) Create(vendor *model.Vendor) error {
	return r.db.Create(vendor).Error
}

func (r *vendorRepositoryImpl) Update(vendor *model.Vendor) error {
	return r.db.Save(vendor).Error
}

func (r *vendorRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&model.Vendor{}, id).Error
}
