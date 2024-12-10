package repositories

import (
	"car_rental/server/models"
	"log"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func (r *OrderRepository) AddOrder(data *models.Order) error {
	tx := r.DB.Create(data)
	if tx.Error != nil {
		log.Println("[error][AddCar_BE][repositories][1] save")
		return tx.Error
	}

	return nil
}

func (r *OrderRepository) GetOrderById(id string) (*models.Order, error) {
	var data *models.Order

	tx := r.DB.Where("id = ?", id).First(data)
	if tx.Error != nil {
		log.Println("[error][GetCar_BE][repositories][1] err:", tx.Error)
		return nil, tx.Error
	}

	return data, nil
}

func (r *OrderRepository) GetListOrders() (*[]models.Order, error) {
	data := &[]models.Order{}

	tx := r.DB.Find(data)
	if tx.Error != nil {
		log.Println("[error][GetListCars_BE][repositories][1] err:", tx.Error)
		return nil, tx.Error
	}

	return data, nil
}

func (r *OrderRepository) UpdateOrder(id string, data *models.Order) error {
	tx := r.DB.Where("id = ?", id).Updates(*data)
	if tx.Error != nil {
		log.Println("[error][GetListCars_BE][repositories][1] err:", tx.Error)
		return tx.Error
	}

	return nil
}

func (r *OrderRepository) DeleteOrder(id string) error {
	data := &models.Order{}

	tx := r.DB.Where("id = ?", id).Delete(data)
	if tx.Error != nil {
		log.Println("[error][GetListCars_BE][repositories][1] err:", tx.Error)
		return tx.Error
	}

	return nil
}
