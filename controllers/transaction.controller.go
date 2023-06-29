package controllers

import (
	"finalProject/db"
	"finalProject/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTransaction(c echo.Context) error {
	db := db.DB()

	var transaction []models.Transaction

	err := db.Preload("Company").Find(&transaction).Error
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Internal Server Error"})
	}

	response := map[string]interface{} {
		"success": true,
		"message": "Success",
		"data": transaction,
	}

	return c.JSON(http.StatusOK, response)
}