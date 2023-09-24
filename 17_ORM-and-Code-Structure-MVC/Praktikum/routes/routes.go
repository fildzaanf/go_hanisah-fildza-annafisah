package routes

import{
	"praktikum/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
}

func New() *echo.Echo {

	e := echo.New()
	
	e.Pre(middleware.AddTrailingSlash())

	userGroup := e.Group("/users")
	userGroup.GET("/", controller.GetUsersController)
	userGroup.GET("/:id/", controller.GetUserController)
	userGroup.POST("/", controller.CreateUserController)
	userGroup.DELETE("/:id/", controller.DeleteUserController)
	userGroup.PUT("/:id/", controller.UpdateUserController)


	bookGroup := e.Group("/books")
	bookGroup.GET("/", controller.GetBooksController)
	bookGroup.GET("/:id/", controller.GetBookController)
	bookGroup.POST("/", controller.CreateBookController)
	bookGroup.DELETE("/:id/", controller.DeleteBookController)
	bookGroup.PUT("/:id/", controller.UpdateBookController)
	

	blogGroup := e.Group("/blogs")
	blogGroup.GET("/", controller.GetBlogsController)
	blogGroup.GET("/:id/", controller.GetBlogController)
	blogGroup.POST("/", controller.CreateBlogController)
	blogGroup.DELETE("/:id/", controller.DeleteBlogController)
	blogGroup.PUT("/:id/", controller.UpdateBlogController)

	return e
}