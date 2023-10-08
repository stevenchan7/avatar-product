package config

import (
	"example.com/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	// Setup from gorm docs
	dsn := "root:Mismag0203i9@tcp(127.0.0.1:3306)/avatar_product?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Checking error
	if err != nil {
		panic("Failed to connect database")
	}

	// Craete table, sync model
	db.AutoMigrate(models.Product{})
	db.AutoMigrate(models.User{})

	return db
}
