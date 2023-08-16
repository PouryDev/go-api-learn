package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	UserPhone string
	Balance   int64
}
