package repository

import (
	"fmt"
	"go-crud-app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "host=db user=postgres password=password dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Database connected successfully.")
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Todo{})
}
