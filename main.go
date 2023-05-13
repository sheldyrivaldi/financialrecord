package main

import (
	"prakerja/configs"
	"prakerja/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.ConnectDatabase()

	e := echo.New()

	routes.LoginRoutes(e)
	routes.RegisterRoutes(e)
	routes.UserRoutes(e)
	routes.ExpenseCategories(e)
	routes.IncomeCategories(e)
	routes.TransactionRoutes(e)
	routes.ReportsRoutes(e)

	e.Start(":8080")
}
