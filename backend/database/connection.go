package database

import (
	"github.com/aungmyatmoe11/Go-JWT-Auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// dsn := "root:/Go-JWT-Auth"
	// dsn := "root:@tcp(127.0.0.1:3306)/Go-JWT-Auth?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:@/Go-JWT-Auth?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
