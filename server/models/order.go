package models

import "time"

type Order struct {
	OrderId         uint      `gorm:"unique;primaryKey;autoIncrement;not null"`
	CarId           uint      `gorm:"not null"`
	OrderDate       time.Time `gorm:"not null"`
	PickupDate      time.Time `gorm:"not null"`
	DropoffDate     time.Time `gorm:"not null"`
	PickupLocation  string    `gorm:"not null"`
	DropoffLocation string    `gorm:"not null"`
}
