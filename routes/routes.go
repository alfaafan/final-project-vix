package routes

import (
	"finalProject/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome")
	})

	positionsRoute := e.Group("/positions")
	positionsRoute.POST("/", controllers.CreatePosition)
	positionsRoute.GET("/", controllers.GetAllPositions)
	positionsRoute.GET("/:id", controllers.GetPosition)
	positionsRoute.DELETE("/:id", controllers.DeletePosition)

	employeeRoute := e.Group("/employee")
	employeeRoute.POST("/", controllers.CreateEmployee)
	employeeRoute.GET("/:id", controllers.GetEmployee)
	employeeRoute.GET("/", controllers.GetAllEmployee)
	employeeRoute.DELETE("/:id", controllers.DeleteEmployee)

	return e
}
