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

	company := new(models.Company)

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

	transaction := new(models.Transaction)
    if err := c.Bind(transaction); err != nil {
        return c.String(http.StatusBadRequest, "Invalid request")
    }

	tx := db.Begin()

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, err) 
	}

	company := new(models.Company)
	err := tx.First(&company, transaction.CompaniesID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return c.JSON(404, map[string]string{"error": "Company not found"})
		}
		tx.Rollback()
		return c.JSON(500, map[string]string{"error": "Internal Server Error"})
	}

	company.Balance += transaction.Amount

	if err := tx.Save(&company).Error; err != nil {
		tx.Rollback()
        return c.String(http.StatusInternalServerError, "Failed to update balance")
    }

	tx.Commit()

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