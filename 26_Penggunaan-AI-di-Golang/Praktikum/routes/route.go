package routes

import (
	"Praktikum/controllers"
	"Praktikum/middlewares"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func New() *echo.Echo {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Cannot load .env file. Err: %s", err)
	}
	e := echo.New()
	middlewares.AddTrailingSlash(e)
	middlewares.Logger(e)

	e.POST("/rekomendasi-laptop/", controllers.RecommendedLaptopController)

	return e
}
