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
	CompanyBaseUrl  = "http://company-api:8080/v1"
	CompanyEndpoint = "companies"
)

type CompanyClient interface {
	GetById(ctx *fiber.Ctx, ID string) (*domains.CompanyDomain, error)
}

type companyClient struct {
	http *http.Client
}

func NewCompanyClient() CompanyClient {
	return &companyClient{
		http: http.DefaultClient,
	}
}

func (client *companyClient) GetById(ctx *fiber.Ctx, ID string) (*domains.CompanyDomain, error) {
	url := fmt.Sprintf("%s/%s/%s", CompanyBaseUrl, CompanyEndpoint, ID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", security.GetAccessToken(ctx))

	response, err := client.http.Do(req)
	if err != nil {
		return nil, err
	}

	domain := new(common.ApiResponse[domains.CompanyDomain])
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
