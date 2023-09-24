package main

import (
	"Praktikum/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
