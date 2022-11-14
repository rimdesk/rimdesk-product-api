package product

import (
	"github.com/rimdesk/product-api/api/data/entities"
	"github.com/rimdesk/product-api/api/database"
)

func FindAll() []entities.Product {
	var products []entities.Product
	database.DB.Find(&products)
	return products
}

func FindById(id string) (*entities.Product, error) {
	var product entities.Product
	err := database.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func Create(product *entities.Product) (*entities.Product, error) {
	err := database.DB.Create(product).Error
	return product, err
}

func Update(product *entities.Product) (*entities.Product, error) {
	err := database.DB.Updates(product).Error
	return product, err
}

func Delete(product *entities.Product) error {
	err := database.DB.Delete(product).Error
	if err != nil {
		return err
	}
	return nil
}
