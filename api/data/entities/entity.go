package entities

import (
	"github.com/rimdesk/product-api/api/data/domains"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	Type        string
	Name        string `gorm:"uniqueIndex:idx_name_barcode"`
	CategoryID  string
	Barcode     string `gorm:"uniqueIndex:idx_name_barcode"`
	Description string
	Amount      float32
	UnitPrice   float32
}

func (p *Product) ToDomain() domains.ProductDomain {
	return domains.ProductDomain{
		ID:          p.ID,
		Type:        p.Type,
		Name:        p.Name,
		CategoryID:  p.CategoryID,
		Barcode:     p.Barcode,
		Description: p.Description,
		Amount:      p.Amount,
		UnitPrice:   p.UnitPrice,
		CreatedAt:   &p.CreatedAt,
	}
}
