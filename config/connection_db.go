package config

import (
	"car_rental/server/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitSQL() *gorm.DB {
	dsn := "host=localhost user=postgres password=123 dbname=car_rental port=5432 sslmode=disable TimeZone=UTC"
	// dsn := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("[error][InitSQL_BE][db][1]", err.Error())
		return nil
	}

	err = db.AutoMigrate(models.Car{}, models.Order{})
	if err != nil {
		log.Println("[error][InitSQL_BE][db][2]", err.Error())
		return nil
	}

	return db
}
