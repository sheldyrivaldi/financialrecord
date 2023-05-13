package controllers

import (
	"net/http"
	"prakerja/configs"
	"prakerja/models"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func AdminGenerateReport(c echo.Context) error {
	// Get start and end dates for the report
	endDate := time.Now().Format("02-01-2006")
	startDate := time.Now().AddDate(0, 0, -30).Format("02-01-2006")

	// Calculate total expenses
	var totalExpenses int
	err := configs.DB.Model(&models.Transactions{}).Select("COALESCE(SUM(amount), 0)").Where("type = 'expense' AND date >= ? AND date <= ?", startDate, endDate).Scan(&totalExpenses).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to calculate total expenses")
	}

	// Calculate total income
	var totalIncome int
	err = configs.DB.Model(&models.Transactions{}).Select("COALESCE(SUM(amount), 0)").Where("type = 'income' AND date >= ? AND date <= ?", startDate, endDate).Scan(&totalIncome).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to calculate total income")
	}

	// Calculate difference (income - expenses)
	difference := totalIncome - totalExpenses

	// Get category-wise report
	var categoryReport []models.CategoryReport
	err = configs.DB.Raw("SELECT category, type, COALESCE(SUM(amount), 0) AS total FROM transactions WHERE date >= ? AND date <= ? GROUP BY category, type", startDate, endDate).Scan(&categoryReport).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to generate category report")
	}

	report := models.Report{
		TotalExpenses:  totalExpenses,
		TotalIncome:    totalIncome,
		Difference:     difference,
		CategoryReport: categoryReport,
	}

	return c.JSON(http.StatusOK, report)

}
func AdminGenerateReportByUserID(c echo.Context) error {
	// Get start and end dates for the report
	endDate := time.Now().Format("02-01-2006")
	startDate := time.Now().AddDate(0, 0, -30).Format("02-01-2006")

	// Calculate total expenses
	var totalExpenses int
	userID := c.Param("user_id")
	err := configs.DB.Model(&models.Transactions{}).Select("COALESCE(SUM(amount), 0)").Where("user_id = ? AND type = 'expense' AND date >= ? AND date <= ?", userID, startDate, endDate).Scan(&totalExpenses).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to calculate total expenses")
	}

	// Calculate total income
	var totalIncome int
	err = configs.DB.Model(&models.Transactions{}).Select("COALESCE(SUM(amount), 0)").Where("user_id = ? AND type = 'income' AND date >= ? AND date <= ?", userID, startDate, endDate).Scan(&totalIncome).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to calculate total income")
	}

	// Calculate difference (income - expenses)
	difference := totalIncome - totalExpenses

	// Get category-wise report
	var categoryReport []models.CategoryReport
	err = configs.DB.Raw("SELECT category, type, COALESCE(SUM(amount), 0) AS total FROM transactions WHERE user_id = ? AND date >= ? AND date <= ? GROUP BY category, type", userID, startDate, endDate).Scan(&categoryReport).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to generate category report")
	}

	report := models.Report{
		TotalExpenses:  totalExpenses,
		TotalIncome:    totalIncome,
		Difference:     difference,
		CategoryReport: categoryReport,
	}

	return c.JSON(http.StatusOK, report)
}

func UserGenerateReport(c echo.Context) error {
	// Get JWT token from request header
	authToken := c.Request().Header.Get("Authorization")
	splitString := strings.Split(authToken, "Bearer ")
	tokenString := splitString[1]

	// Parse JWT token
	token, _ := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("YoUr-SeCreT-KeY"), nil
	})
	claims, _ := token.Claims.(*models.Claims)

	userID := strconv.Itoa(claims.ID)

	// Calculate the start and end dates for the report (last 30 days)
	endDate := time.Now().Format("02-01-2006")
	startDate := time.Now().AddDate(0, 0, -30).Format("02-01-2006")

	// Calculate total expenses
	var totalExpenses int
	err := configs.DB.Model(&models.Transactions{}).Select("COALESCE(SUM(amount), 0)").Where("user_id = ? AND type = 'expense' AND date >= ? AND date <= ?", userID, startDate, endDate).Scan(&totalExpenses).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to calculate total expenses")
	}

	// Calculate total income
	var totalIncome int
	err = configs.DB.Model(&models.Transactions{}).Select("COALESCE(SUM(amount), 0)").Where("user_id = ? AND type = 'income' AND date >= ? AND date <= ?", userID, startDate, endDate).Scan(&totalIncome).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to calculate total income")
	}

	// Calculate difference (income - expenses)
	difference := totalIncome - totalExpenses

	// Get category-wise report
	var categoryReport []models.CategoryReport
	err = configs.DB.Raw("SELECT category, type, COALESCE(SUM(amount), 0) AS total FROM transactions WHERE user_id = ? AND date >= ? AND date <= ? GROUP BY category, type", userID, startDate, endDate).Scan(&categoryReport).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to generate category report")
	}

	report := models.Report{
		TotalExpenses:  totalExpenses,
		TotalIncome:    totalIncome,
		Difference:     difference,
		CategoryReport: categoryReport,
	}

	return c.JSON(http.StatusOK, report)
}
