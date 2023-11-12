package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rimdesk/product-api/pkg/common"
	"github.com/rimdesk/product-api/pkg/exceptions"
)

type AppMiddleware interface {
	CompanyHeaderPresent(*fiber.Ctx) error
}

type fiberMiddleware struct {
}

func (ware *fiberMiddleware) CompanyHeaderPresent(ctx *fiber.Ctx) error {
	apiResponse := common.NewApiResponse()
	company := ctx.GetReqHeaders()["X-Company-Id"]
	if company == "" {
		apiResponse.Success = false
		apiResponse.Errors = []string{exceptions.ErrCompanyHeaderMustBePresent.Error()}
		return ctx.Status(fiber.StatusBadRequest).JSON(apiResponse)
	}

	return ctx.Next()
}

func NewFiberMiddleware() AppMiddleware {
	return &fiberMiddleware{}
}
