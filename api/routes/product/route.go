package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rimdesk/product-api/api/http/handler/product"
)

func SetupProductRoutes(router fiber.Router) {
	productRoutes := router.Group("products")
	productRoutes.Get("", product.List)
	productRoutes.Get("/:id", product.Get)
	productRoutes.Post("", product.Post)
	productRoutes.Patch("/:id", product.Patch)
	productRoutes.Delete("/:id", product.Delete)
}
