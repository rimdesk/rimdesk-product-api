package bootstrap

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rimdesk/product-api/api/config"
	"github.com/rimdesk/product-api/api/database"
	"github.com/rimdesk/product-api/api/http/middleware"
	"github.com/rimdesk/product-api/api/routes"
	"log"
)

type AppRunner struct {
	App *fiber.App
}

func New() *AppRunner {
	config.Init()
	appRunner := new(AppRunner)
	app := fiber.New(config.AppConfig())
	appRunner.App = app
	middleware.RegisterMiddlewares(app)
	database.ConnectDB()
	routes.SetupApiRoutes(app)

	return appRunner
}

func (r AppRunner) StartAndListen(host string, port string) {
	log.Fatalln(r.App.Listen(fmt.Sprintf("%s:%s", host, port)))
}
