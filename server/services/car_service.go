package services

import (
	"car_rental/server"
	"car_rental/server/models"
	"car_rental/server/repositories"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type CarService struct {
	CarRepo *repositories.CarRepository
}

func (s *CarService) AddCar(dataMap map[string]any) error {
	var err error
	data := &models.Car{}

	data.CarName = strings.TrimSpace(dataMap["name"].(string))
	data.DayRate, err = strconv.ParseFloat(dataMap["date_rate"].(string), 64)
	if err != nil {
		log.Println("[error][AddCar_BE][handlers][1] parseFloat DayRate error, err:", err)
		return server.ErrBadRequest
	}
	data.DayMonth, err = strconv.ParseFloat(dataMap["day_month"].(string), 64)
	if err != nil {
		log.Println("[error][AddCar_BE][handlers][2] parseFloat DayMonth error, err:", err)
		return server.ErrBadRequest
	}

	data.Image = strings.TrimSpace(dataMap["image"].(string))
	return s.CarRepo.AddCar(data)
}

func (s *CarService) GetCars(dataMap map[string]string) (any, error) {
	isMatch, _ := regexp.MatchString("^[1-9][0-9]*$", dataMap["id"])
	if !isMatch {
		return nil, server.ErrBadRequest
	}

	data, err := s.CarRepo.GetCarById(dataMap["id"])
	return *data, err
}

func (s *CarService) DeleteCar(dataMap map[string]string) error {
	isMatch, _ := regexp.MatchString("^[1-9][0-9]*$", dataMap["id"])
	if !isMatch {
		return server.ErrBadRequest
	}

	return s.CarRepo.DeleteCar(dataMap["id"])
}
