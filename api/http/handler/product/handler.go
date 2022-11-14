package product

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rimdesk/product-api/api/common"
	"github.com/rimdesk/product-api/api/data/domains"
	"github.com/rimdesk/product-api/api/http/request/product"
	productService "github.com/rimdesk/product-api/api/service/product"
	"time"
)

func List(ctx *fiber.Ctx) error {
	apiResponse := common.ApiResponse[[]domains.ProductDomain]{
		Success:   true,
		Timestamp: time.Now().UnixMilli(),
		Message:   "Products fetched successfully",
		Code:      fiber.StatusOK,
	}

	apiResponse.Data = productService.GetAllProducts()

	return ctx.Status(fiber.StatusOK).JSON(apiResponse)
}

func Get(ctx *fiber.Ctx) error {
	apiResponse := common.ApiResponse[*domains.ProductDomain]{
		Success:   true,
		Timestamp: time.Now().UnixMilli(),
		Code:      fiber.StatusOK,
	}

	productDomain, err := productService.GetProductById(ctx.Params("id"))
	if err != nil {
		apiResponse.Success = false
		apiResponse.Errors = []string{err.Error()}
		apiResponse.Code = fiber.StatusNotFound

		return ctx.Status(fiber.StatusNotFound).JSON(apiResponse)
	}

	apiResponse.Data = productDomain
	apiResponse.Message = "Product fetched successfully"

	return ctx.Status(fiber.StatusOK).JSON(apiResponse)
}

func Post(ctx *fiber.Ctx) error {
	apiResponse := common.ApiResponse[*domains.ProductDomain]{
		Success:   true,
		Timestamp: time.Now().UnixMilli(),
		Code:      fiber.StatusCreated,
	}

	dto := new(product.Dto)
	if err := ctx.BodyParser(dto); err != nil {
		apiResponse.Success = false
		apiResponse.Errors = []string{err.Error()}
		apiResponse.Code = fiber.StatusBadRequest

		return ctx.Status(fiber.StatusBadRequest).JSON(apiResponse)
	}

	if err := validator.New().Struct(dto); err != nil {
		apiResponse.Success = false
		apiResponse.Errors = []string{err.Error()}
		apiResponse.Code = fiber.StatusBadRequest

		return ctx.Status(fiber.StatusBadRequest).JSON(apiResponse)
	}

	productDomain, err := productService.CreateProduct(dto)
	if err != nil {
		apiResponse.Success = false
		apiResponse.Errors = []string{err.Error()}
		apiResponse.Code = fiber.StatusInternalServerError

		return ctx.Status(fiber.StatusInternalServerError).JSON(apiResponse)
	}

	apiResponse.Data = productDomain
	apiResponse.Message = "Product created successfully"

	return ctx.Status(fiber.StatusCreated).JSON(apiResponse)
}

func Patch(ctx *fiber.Ctx) error {
	apiResponse := common.ApiResponse[*domains.ProductDomain]{
		Success:   true,
		Timestamp: time.Now().UnixMilli(),
		Code:      fiber.StatusOK,
	}

	dto := new(product.Dto)
	if err := ctx.BodyParser(dto); err != nil {
		apiResponse.Success = false
		apiResponse.Errors = []string{err.Error()}
		apiResponse.Code = fiber.StatusBadRequest

		return ctx.Status(fiber.StatusBadRequest).JSON(apiResponse)
	}

	if err := validator.New().Struct(dto); err != nil {
		apiResponse.Success = false
		apiResponse.Errors = []string{err.Error()}
		apiResponse.Code = fiber.StatusBadRequest

		return ctx.Status(fiber.StatusBadRequest).JSON(apiResponse)
	}

	productDomain, err := productService.UpdateProduct(ctx.Params("id"), dto)
	if err != nil {
		apiResponse.Success = false
		apiResponse.Errors = []string{err.Error()}
		apiResponse.Code = fiber.StatusInternalServerError

		return ctx.Status(fiber.StatusInternalServerError).JSON(apiResponse)
	}

	apiResponse.Data = productDomain
	apiResponse.Message = "Product updated successfully"

	return ctx.Status(fiber.StatusOK).JSON(apiResponse)
}

func Delete(ctx *fiber.Ctx) error {
	apiResponse := common.ApiResponse[*domains.ProductDomain]{
		Success:   true,
		Timestamp: time.Now().UnixMilli(),
		Code:      fiber.StatusOK,
	}

	err := productService.DeleteProduct(ctx.Params("id"))
	if err != nil {
		apiResponse.Success = false
		apiResponse.Errors = []string{err.Error()}
		apiResponse.Code = fiber.StatusInternalServerError

		return ctx.Status(fiber.StatusInternalServerError).JSON(apiResponse)
	}

	apiResponse.Message = "Product deleted successfully"

	return ctx.Status(fiber.StatusOK).JSON(apiResponse)
}
