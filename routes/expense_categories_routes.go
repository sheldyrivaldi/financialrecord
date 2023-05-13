package routes

import (
	"prakerja/controllers"
	"prakerja/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ExpenseCategories(e *echo.Echo) *echo.Echo {
	// Middleware for authentication and role-based access control
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define route groups based on role
	adminGroup := e.Group("/admin", middlewares.AdminAuthMiddleware)
	userGroup := e.Group("/user", middlewares.UserAuthMiddleware)

	adminGroup.GET("/categories/expense", controllers.AdminGetExpenseCategoreies)
	adminGroup.GET("/categories/expense/:user_id", controllers.AdminGetExpenseCategoreiesByUserID)

	userGroup.GET("/categories/expense", controllers.UserGetExpenseCategories)
	userGroup.PUT("/categories/expense/:id", controllers.UserUpdateExpenseCategoryById)
	userGroup.POST("/categories/expense", controllers.UserCreateExpenseCategory)
	userGroup.DELETE("/categories/expense/:id", controllers.UserDeleteExpenseCategoryById)

	return e
}
