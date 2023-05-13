package routes

import (
	"prakerja/controllers"
	"prakerja/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ReportsRoutes(e *echo.Echo) *echo.Echo {
	// Middleware for authentication and role-based access control
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define route groups based on role
	adminGroup := e.Group("/admin", middlewares.AdminAuthMiddleware)
	userGroup := e.Group("/user", middlewares.UserAuthMiddleware)

	adminGroup.GET("/reports", controllers.AdminGenerateReport)
	adminGroup.GET("/reports/:user_id", controllers.AdminGenerateReportByUserID)

	userGroup.GET("/reports", controllers.UserGenerateReport)

	return e
}
