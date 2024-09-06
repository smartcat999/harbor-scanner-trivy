package repo

import (
	"github.com/aquasecurity/harbor-scanner-trivy/pkg/output/opencve/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CveRepository interface {
	FindByID(id uuid.UUID) (*model.CVE, error)
	FindAll() ([]model.CVE, error)
	Create(vendor *model.CVE) error
	Update(vendor *model.CVE) error
	Delete(id uuid.UUID) error
}

type cveRepositoryImpl struct {
	db *gorm.DB
}

func NewCveRepository(db *gorm.DB) CveRepository {
	return &cveRepositoryImpl{db: db}
}

func (c *cveRepositoryImpl) FindByID(id uuid.UUID) (*model.CVE, error) {
	var cve model.CVE
	if err := c.db.First(&cve, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &cve, nil
}

func (c *cveRepositoryImpl) FindAll() ([]model.CVE, error) {
	var cves []model.CVE
	if err := c.db.Find(&cves).Error; err != nil {
		return nil, err
	}
	return cves, nil
}

func (c *cveRepositoryImpl) Create(cve *model.CVE) error {
	return c.db.Create(cve).Error
}

func (c *cveRepositoryImpl) Update(cve *model.CVE) error {
	return c.db.Save(cve).Error
}

func (c *cveRepositoryImpl) Delete(id uuid.UUID) error {
	return c.db.Delete(&model.CVE{}, id).Error
}
