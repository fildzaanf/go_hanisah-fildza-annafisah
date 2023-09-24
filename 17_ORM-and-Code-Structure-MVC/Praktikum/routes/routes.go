package routes

import (
	"Praktikum/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()

	e.Pre(middleware.AddTrailingSlash())

	userGroup := e.Group("/users")
	userGroup.GET("/", controllers.GetUsersController)
	userGroup.GET("/:id/", controllers.GetUserController)
	userGroup.POST("/", controllers.CreateUserController)
	userGroup.DELETE("/:id/", controllers.DeleteUserController)
	userGroup.PUT("/:id/", controllers.UpdateUserController)

	bookGroup := e.Group("/books")
	bookGroup.GET("/", controllers.GetBooksController)
	bookGroup.GET("/:id/", controllers.GetBookController)
	bookGroup.POST("/", controllers.CreateBookController)
	bookGroup.DELETE("/:id/", controllers.DeleteBookController)
	bookGroup.PUT("/:id/", controllers.UpdateBookController)

	blogGroup := e.Group("/blogs")
	blogGroup.GET("/", controllers.GetBlogsController)
	blogGroup.GET("/:id/", controllers.GetBlogController)
	blogGroup.POST("/", controllers.CreateBlogController)
	blogGroup.DELETE("/:id/", controllers.DeleteBlogController)
	blogGroup.PUT("/:id/", controllers.UpdateBlogController)

	return e
}
