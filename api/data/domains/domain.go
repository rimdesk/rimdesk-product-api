package domains

import "time"

type ProductDomain struct {
	ID          string     `json:"id"`
	Type        string     `json:"type"`
	Name        string     `json:"name"`
	CategoryID  string     `json:"category_id"`
	Barcode     string     `json:"barcode"`
	Description string     `json:"description"`
	Amount      float32    `json:"amount"`
	UnitPrice   float32    `json:"unitPrice"`
	CreatedAt   *time.Time `json:"createdAt"`
}
