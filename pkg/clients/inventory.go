package clients

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rimdesk/product-api/pkg/common"
	"github.com/rimdesk/product-api/pkg/data/domains"
	"github.com/rimdesk/product-api/pkg/security"
	"net/http"
)

const (
	InventoryBaseUrl  = "http://inventory-api:8080/v1"
	InventoryEndpoint = "inventory"
)

type InventoryClient interface {
	GetById(ctx *fiber.Ctx, ID string) (*domains.InventoryDomain, error)
}

type inventoryClient struct {
	http *http.Client
}

func NewInventoryClient() InventoryClient {
	return &inventoryClient{
		http: http.DefaultClient,
	}
}

func (client *inventoryClient) GetById(ctx *fiber.Ctx, ID string) (*domains.InventoryDomain, error) {
	url := fmt.Sprintf("%s/%s/%s", InventoryBaseUrl, InventoryEndpoint, ID)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", security.GetAccessToken(ctx))

	response, err := client.http.Do(req)
	if err != nil {
		return nil, err
	}

	domain := new(common.ApiResponse[domains.InventoryDomain])
	if err := json.NewDecoder(response.Body).Decode(domain); err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("%s", fmt.Errorf("%s", domain.Errors[0]))
	}

	if !domain.Success {
		return nil, fmt.Errorf("%s", domain.Errors[0])
	}

	return &domain.Data, nil
}
