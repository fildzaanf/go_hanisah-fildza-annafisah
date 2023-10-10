package config

import (
	"project/models"

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
	DB, err = gorm.Open("mysql", "root:fildzaanf123@tcp(dbfildzaanf.cpdi16ed6gbs.ap-southeast-1.rds.amazonaws.com:3306)/dbfildza")
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.Blog{})
}
