package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rimdesk/product-api/pkg/controllers"
	"github.com/rimdesk/product-api/pkg/security"
)

type AppRouter interface {
	ApiR0utes()
}

type fiberRouter struct {
	app               *fiber.App
	productController controllers.ProductController
}

func (engine *fiberRouter) ApiR0utes() {
	v1 := engine.app.Group("v1")
	// Register routes here
	route := v1.Group("products", logger.New())
	route.Use(security.IsAuthorizedJWT)
	route.Get("", engine.productController.List)
	route.Get("/:id", engine.productController.Get)
	route.Post("", engine.productController.Post)
	route.Patch("/:id", engine.productController.Patch)
	route.Delete("/:id", engine.productController.Delete)
	route.Get("/search", engine.productController.Search)
}

func NewFiberRouter(app *fiber.App, controller controllers.ProductController) AppRouter {
	return &fiberRouter{
		app:               app,
		productController: controller,
	}
}
