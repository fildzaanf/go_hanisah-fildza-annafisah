package routes

import (
	"project/constants"
	"project/controllers"
	"project/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()

	middlewares.AddTrailingSlash(e)
	middlewares.Logger(e)

	JWT := middleware.JWT([]byte(constants.SECRET_JWT))

	userGroup := e.Group("/users")
	userGroup.GET("/", controllers.GetUsersController, JWT)
	userGroup.GET("/:id/", controllers.GetUserController, JWT)
	userGroup.POST("/", controllers.CreateUserController)
	userGroup.DELETE("/:id/", controllers.DeleteUserController, JWT)
	userGroup.PUT("/:id/", controllers.UpdateUserController, JWT)
	userGroup.POST("/login/", controllers.LoginUserController)

	bookGroup := e.Group("/books")
	bookGroup.GET("/", controllers.GetBooksController, JWT)
	bookGroup.GET("/:id/", controllers.GetBookController, JWT)
	bookGroup.POST("/", controllers.CreateBookController, JWT)
	bookGroup.DELETE("/:id/", controllers.DeleteBookController, JWT)
	bookGroup.PUT("/:id/", controllers.UpdateBookController, JWT)

	blogGroup := e.Group("/blogs")
	blogGroup.GET("/", controllers.GetBlogsController, JWT)
	blogGroup.GET("/:id/", controllers.GetBlogController, JWT)
	blogGroup.POST("/", controllers.CreateBlogController, JWT)
	blogGroup.DELETE("/:id/", controllers.DeleteBlogController, JWT)
	blogGroup.PUT("/:id/", controllers.UpdateBlogController, JWT)

	return e
}
