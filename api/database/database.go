package database

import (
	"fmt"
	"github.com/rimdesk/product-api/api/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	dbType := config.Get("DB_TYPE")
	log.Println("Loading db config for: ", dbType)
	var err error
	switch dbType {
	case "mysql":
		DB, err = LoadMysqlDB()
		if err != nil {
			log.Fatalln("unable to load sqlite db", err.Error())
		}
	case "sqlite":
		DB, err = LoadSqliteDB()
		if err != nil {
			log.Fatalln("unable to load sqlite db", err.Error())
		}
	case "postgres":
		log.Println("loading config for postgres")
		DB, err = LoadPostgresDB()
		if err != nil {
			log.Fatalln("unable to load postgres db", err.Error())
		}

		log.Println("postgres connected successfully")
	}

	HandleMigrations(DB)
}

func LoadSqliteDB() (*gorm.DB, error) {
	dbName := fmt.Sprintf("%s.db", config.Get("DB_NAME"))
	db, err := gorm.Open(sqlite.Open(dbName), config.DatabaseConfig())

	if err != nil {
		panic("failed to connect database")
		return nil, err
	}
	log.Println("connected to sqlite db")

	return db, nil
}

func LoadPostgresDB() (*gorm.DB, error) {
	dsn := config.Get("DB_POSTGRES_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		return nil, err
	}

	return db, nil
}

func LoadMysqlDB() (*gorm.DB, error) {
	dsn := config.Get("DB_MYSQL_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		return nil, err
	}

	return db, nil
}
