package product

import (
	"github.com/rimdesk/product-api/api/data/domains"
	productRepository "github.com/rimdesk/product-api/api/data/repository/product"
	"github.com/rimdesk/product-api/api/http/request/product"
)

func GetAllProducts() []domains.ProductDomain {
	products := productRepository.FindAll()
	productDomains := make([]domains.ProductDomain, 0)
	for _, entity := range products {
		productDomains = append(productDomains, entity.ToDomain())
	}

	return productDomains
}

func GetProductById(id string) (*domains.ProductDomain, error) {
	entity, err := productRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	productDomain := entity.ToDomain()
	return &productDomain, nil
}

func CreateProduct(dto *product.Dto) (*domains.ProductDomain, error) {
	entity := dto.ToEntity()
	productEntity, err := productRepository.Create(&entity)
	if err != nil {
		return nil, err
	}

	productDomain := productEntity.ToDomain()
	return &productDomain, nil
}

func UpdateProduct(id string, dto *product.Dto) (*domains.ProductDomain, error) {
	entity, err := productRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	dto.CopyToEntity(entity)
	productEntity, err := productRepository.Update(entity)
	if err != nil {
		return nil, err
	}

	productDomain := productEntity.ToDomain()

	return &productDomain, nil
}

func DeleteProduct(id string) error {
	entity, err := productRepository.FindById(id)
	if err != nil {
		return err
	}

	return productRepository.Delete(entity)
}
