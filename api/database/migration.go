package database

import (
	"github.com/rimdesk/product-api/api/data/entities"
	"gorm.io/gorm"
	"log"
)

func HandleMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&entities.Product{})
	if err != nil {
		log.Fatalf("migration failed: %s", err.Error())
	}
	log.Println("tables migrated successfully!")
}
