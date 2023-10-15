package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func AddTrailingSlash(e *echo.Echo) {
	e.Pre(middleware.AddTrailingSlash())
}
