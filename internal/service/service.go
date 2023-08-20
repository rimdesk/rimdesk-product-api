package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rimdesk/product-api/internal/clients"
	"github.com/rimdesk/product-api/internal/data/domains"
	"github.com/rimdesk/product-api/internal/data/dtos"
	"github.com/rimdesk/product-api/internal/data/repository"
)

type ProductService interface {
	GetAllProducts(*fiber.Ctx) []*domains.ProductDomain
	GetProductById(*fiber.Ctx, string) (*domains.ProductDomain, error)
	CreateProduct(*fiber.Ctx, *dtos.ProductDto) (*domains.ProductDomain, error)
	UpdateProduct(*fiber.Ctx, string, *dtos.ProductDto) (*domains.ProductDomain, error)
	DeleteProduct(*fiber.Ctx, string) error
	SearchWarehouse(*fiber.Ctx, *dtos.ProductSearchDto) ([]domains.ProductDomain, error)
}

type productService struct {
	productRepository repository.ProductRepository
	warehouseClient   clients.WarehouseClient
}

func NewProductService(productRepository repository.ProductRepository, client clients.WarehouseClient) ProductService {
	return &productService{productRepository: productRepository, warehouseClient: client}
}

func (service *productService) GetAllProducts(ctx *fiber.Ctx) []*domains.ProductDomain {
	products := service.productRepository.FindAll()
	productDomains := make([]*domains.ProductDomain, 0)
	for _, entity := range products {
		productDomains = append(productDomains, entity.ToDomain())
	}

	return productDomains
}

func (service *productService) GetProductById(ctx *fiber.Ctx, id string) (*domains.ProductDomain, error) {
	entity, err := service.productRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	productDomain := entity.ToDomain()
	return productDomain, nil
}

func (service *productService) CreateProduct(ctx *fiber.Ctx, dto *dtos.ProductDto) (*domains.ProductDomain, error) {
	entity := dto.ToEntity()
	productEntity, err := service.productRepository.Create(entity)
	if err != nil {
		return nil, err
	}

	productDomain := productEntity.ToDomain()
	return productDomain, nil
}

func (service *productService) UpdateProduct(ctx *fiber.Ctx, id string, dto *dtos.ProductDto) (*domains.ProductDomain, error) {
	entity, err := service.productRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	dto.CopyToEntity(entity)
	productEntity, err := service.productRepository.Update(entity)
	if err != nil {
		return nil, err
	}

	productDomain := productEntity.ToDomain()

	return productDomain, nil
}

func (service *productService) DeleteProduct(ctx *fiber.Ctx, id string) error {
	entity, err := service.productRepository.FindById(id)
	if err != nil {
		return err
	}

	return service.productRepository.Delete(entity)
}

func (service *productService) SearchWarehouse(ctx *fiber.Ctx, params *dtos.ProductSearchDto) ([]domains.ProductDomain, error) {
	_, err := service.warehouseClient.GetById(ctx, params.WarehouseID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
