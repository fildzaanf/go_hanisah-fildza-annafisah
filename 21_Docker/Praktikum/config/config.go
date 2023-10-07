package config

import (
	"Praktikum/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	InitDB()
	InitialMigration()
}

func InitDB() {

	var err error
	DB, err = gorm.Open("mysql", "root:my-secret-pw@tcp(172.17.0.2:3306)/crud_go")
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.Blog{})
}
