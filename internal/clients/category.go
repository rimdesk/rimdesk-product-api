package clients

import (
	"encoding/json"
	"fmt"
	"github.com/rimdesk/product-api/internal/common"
	"github.com/rimdesk/product-api/internal/data/domains"
	"net/http"
)

const (
	CategoryBaseUrl  = "http://category-api:8080/v1"
	CategoryEndpoint = "categories"
)

type CategoryClient struct {
	Client *http.Client
}

func NewCategoryClient() *CategoryClient {
	return &CategoryClient{
		Client: &http.Client{},
	}
}

func (c *CategoryClient) GetById(ID string) (*domains.CategoryDomain, error) {
	url := fmt.Sprintf("%s/%s/%s", CategoryBaseUrl, CategoryEndpoint, ID)
	req, _ := http.NewRequest("GET", url, nil)
	reqHeader := http.Header{}
	reqHeader.Set("Authorization", "Bearer token-goes-here")
	req.Header = reqHeader

	response, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	domain := new(common.ApiResponse[domains.CategoryDomain])
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
