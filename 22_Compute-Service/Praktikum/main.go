package main

import (
	"Praktikum/config"
	"Praktikum/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
