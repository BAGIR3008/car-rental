package repositories

import (
	"car_rental/server/models"
	"log"

	"gorm.io/gorm"
)

type CarRepository struct {
	DB *gorm.DB
}

func (r *CarRepository) AddCar(data *models.Car) error {
	tx := r.DB.Create(data)
	if tx.Error != nil {
		log.Println("[error][AddCar_BE][repositories][1] save")
		return tx.Error
	}

	return nil
}

func (r *CarRepository) GetCarById(id string) (*models.Car, error) {
	var data *models.Car

	tx := r.DB.Where("id = ?", id).First(data)
	if tx.Error != nil {
		log.Println("[error][GetCar_BE][repositories][1] err:", tx.Error)
		return nil, tx.Error
	}

	return data, nil
}

func (r *CarRepository) GetListCars() (*[]models.Car, error) {
	data := &[]models.Car{}

	tx := r.DB.Find(data)
	if tx.Error != nil {
		log.Println("[error][GetListCars_BE][repositories][1] err:", tx.Error)
		return nil, tx.Error
	}

	return data, nil
}

func (r *CarRepository) UpdateCar(id string, data *models.Car) error {
	tx := r.DB.Where("id = ?", id).Updates(*data)
	if tx.Error != nil {
		log.Println("[error][GetListCars_BE][repositories][1] err:", tx.Error)
		return tx.Error
	}

	return nil
}

func (r *CarRepository) DeleteCar(id string) error {
	data := &models.Car{}

	tx := r.DB.Where("id = ?", id).Delete(data)
	if tx.Error != nil {
		log.Println("[error][GetListCars_BE][repositories][1] err:", tx.Error)
		return tx.Error
	}

	return nil
}
