package routes

import (
	"Praktikum/config"
	"Praktikum/constants"
	"Praktikum/controllers"
	"Praktikum/repository"
	"Praktikum/usecase"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route() *echo.Echo {

	e := echo.New()

	repository := repository.New(config.DB)
	usecase := usecase.NewUserUseCase(repository)
	controllers := controllers.New(usecase)

	JWT := middleware.JWT([]byte(constants.SECRET_JWT))

	gUser := e.Group("/users")
	gUser.GET("/", controllers.GetUsersController, JWT)
	gUser.POST("/login/", controllers.LoginUserController)
	gUser.POST("/", controllers.CreateUserController)

	return e
}
