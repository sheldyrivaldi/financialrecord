package middlewares

import (
	"log"
	"net/http"
	"os"
	"prakerja/models"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func AdminAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		errEnv := godotenv.Load()
		if errEnv != nil {
			log.Fatal("Error loading .env file")
		}

		// Get JWT token from request header
		authToken := c.Request().Header.Get("Authorization")
		if authToken == "" {
			return c.JSON(http.StatusUnauthorized, "Unauthorized or invalid token")
		}
		splitString := strings.Split(authToken, "Bearer ")
		tokenString := splitString[1]

		// Parse JWT token
		JWT_KEY := os.Getenv("JWT_KEY")
		token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_KEY), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Unauthorized or invalid token")
		}

		claims, ok := token.Claims.(*models.Claims)
		if !ok || !token.Valid {
			return c.JSON(http.StatusUnauthorized, "Unauthorized or invalid token")
		}
		if claims.Role != "admin" {
			return c.JSON(http.StatusUnauthorized, "Unauthorized or invalid token")
		}

		return next(c)
	}
}

func UserAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get JWT token from request header
		authToken := c.Request().Header.Get("Authorization")
		if authToken == "" {
			return c.JSON(http.StatusUnauthorized, "Unauthorized or invalid token")
		}
		splitString := strings.Split(authToken, "Bearer ")
		tokenString := splitString[1]

		// Parse JWT token
		JWT_KEY := os.Getenv("JWT_KEY")
		token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_KEY), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Unauthorized or invalid token")
		}

		claims, ok := token.Claims.(*models.Claims)
		if !ok || !token.Valid {
			return c.JSON(http.StatusUnauthorized, "Unauthorized or invalid token")
		}
		if claims.Role != "user" {
			return c.JSON(http.StatusUnauthorized, "Unauthorized or invalid token")
		}

		return next(c)
	}
}
