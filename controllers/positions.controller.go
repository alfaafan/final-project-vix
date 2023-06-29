package controllers

import (
	"finalProject/db"
	"finalProject/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePosition(c echo.Context) error {
	db := db.DB()
	p := new(models.Positions)

	if err := c.Bind(p); err != nil {
		data := map[string]interface{} {
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	db.Create(p)

	response := map[string]interface{} {
		"success": true,
		"message": "Success",
		"data": p,
	}
	return c.JSONPretty(http.StatusOK, response, "	")
}

func GetAllPositions(c echo.Context) error {
	db := db.DB()
	var positions []models.Positions

	err := db.Find(&positions).Error
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Internal server error"})
	}

	response := map[string]interface{} {
		"success": true,
		"message": "Success",
		"data": positions,
	}

	return c.JSON(http.StatusOK, response)
}

func GetPosition(c echo.Context) error {
	db := db.DB()
	position := new(models.Positions)

	err := db.First(position, c.Param("id")).Error
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Position not found"})
	}

	response := map[string]interface{} {
		"success": true,
		"message": "Success",
		"data": position,
	}

	return c.JSON(http.StatusOK, response)
}

func DeletePosition(c echo.Context) error {
	db := db.DB()
	position := new(models.Positions)

	err := db.First(position, c.Param("id")).Error
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Position not found"})
	}

	db.Delete(position)

	response := map[string]interface{} {
		"success": true,
		"message": "Success",
		"data": "",
	}

	return c.JSON(http.StatusOK, response)
}