package routes

import (
	"prakerja/controllers"

	"github.com/labstack/echo/v4"
)

func LoginRoutes(e *echo.Echo) *echo.Echo {
	e.POST("/login", controllers.Login)

	return e
}
