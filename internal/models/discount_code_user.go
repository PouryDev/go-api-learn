package models

import "gorm.io/gorm"

type DiscountCodeUser struct {
	gorm.Model
	PhoneNumber    string       `gorm:"uniqueIndex:idx_phone_discount; not null"`
	DiscountCodeID uint         `gorm:"uniqueIndex:idx_phone_discount; not null"`
	DiscountCode   DiscountCode `gorm:"foreignKey:DiscountCodeID"`
}
