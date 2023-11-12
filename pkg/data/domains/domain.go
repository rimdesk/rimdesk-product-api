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
	SupplyPrice float32    `json:"supply_price"`
	RetailPrice float32    `json:"retail_price"`
	CreatedAt   *time.Time `json:"created_at"`
}

type WarehouseDomain struct {
	ID                    string     `json:"id"`
	Name                  string     `json:"name"`
	PhoneNumber           string     `json:"phone_number"`
	DiscountLimit         float32    `json:"discount_limit"`
	SupervisorPercentage  float32    `json:"supervisor_percentage"`
	SalesPersonPercentage float32    `json:"sales_person_percentage"`
	Enabled               bool       `json:"enabled"`
	IsCommissionEnabled   bool       `json:"is_commission_enabled"`
	IsWholesaleEnabled    bool       `json:"is_wholesale_enabled"`
	IsRawMaterialsEnabled bool       `json:"is_raw_materials_enabled"`
	CreatedAt             *time.Time `json:"created_at"`
}

type CategoryDomain struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
}

type CompanyDomain struct {
	ID                 string          `json:"id"`
	Name               string          `json:"name"`
	Email              string          `json:"email"`
	PhoneNumber        string          `json:"phone_number"`
	TinNumber          string          `json:"tin_number"`
	RegistrationNumber string          `json:"registration_number"`
	Currency           string          `json:"currency"`
	Logo               string          `json:"logo"`
	Category           *CategoryDomain `json:"category"`
}

type InventoryDomain struct {
	ID        string        `json:"id"`
	Product   ProductDomain `json:"product"`
	Quantity  uint          `json:"quantity"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type UserDomain struct {
	ID           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	EmailAddress string `json:"email_address"`
	PhoneNumber  string `json:"phone_number"`
	IsEnabled    bool   `json:"is_enabled"`
}
