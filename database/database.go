package database

import (
	"go-gin-crud/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	// dsn := "yourusername:yourpassword@tcp(localhost:3306)/yourdbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:@tcp(localhost)/allugoh?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		os.Exit(1)
	}

	err = DB.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
		os.Exit(1)
	}
}
