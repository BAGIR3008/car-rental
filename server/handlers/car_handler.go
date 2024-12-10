package handlers

import (
	"car_rental/server/services"
	"log"

	"github.com/labstack/echo/v4"
)

type CarHandler struct {
	CarSevice *services.CarService
}

func (h *CarHandler) AddCar() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := h.CarSevice.AddCar(map[string]any{
			"name":       c.FormValue("name"),
			"day_rate":   c.FormValue("day_rate"),
			"month_rate": c.FormValue("month_rate"),
			"image":      c.FormValue("image"),
		})

		if err != nil {
			log.Println("[error][AddCar_BE][handlers][1] call AddCar error, err:", err)
			return echo.ErrInternalServerError
		}

		return c.JSON(200, map[string]any{
			"code":    "200",
			"message": "success",
		})
	}
}

func (h *CarHandler) GetCars() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := h.CarSevice.GetCars(map[string]string{
			"id": c.FormValue("id"),
		})

		if err != nil {
			log.Println("[error][AddCar_BE][handlers][1] call AddCar error, err:", err)
			return echo.ErrInternalServerError
		}

		return c.JSON(200, map[string]any{
			"code":    "200",
			"message": "success",
			"data":    data,
		})
	}
}

func (h *CarHandler) DeleteCar() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := h.CarSevice.DeleteCar(map[string]string{
			"id": c.FormValue("id"),
		})

		if err != nil {
			log.Println("[error][AddCar_BE][handlers][1] call AddCar error, err:", err)
			return echo.ErrInternalServerError
		}

		return c.JSON(200, map[string]any{
			"code":    "200",
			"message": "success",
		})
	}
}
