package repo

import (
	"github.com/aquasecurity/harbor-scanner-trivy/pkg/output/opencve/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindByID(id uuid.UUID) (*model.Product, error)
	FindAll() ([]model.Product, error)
	Create(vendor *model.Product) error
	Update(vendor *model.Product) error
	Delete(id uuid.UUID) error
}

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepositoryImpl{db: db}
}

func (p *productRepositoryImpl) FindByID(id uuid.UUID) (*model.Product, error) {
	var product model.Product
	if err := p.db.First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *productRepositoryImpl) FindAll() ([]model.Product, error) {
	var products []model.Product
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *productRepositoryImpl) Create(product *model.Product) error {
	return p.db.Create(product).Error
}

func (p *productRepositoryImpl) Update(product *model.Product) error {
	return p.db.Save(product).Error
}

func (p *productRepositoryImpl) Delete(id uuid.UUID) error {
	return p.db.Delete(&model.Product{}, id).Error
}
