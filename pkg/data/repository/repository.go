package repository

import (
	"github.com/rimdesk/product-api/pkg/data/entities"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(string) ([]*entities.Product, error)
	FindById(id string) (*entities.Product, error)
	FindByCompanyIdAndId(string, string) (*entities.Product, error)
	Create(*entities.Product) error
	Update(*entities.Product) error
	Delete(*entities.Product) error
}

type productRepository struct {
	store *gorm.DB
}

func (repository *productRepository) FindByCompanyIdAndId(companyId string, id string) (*entities.Product, error) {
	var product entities.Product
	err := repository.store.Where("company_id = ?", companyId).First(&product, "id = ?", id).Error
	return &product, err
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{store: db}
}

func (repository *productRepository) FindAll(companyID string) ([]*entities.Product, error) {
	var products []*entities.Product
	err := repository.store.Where("company_id = ?", companyID).Find(&products).Error
	return products, err
}

func (repository *productRepository) FindById(id string) (*entities.Product, error) {
	var product entities.Product
	err := repository.store.First(&product, "id = ?", id).Error
	return &product, err
}

func (repository *productRepository) Create(product *entities.Product) error {
	return repository.store.Create(product).Error
}

func (repository *productRepository) Update(product *entities.Product) error {
	return repository.store.Updates(product).Error
}

func (repository *productRepository) Delete(product *entities.Product) error {
	return repository.store.Delete(product).Error
}
