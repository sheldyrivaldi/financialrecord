package routes

import (
	"prakerja/controllers"
	"prakerja/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserRoutes(e *echo.Echo) *echo.Echo {
	// Middleware for authentication and role-based access control
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define route groups based on role
	adminGroup := e.Group("/admin", middlewares.AdminAuthMiddleware)
	userGroup := e.Group("/user", middlewares.UserAuthMiddleware)

	adminGroup.GET("/users", controllers.AdminGetUsers)
	adminGroup.GET("/users/:id", controllers.AdminGetUserById)

	userGroup.GET("/profile", controllers.UserGetProfile)
	userGroup.PUT("/profile", controllers.UserUpdateProfile)

	return e
}
