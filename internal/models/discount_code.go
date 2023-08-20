package models

import "gorm.io/gorm"

const (
	GIFT_CODE     = "gift"
	DISCOUNT_CODE = "discount"
	DEFAULT_CODE  = "default"
)

type DiscountCodeType string

func (dct DiscountCodeType) String() string {
	str := string(dct)
	if str != GIFT_CODE && str != DISCOUNT_CODE {
		return DEFAULT_CODE
	}
	return str
}

type DiscountCode struct {
	gorm.Model
	Code         string           `gorm:"unique; not null"`
	MaxUsers     int              `gorm:"not null"`
	UsedCount    int              `gorm:"default:0"`
	Value        int64            `gorm:"not null"`
	Type         DiscountCodeType `gorm:"not null"`
	Transactions []Transaction
}
