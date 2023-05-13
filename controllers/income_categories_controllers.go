package controllers

import (
	"log"
	"net/http"
	"os"
	"prakerja/configs"
	"prakerja/models"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func AdminGetIncomeCategoreies(c echo.Context) error {
	getCategories := []models.IncomeCategories{}
	result := configs.DB.Find(&getCategories)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Failed to get categories", Data: nil})
	} else {
		return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: getCategories})
	}
}
func AdminGetIncomeCategoreiesByUserID(c echo.Context) error {
	user_id := c.Param("user_id")
	getCategory := []models.IncomeCategories{}
	result := configs.DB.Where("user_id = ?", user_id).Find(&getCategory)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Category not found", Data: nil})
	} else {
		return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: getCategory})
	}
}

func UserGetIncomeCategories(c echo.Context) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get JWT token from request header
	authToken := c.Request().Header.Get("Authorization")
	splitString := strings.Split(authToken, "Bearer ")
	tokenString := splitString[1]

	// Parse JWT token
	token, _ := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	claims, _ := token.Claims.(*models.Claims)

	userID := strconv.Itoa(claims.ID)

	var categories []models.IncomeCategories
	result := configs.DB.Where("user_id = ?", userID).Find(&categories)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Category not found", Data: nil})
	} else {
		return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: categories})
	}
}

func UserCreateIncomeCategory(c echo.Context) error {
	// Get JWT token from request header
	authToken := c.Request().Header.Get("Authorization")
	splitString := strings.Split(authToken, "Bearer ")
	tokenString := splitString[1]

	// Parse JWT token
	token, _ := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	claims, _ := token.Claims.(*models.Claims)

	userID := claims.ID

	insertCategory := models.IncomeCategories{
		User_ID: userID,
	}
	c.Bind(&insertCategory)
	result := configs.DB.Create(&insertCategory)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Failed to insert category", Data: nil})
	} else {
		return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: insertCategory})
	}
}

func UserUpdateIncomeCategoryById(c echo.Context) error {
	// Get JWT token from request header
	authToken := c.Request().Header.Get("Authorization")
	splitString := strings.Split(authToken, "Bearer ")
	tokenString := splitString[1]

	// Parse JWT token
	token, _ := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	claims, _ := token.Claims.(*models.Claims)

	userID := claims.ID
	id := c.Param("id")

	var updateCategory models.IncomeCategories
	if err := configs.DB.Where("id = ? AND user_id = ?", id, userID).First(&updateCategory).Error; err != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{Status: false, Message: "Category not found", Data: nil})
	}

	if err := c.Bind(&updateCategory); err != nil {
		return err
	}

	if err := configs.DB.Save(&updateCategory).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Failed to update category", Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Category updated successfully", Data: updateCategory})
}

func UserDeleteIncomeCategoryById(c echo.Context) error {
	// Get JWT token from request header
	authToken := c.Request().Header.Get("Authorization")
	splitString := strings.Split(authToken, "Bearer ")
	tokenString := splitString[1]

	// Parse JWT token
	token, _ := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	claims, _ := token.Claims.(*models.Claims)

	userID := claims.ID
	id := c.Param("id")

	var deleteCategory models.IncomeCategories
	if err := configs.DB.Where("id = ? AND user_id = ?", id, userID).First(&deleteCategory).Error; err != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{Status: false, Message: "Category not found", Data: nil})
	}

	if err := c.Bind(&deleteCategory); err != nil {
		return err
	}

	if err := configs.DB.Delete(&deleteCategory).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Failed to delete category", Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Category deleted successfully", Data: deleteCategory})
}
