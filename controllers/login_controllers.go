package controllers

import (
	"log"
	"net/http"
	"os"
	"prakerja/configs"
	"prakerja/models"
	"time"

	"github.com/joho/godotenv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	insertUser := models.Users{}
	if err := c.Bind(&insertUser); err != nil {
		return err
	}
	result := configs.DB.Where("email = ?", insertUser.Email).First(&insertUser)

	if result.Error != nil {
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{Status: false, Message: "Invalid email or password", Data: nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := &models.Claims{
		ID:   insertUser.Id,
		Role: insertUser.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token.Claims = claims

	// Generate encoded token and send it as response

	t, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Failed to generate token"})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "User successfully login!",
		Data:    t,
	})
}
