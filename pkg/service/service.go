package service

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/rimdesk/product-api/pkg/clients"
	"github.com/rimdesk/product-api/pkg/data/domains"
	"github.com/rimdesk/product-api/pkg/data/dtos"
	"github.com/rimdesk/product-api/pkg/data/repository"
)

type ProductService interface {
	GetAllProducts(*fiber.Ctx, string) ([]*domains.ProductDomain, error)
	GetProductById(*fiber.Ctx, string) (*domains.ProductDomain, error)
	GetProductByCompanyAndId(*fiber.Ctx, string, string) (*domains.ProductDomain, error)
	CreateProduct(*fiber.Ctx, string, *dtos.ProductDto) (*domains.ProductDomain, error)
	UpdateProduct(*fiber.Ctx, string, string, *dtos.ProductDto) (*domains.ProductDomain, error)
	DeleteProduct(*fiber.Ctx, string, string) error
	SearchWarehouse(*fiber.Ctx, string, *dtos.ProductSearchDto) ([]*domains.ProductDomain, error)
}

type productService struct {
	productRepository repository.ProductRepository
	warehouseClient   clients.WarehouseClient
}

func (service *productService) GetProductByCompanyAndId(ctx *fiber.Ctx, s string, s2 string) (*domains.ProductDomain, error) {
	entity, err := service.productRepository.FindByCompanyIdAndId(s, s2)
	if err != nil {
		return nil, err
	}

	productDomain := entity.ToDomain()
	return productDomain, nil
}

func NewProductService(productRepository repository.ProductRepository, client clients.WarehouseClient) ProductService {
	return &productService{productRepository: productRepository, warehouseClient: client}
}

func (service *productService) GetAllProducts(ctx *fiber.Ctx, companyID string) ([]*domains.ProductDomain, error) {
	products, err := service.productRepository.FindAll(companyID)
	if err != nil {
		return nil, err
	}

	productDomains := make([]*domains.ProductDomain, 0)
	for _, entity := range products {
		productDomains = append(productDomains, entity.ToDomain())
	}

	return productDomains, nil
}

func (service *productService) GetProductById(ctx *fiber.Ctx, id string) (*domains.ProductDomain, error) {
	entity, err := service.productRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	productDomain := entity.ToDomain()
	return productDomain, nil
}

func (service *productService) CreateProduct(ctx *fiber.Ctx, companyID string, dto *dtos.ProductDto) (*domains.ProductDomain, error) {
	product := dto.ToEntity()
	product.CompanyID = companyID

	err := service.productRepository.Create(product)
	if err != nil {
		return nil, err
	}

	return product.ToDomain(), nil
}

func (service *productService) UpdateProduct(ctx *fiber.Ctx, companyID string, id string, dto *dtos.ProductDto) (*domains.ProductDomain, error) {
	product, err := service.productRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	dto.CopyToEntity(product)
	if err := service.productRepository.Update(product); err != nil {
		return nil, err
	}

	return product.ToDomain(), nil
}

func (service *productService) DeleteProduct(ctx *fiber.Ctx, companyID string, id string) error {
	entity, err := service.productRepository.FindById(id)
	if err != nil {
		return err
	}

	if companyID != entity.CompanyID {
		return errors.New("you are not the resource owner")
	}

	return service.productRepository.Delete(entity)
}

func (service *productService) SearchWarehouse(ctx *fiber.Ctx, companyID string, params *dtos.ProductSearchDto) ([]*domains.ProductDomain, error) {
	_, err := service.warehouseClient.GetById(ctx, params.WarehouseID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
