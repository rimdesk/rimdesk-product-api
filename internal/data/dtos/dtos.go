package dtos

import (
	"github.com/jinzhu/copier"
	"github.com/rimdesk/product-api/internal/data/entities"
	"log"
)

type ProductDto struct {
	Name        string  `json:"name" validate:"required"`
	Type        string  `json:"type" validate:"required"`
	CompanyID   string  `json:"company_id" validate:"required"`
	CategoryID  string  `json:"category_id" validate:"required"`
	Barcode     string  `json:"barcode" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Amount      float64 `json:"amount" validate:"required"`
	SupplyPrice float64 `json:"supply_price" validate:"required"`
	RetailPrice float64 `json:"retail_price" validate:"required"`
}

type ProductSearchDto struct {
	WarehouseID string `json:"warehouse_id" validate:"required"`
	Query       string `json:"query" validate:"required"`
}

func (d *ProductDto) ToEntity() *entities.Product {
	product := new(entities.Product)
	if err := copier.Copy(product, d); err != nil {
		log.Println("failed to copy entity:", err)
	}

	return product
}

func (d *ProductDto) CopyToEntity(product *entities.Product) {
	if err := copier.Copy(product, d); err != nil {
		log.Println("failed to copy entity:", err)
	}
}
