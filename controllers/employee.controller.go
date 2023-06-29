package controllers

import (
	"errors"
	"finalProject/db"
	"finalProject/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateEmployee(c echo.Context) error {
db := db.DB()
db.AutoMigrate(&models.Employees{})	

employee := new(models.Employees)

if err := c.Bind(employee); err != nil {
	return c.JSON(400, map[string]string{"error": err.Error()})
}

db.Preload("Positions").Create(employee)

response := map[string]interface{} {
	"success": true,
	"message": "Success",
	"data": employee,
}

return c.JSON(201, response)
}

func GetEmployee(c echo.Context) error {
	db := db.DB()
	employee := new(models.Employees)
	err := db.Preload("Positions").First(employee, c.Param("id")).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(404, map[string]string{"error": "Employee not found"})
		}
		return c.JSON(500, map[string]string{"error": "Internal server error"})
	}

	response := map[string]interface{} {
		"success": true,
		"message": "Success",
		"data": employee,
	}

	return c.JSON(200, response)
}

func GetAllEmployee(c echo.Context) error {
	db := db.DB()
	var employees []models.Employees

	err := db.Preload("Positions").Find(&employees).Error
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Internal server error"})
	}

	response := map[string]interface{} {
		"success": true,
		"message": "Success",
		"data": employees,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteEmployee(c echo.Context) error {
	db := db.DB()
	employee := new(models.Employees)

	err := db.First(employee, c.Param("id")).Error
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Employee not found"})
	}

	db.Delete(employee)

	response := map[string]interface{} {
		"success": true,
		"message": "Success",
		"data": "",
	}

	return c.JSON(http.StatusOK, response)
}

// func WithdrawSalary(c echo.Context) error {
	
// }