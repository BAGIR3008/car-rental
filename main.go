package main

import (
	"car_rental/config"
	"car_rental/server/handlers"
	"car_rental/server/repositories"
	"car_rental/server/routes"
	"car_rental/server/services"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitSQL()
	if db == nil {
		log.Println("[error][InitSQL_BE][db][1] db nil")
		return
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Welcome to CAR RENTAL!")
	})

	carRepo := &repositories.CarRepository{DB: db}
	carService := &services.CarService{CarRepo: carRepo}
	carHandler := &handlers.CarHandler{CarSevice: carService}
	routes.SetupCarRoutes(e, carHandler)

	orderRepo := &repositories.OrderRepository{DB: db}
	orderService := &services.OrderService{OrderRepo: orderRepo}
	orderHandler := &handlers.OrderHandler{OrderSevice: orderService}
	routes.SetupOrderRoutes(e, orderHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
