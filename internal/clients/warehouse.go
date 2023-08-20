package clients

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rimdesk/product-api/internal/common"
	"github.com/rimdesk/product-api/internal/data/domains"
	"github.com/rimdesk/product-api/internal/security"
	"net/http"
)

var (
	WarehouseURL      = "http://warehouse-api:8080/v1"
	WarehouseEndpoint = "warehouses"
)

type WarehouseClient interface {
	GetById(*fiber.Ctx, string) (*domains.WarehouseDomain, error)
}

type warehouseClient struct {
	Client *http.Client
}

func NewWarehouseClient() WarehouseClient {
	return &warehouseClient{
		Client: http.DefaultClient,
	}
}

func (warehouse *warehouseClient) GetById(ctx *fiber.Ctx, ID string) (*domains.WarehouseDomain, error) {
	url := fmt.Sprintf("%s/%s/%s", WarehouseURL, WarehouseEndpoint, ID)
	req, _ := http.NewRequest("GET", url, nil)

	accessToken := security.GetAccessToken(ctx)
	req.Header.Set("Authorization", accessToken)

	response, err := warehouse.Client.Do(req)
	if err != nil {
		return nil, err
	}

	d := new(common.ApiResponse[domains.WarehouseDomain])
	if err := json.NewDecoder(response.Body).Decode(d); err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("%s", fmt.Errorf("%s", d.Errors[0]))
	}

	if !d.Success {
		return nil, fmt.Errorf("%s", d.Errors[0])
	}

	return &d.Data, nil
}
