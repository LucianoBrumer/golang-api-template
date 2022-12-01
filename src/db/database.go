package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"golang-api/src/models"
)

var DB *gorm.DB

func Database() {
	dsn := "root@tcp(127.0.0.1:3306)/gotest?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database.")
	}

	DB.AutoMigrate(&models.User{})
}
