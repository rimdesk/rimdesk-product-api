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

var (
	WarehouseURL      = "http://warehouse-api:8080/v1"
	WarehouseEndpoint = "warehouses"
)

type WarehouseClient interface {
	GetById(*fiber.Ctx, string) (*domains.WarehouseDomain, error)
}

type warehouseClient struct {
	http *http.Client
}

func NewWarehouseClient() WarehouseClient {
	return &warehouseClient{
		http: http.DefaultClient,
	}
}

func (warehouse *warehouseClient) GetById(ctx *fiber.Ctx, ID string) (*domains.WarehouseDomain, error) {
	url := fmt.Sprintf("%s/%s/%s", WarehouseURL, WarehouseEndpoint, ID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", security.GetAccessToken(ctx))

	response, err := warehouse.http.Do(req)
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
