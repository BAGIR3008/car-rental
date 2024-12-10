package routes

import (
	"car_rental/server/handlers"

	"github.com/labstack/echo/v4"
)

func SetupCarRoutes(e *echo.Echo, h *handlers.CarHandler) {
	e.POST("/car", h.AddCar())
}
