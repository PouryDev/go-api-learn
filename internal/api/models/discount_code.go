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
	Code      string
	MaxUsers  int
	UsedCount int
	Value     int64
	Type      DiscountCodeType
}
