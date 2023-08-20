package repository

import (
	"github.com/rimdesk/product-api/internal/data/entities"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() []entities.Product
	FindById(id string) (*entities.Product, error)
	Create(product *entities.Product) (*entities.Product, error)
	Update(product *entities.Product) (*entities.Product, error)
	Delete(product *entities.Product) error
}

type productRepository struct {
	store *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{store: db}
}

func (repository *productRepository) FindAll() []entities.Product {
	var products []entities.Product
	repository.store.Find(&products)
	return products
}

func (repository *productRepository) FindById(id string) (*entities.Product, error) {
	var product entities.Product
	err := repository.store.First(&product, "id = ?", id).Error
	return &product, err
}

func (repository *productRepository) Create(product *entities.Product) (*entities.Product, error) {
	err := repository.store.Create(product).Error
	return product, err
}

func (repository *productRepository) Update(product *entities.Product) (*entities.Product, error) {
	err := repository.store.Updates(product).Error
	return product, err
}

func (repository *productRepository) Delete(product *entities.Product) error {
	err := repository.store.Delete(product).Error
	if err != nil {
		return err
	}
	return nil
}
