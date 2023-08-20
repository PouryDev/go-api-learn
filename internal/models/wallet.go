package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	PhoneNumber string `gorm:"unique;not null"`
	Balance     int64  `gorm:"default:0;not null"`
}
