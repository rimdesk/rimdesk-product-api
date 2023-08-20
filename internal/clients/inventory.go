package clients

import (
	"encoding/json"
	"fmt"
	"github.com/rimdesk/product-api/internal/common"
	"github.com/rimdesk/product-api/internal/data/domains"
	"net/http"
)

const (
	InventoryBaseUrl  = "http://inventory-api:8080/v1"
	InventoryEndpoint = "inventory"
)

type InventoryClient struct {
	Client *http.Client
}

func NewInventoryClient() *InventoryClient {
	return &InventoryClient{
		Client: &http.Client{},
	}
}

func (c *InventoryClient) GetById(ID string) (*domains.InventoryDomain, error) {
	url := fmt.Sprintf("%s/%s/%s", InventoryBaseUrl, InventoryEndpoint, ID)
	req, _ := http.NewRequest("GET", url, nil)
	reqHeader := http.Header{}
	reqHeader.Set("Authorization", "Bearer token-goes-here")
	req.Header = reqHeader

	response, err := c.Client.Do(req)
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
