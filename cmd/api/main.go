package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rimdesk/product-api/internal/clients"
	"github.com/rimdesk/product-api/internal/config"
	"github.com/rimdesk/product-api/internal/controllers"
	"github.com/rimdesk/product-api/internal/data/repository"
	"github.com/rimdesk/product-api/internal/database"
	"github.com/rimdesk/product-api/internal/router"
	"github.com/rimdesk/product-api/internal/service"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	cfg = config.New()
	db  = database.NewGormDatabase()
)

func init() {
	cfg.LoadEnv()
	db.SetConfig(cfg.DatabaseConfig())
	db.ConnectDB()
}

func main() {
	serverAddr := fmt.Sprintf(":%s", os.Getenv("APP.PORT"))
	dbEngine := db.GetEngine().(*gorm.DB)
	app := fiber.New(cfg.AppConfig())
	app.Use(cors.New())
	app.Use(requestid.New())

	warehouseClient := clients.NewWarehouseClient()

	productRepository := repository.NewProductRepository(dbEngine)
	productService := service.NewProductService(productRepository, warehouseClient)
	productController := controllers.NewProductController(productService, validator.New())
	rtr := router.NewFiberRouter(app, productController)
	rtr.ApiR0utes()

	log.Println("REST server started on addr:", serverAddr)
	if err := app.Listen(serverAddr); err != nil {
		log.Println("failed to listen on addr:", serverAddr)
		os.Exit(1)
	}
}
