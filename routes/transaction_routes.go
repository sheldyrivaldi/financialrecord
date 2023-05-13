package routes

import (
	"prakerja/controllers"
	"prakerja/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TransactionRoutes(e *echo.Echo) *echo.Echo {
	// Middleware for authentication and role-based access control
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define route groups based on role
	adminGroup := e.Group("/admin", middlewares.AdminAuthMiddleware)
	userGroup := e.Group("/user", middlewares.UserAuthMiddleware)

	adminGroup.GET("/transactions", controllers.AdminGetTransactions)
	adminGroup.GET("/transactions:user_id", controllers.AdminGetTransactionsByUserId)
	adminGroup.GET("/transactions/:transaction_id", controllers.AdminGetTransactionsByTransactionId)

	userGroup.GET("/transactions", controllers.UserGetTransactions)
	userGroup.POST("/transactions", controllers.UserCreateTransaction)
	userGroup.DELETE("/transactions/:transaction_id", controllers.UserDeleteTransactionByTransactionId)
	return e
}
