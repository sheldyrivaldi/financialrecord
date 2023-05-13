package routes

import (
	"prakerja/controllers"
	"prakerja/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func IncomeCategories(e *echo.Echo) *echo.Echo {
	// Middleware for authentication and role-based access control
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define route groups based on role
	adminGroup := e.Group("/admin", middlewares.AdminAuthMiddleware)
	userGroup := e.Group("/user", middlewares.UserAuthMiddleware)

	adminGroup.GET("/categories/income", controllers.AdminGetIncomeCategoreies)
	adminGroup.GET("/categories/income/:user_id", controllers.AdminGetIncomeCategoreiesByUserID)

	userGroup.GET("/categories/income", controllers.UserGetIncomeCategories)
	userGroup.PUT("/categories/income/:id", controllers.UserUpdateIncomeCategoryById)
	userGroup.POST("/categories/income", controllers.UserCreateIncomeCategory)
	userGroup.DELETE("/categories/income/:id", controllers.UserDeleteIncomeCategoryById)

	return e
}
