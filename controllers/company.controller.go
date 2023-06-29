package controllers

import (
	"errors"
	"finalProject/db"
	"finalProject/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateCompany(c echo.Context) error {
	db := db.DB()
	db.AutoMigrate(&models.Company{})

	company :=new(models.Company)

	if err := c.Bind(company); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	db.Create(company)

	response := map[string]interface{} {
		"success": true,
		"message": "Success",
		"data": company,
	}

	return c.JSON(http.StatusOK, response)
}

func TopUpBalance(c echo.Context) error {
	db := db.DB()

	company := new(models.Company)

	err := db.First(company, c.Param("id")).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(404, map[string]string{"error": "Company not found"})
		}
		return c.JSON(500, map[string]string{"error": "Internal Server Error"})
	}

	response := map[string]interface{} {
		"success": true,
		"message": "Success",
		"data": company,
	}

	return c.JSON(http.StatusOK, response)
}

func GetCompany(c echo.Context) error {
	db := db.DB()
	var company []models.Company

	err := db.Find(&company).Error
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Internal Server Error"})
	}

	response := map[string]interface{} {
		"success": true,
		"message": "Success",
		"data": company,
	}

	return c.JSON(http.StatusOK, response)
}