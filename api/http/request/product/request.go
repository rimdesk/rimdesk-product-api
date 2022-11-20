package product

import (
	"github.com/google/uuid"
	"github.com/rimdesk/product-api/api/data/entities"
)

type Dto struct {
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	CategoryID  string  `json:"category_id"`
	Barcode     string  `json:"barcode"`
	Description string  `json:"description"`
	Amount      float32 `json:"amount"`
	UnitPrice   float32 `json:"unitPrice"`
	RetailPrice float32 `json:"retailPrice"`
}

func (d *Dto) ToEntity() entities.Product {
	return entities.Product{
		ID:          uuid.NewString(),
		Type:        d.Type,
		Name:        d.Name,
		CategoryID:  d.CategoryID,
		Barcode:     d.Barcode,
		Description: d.Description,
		Amount:      d.Amount,
		UnitPrice:   d.UnitPrice,
		RetailPrice: d.RetailPrice,
	}
}

func (d *Dto) CopyToEntity(product *entities.Product) {
	product.Name = d.Name
	product.Amount = d.Amount
	product.Type = d.Type
	product.Barcode = d.Barcode
	product.UnitPrice = d.UnitPrice
	product.RetailPrice = d.RetailPrice
	product.Amount = d.Amount
	product.Description = d.Description
	product.CategoryID = d.CategoryID
}
