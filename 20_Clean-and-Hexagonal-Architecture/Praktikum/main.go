package main

import (
	"Praktikum/config"
	"Praktikum/routes"
)

func main() {

	config.InitDB()

	router := routes.Route()

	router.Start(":8000")
}
