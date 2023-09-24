package configs

import{
	"fmt"
	"Praktikum/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
}

var DB *gorm.DB

func init() {
	InitDB()
	InitialMigration()
}

func InitDB() {
	config := LoadConfig()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)
	
	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.Blog{})
}