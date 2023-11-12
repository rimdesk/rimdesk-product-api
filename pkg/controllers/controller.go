package controllers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rimdesk/product-api/pkg/common"
	"github.com/rimdesk/product-api/pkg/data/dtos"
	"github.com/rimdesk/product-api/pkg/service"
)

type ProductController interface {
	List(ctx *fiber.Ctx) error
	Get(ctx *fiber.Ctx) error
	Post(ctx *fiber.Ctx) error
	Patch(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Search(ctx *fiber.Ctx) error
}

type productController struct {
	productService service.ProductService
	validator      *validator.Validate
}

func NewProductController(productService service.ProductService, validate *validator.Validate) ProductController {
	return &productController{productService: productService, validator: validate}
}

func (controller *productController) List(ctx *fiber.Ctx) error {
	apiResponse := common.NewApiResponse()
	companyID := ctx.GetReqHeaders()["X-Company-Id"]

	apiResponse.Data, _ = controller.productService.GetAllProducts(ctx, companyID)

	return ctx.Status(fiber.StatusOK).JSON(apiResponse)
}

func (controller *productController) Get(ctx *fiber.Ctx) error {
	apiResponse := common.NewApiResponse()
	companyID := ctx.GetReqHeaders()["X-Company-Id"]

	productDomain, err := controller.productService.GetProductByCompanyAndId(ctx, companyID, ctx.Params("id"))
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

func (controller *productController) Post(ctx *fiber.Ctx) error {
	apiResponse := common.NewApiResponse()
	companyID := ctx.GetReqHeaders()["X-Company-Id"]

	dto := new(dtos.ProductDto)
	if err := ctx.BodyParser(dto); err != nil {
		apiResponse.Success = false
		apiResponse.Errors = []string{fmt.Sprintf("failed to parse dto: %s", err.Error())}
		apiResponse.Code = fiber.StatusBadRequest

		return ctx.Status(fiber.StatusBadRequest).JSON(apiResponse)
	}

	if err := controller.validator.Struct(dto); err != nil {
		apiResponse.Success = false
		apiResponse.Errors = []string{fmt.Sprintf("failed to parse dto: %s", err.Error())}
		apiResponse.Code = fiber.StatusBadRequest

		return ctx.Status(fiber.StatusBadRequest).JSON(apiResponse)
	}

	productDomain, err := controller.productService.CreateProduct(ctx, companyID, dto)
	if err != nil {
		apiResponse.Success = false
		apiResponse.Errors = []string{err.Error()}
		apiResponse.Code = fiber.StatusInternalServerError

		return ctx.Status(fiber.StatusInternalServerError).JSON(apiResponse)
	}

	apiResponse.Data = productDomain
	apiResponse.Message = "Product created successfully"
	apiResponse.Code = fiber.StatusCreated

	return ctx.Status(fiber.StatusCreated).JSON(apiResponse)
}

func (controller *productController) Patch(ctx *fiber.Ctx) error {
	apiResponse := common.NewApiResponse()
	companyID := ctx.GetReqHeaders()["X-Company-Id"]

	dto := new(dtos.ProductDto)
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

	productDomain, err := controller.productService.UpdateProduct(ctx, companyID, ctx.Params("id"), dto)
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

func (controller *productController) Delete(ctx *fiber.Ctx) error {
	apiResponse := common.NewApiResponse()
	companyID := ctx.GetReqHeaders()["X-Company-Id"]

	err := controller.productService.DeleteProduct(ctx, companyID, ctx.Params("id"))
	if err != nil {
		apiResponse.Success = false
		apiResponse.Errors = []string{err.Error()}
		apiResponse.Code = fiber.StatusInternalServerError

		return ctx.Status(fiber.StatusInternalServerError).JSON(apiResponse)
	}

	apiResponse.Message = "Product deleted successfully"

	return ctx.Status(fiber.StatusOK).JSON(apiResponse)
}

func (controller *productController) Search(ctx *fiber.Ctx) error {
	apiResponse := common.NewApiResponse()
	companyID := ctx.GetReqHeaders()["X-Company-Id"]

	dto := new(dtos.ProductSearchDto)
	if err := ctx.QueryParser(dto); err != nil {
		apiResponse.Success = false
		apiResponse.Errors = []string{err.Error()}
		apiResponse.Code = fiber.StatusBadRequest

		return ctx.Status(fiber.StatusBadRequest).JSON(apiResponse)
	}

	if err := controller.validator.Struct(dto); err != nil {
		apiResponse.Success = false
		apiResponse.Errors = []string{err.Error()}
		apiResponse.Code = fiber.StatusBadRequest

		return ctx.Status(fiber.StatusBadRequest).JSON(apiResponse)
	}

	productDomains, err := controller.productService.SearchWarehouse(ctx, companyID, dto)
	if err != nil {
		apiResponse.Success = false
		apiResponse.Errors = []string{err.Error()}
		apiResponse.Code = fiber.StatusBadRequest

		return ctx.Status(fiber.StatusBadRequest).JSON(apiResponse)
	}

	apiResponse.Data = productDomains
	apiResponse.Message = "Warehouse "

	return ctx.Status(fiber.StatusOK).JSON(apiResponse)
}
