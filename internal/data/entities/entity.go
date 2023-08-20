package entities

import (
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/rimdesk/product-api/internal/data/domains"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	CompanyID   string
	Type        string
	Name        string `gorm:"uniqueIndex:idx_name_barcode"`
	CategoryID  string
	Barcode     string `gorm:"uniqueIndex:idx_name_barcode"`
	Description string
	Amount      float64
	SupplyPrice float64
	RetailPrice float64
}

func (p *Product) ToDomain() *domains.ProductDomain {
	domain := new(domains.ProductDomain)
	if err := copier.Copy(domain, p); err != nil {
		log.Println("failed to copy entity to domain:", err)
	}
	return domain
}

func (p *Product) BeforeCreate(*gorm.DB) error {
	p.ID = uuid.NewString()
	return nil
}
