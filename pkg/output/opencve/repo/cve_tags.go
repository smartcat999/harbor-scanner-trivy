package repo

import (
	"github.com/aquasecurity/harbor-scanner-trivy/pkg/output/opencve/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CVETagRepository interface {
	FindByID(id uuid.UUID) (*model.CVETag, error)
	FindAll() ([]model.CVETag, error)
	Create(tag *model.CVETag) error
	CreateMany(tags []model.CVETag) error
	Update(tag *model.CVETag) error
	Delete(id uuid.UUID) error
}

type cveTagRepositoryImpl struct {
	db *gorm.DB
}

func NewCVETagRepository(db *gorm.DB) CVETagRepository {
	return &cveTagRepositoryImpl{db: db}
}

func (r *cveTagRepositoryImpl) FindByID(id uuid.UUID) (*model.CVETag, error) {
	var tag model.CVETag
	if err := r.db.First(&tag, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *cveTagRepositoryImpl) FindAll() ([]model.CVETag, error) {
	var tags []model.CVETag
	if err := r.db.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *cveTagRepositoryImpl) Create(tag *model.CVETag) error {
	return r.db.Create(tag).Error
}

func (r *cveTagRepositoryImpl) CreateMany(tags []model.CVETag) error {
	return r.db.CreateInBatches(tags, 50).Error
}

func (r *cveTagRepositoryImpl) Update(tag *model.CVETag) error {
	return r.db.Save(tag).Error
}

func (r *cveTagRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&model.CVETag{}, id).Error
}
