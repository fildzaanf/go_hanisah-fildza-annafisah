package main

import{
	"Praktikum/configs"
	"Praktikum/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
}

func main(){
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}