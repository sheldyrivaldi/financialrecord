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

func AdminGetUsers(c echo.Context) error {
	var users []models.Users
	result := configs.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "User not found", Data: nil})
	} else {
		return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: users})
	}
}

func AdminGetUserById(c echo.Context) error {
	errEnv := godotenv.Load()
	if errEnv != nil {
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

	role := claims.Role

	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{Status: false, Message: "Unauthorized", Data: nil})
	}
	var users []models.Users
	userID := c.Param("id")
	result := configs.DB.Where("id = ?", userID).Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "User not found", Data: nil})
	} else {
		return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: users})
	}
}

func UserGetProfile(c echo.Context) error {
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

	var users []models.Users
	result := configs.DB.Where("id = ?", userID).Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "User not found", Data: nil})
	} else {
		return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: users})
	}
}

func UserUpdateProfile(c echo.Context) error {
	// Get JWT token from request header
	authToken := c.Request().Header.Get("Authorization")
	splitString := strings.Split(authToken, "Bearer ")
	tokenString := splitString[1]

	// Parse JWT token
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid token")
	}

	claims, ok := token.Claims.(*models.Claims)
	if !ok || !token.Valid {
		return c.JSON(http.StatusUnauthorized, "Invalid token")
	}

	id := strconv.Itoa(claims.ID)
	var user models.Users
	if err := configs.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := configs.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update user")
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Updated successfully", Data: user})

}
