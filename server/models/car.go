package models

import "gorm.io/gorm"

type Car struct {
	CarId     uint           `gorm:"unique;primaryKey;autoIncrement;not null"`
	CarName   string         `gorm:"not null"`
	DayRate   float64        `gorm:"not null"`
	DayMonth  float64        `gorm:"not null"`
	Image     string         `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
