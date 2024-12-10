package services

import (
	"car_rental/server"
	"car_rental/server/models"
	"car_rental/server/repositories"
	"log"
	"regexp"
	"time"
)

type OrderService struct {
	OrderRepo *repositories.OrderRepository
}

func (s *OrderService) AddOrder(dataMap map[string]string) error {
	var err error
	data := &models.Order{}

	data.OrderDate, err = time.Parse(time.RFC1123, dataMap["order_date"])
	if err != nil {
		log.Println("[error][AddOrder_BE][handlers][1] parse time OrderDate error, err:", err)
		return server.ErrBadRequest
	}
	data.PickupDate, err = time.Parse(time.RFC1123, dataMap["pickup_date"])
	if err != nil {
		log.Println("[error][AddOrder_BE][handlers][2] parse time PickupDate error, err:", err)
		return server.ErrBadRequest
	}
	data.DropoffDate, err = time.Parse(time.RFC1123, dataMap["dropoff_date"])
	if err != nil {
		log.Println("[error][AddOrder_BE][handlers][3] parse time DropoffDate error, err:", err)
		return server.ErrBadRequest
	}

	return s.OrderRepo.AddOrder(data)
}

func (s *OrderService) GetOrders(dataMap map[string]string) (any, error) {
	isMatch, _ := regexp.MatchString("^[1-9][0-9]*$", dataMap["id"])
	if !isMatch {
		log.Println("[error][AddOrder_BE][handlers][1] invalid id regex, err:", dataMap["id"])
		return nil, server.ErrBadRequest
	}

	data, err := s.OrderRepo.GetOrderById(dataMap["id"])
	return *data, err
}

func (s *OrderService) DeleteOrder(dataMap map[string]string) error {
	isMatch, _ := regexp.MatchString("^[1-9][0-9]*$", dataMap["id"])
	if !isMatch {
		log.Println("[error][AddOrder_BE][handlers][1] invalid id regex, err:", dataMap["id"])
		return server.ErrBadRequest
	}

	return s.OrderRepo.DeleteOrder(dataMap["id"])
}
