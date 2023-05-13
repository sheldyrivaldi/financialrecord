package routes

import (
	"prakerja/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) *echo.Echo {
	e.POST("/register", controllers.Registration)

	return e
}
