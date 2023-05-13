package controllers

import (
	"net/http"
	"prakerja/configs"
	"prakerja/models"

	"github.com/labstack/echo/v4"
)

func Registration(c echo.Context) error {
	insertUser := models.Users{}
	c.Bind(&insertUser)
	result := configs.DB.Create(&insertUser)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Failed to insert user", Data: nil})
	} else {
		return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: insertUser})
	}
}
