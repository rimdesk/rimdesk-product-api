package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rimdesk/product-api/api/routes/product"
)

func SetupApiRoutes(app *fiber.App) {
	v1 := app.Group("v1", logger.New())
	// Register routes here
	product.SetupProductRoutes(v1)
}
