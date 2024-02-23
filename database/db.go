package database

import (
	"fmt"
	"log"

	"github.com/haviz000/superindo-retail/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB_HOST     = "localhost"
	DB_USER     = "postgres"
	DB_PASSWORD = "admin123"
	DB_NAME     = "superindo_retail"
	DB_PORT     = "5432"

	db  *gorm.DB
	err error
)

func ConnectDB() {
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("failed connecting to database", err)
	}

	db.Debug().AutoMigrate(model.User{}, model.ProductCategory{})
}

func GetDB() *gorm.DB {
	return db
}
