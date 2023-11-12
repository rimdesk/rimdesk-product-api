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
	UserBaseUrl  = "http://user-api:8080/v1"
	UserEndpoint = "users"
)

type UserClient interface {
	GetById(ctx *fiber.Ctx, ID string) (*domains.UserDomain, error)
}

type userClient struct {
	http *http.Client
}

func NewUserClient() UserClient {
	return &userClient{
		http: http.DefaultClient,
	}
}

func (client *userClient) GetById(ctx *fiber.Ctx, ID string) (*domains.UserDomain, error) {
	url := fmt.Sprintf("%s/%s/%s", UserBaseUrl, UserEndpoint, ID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", security.GetAccessToken(ctx))

	response, err := client.http.Do(req)
	if err != nil {
		return nil, err
	}

	domain := new(common.ApiResponse[*domains.UserDomain])
	if err := json.NewDecoder(response.Body).Decode(domain); err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("%s", fmt.Errorf("%s", domain.Errors[0]))
	}

	if !domain.Success {
		return nil, fmt.Errorf("%s", domain.Errors[0])
	}

	return domain.Data, nil
}
