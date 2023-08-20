package models

import "gorm.io/gorm"

const (
	GIFT_TYPE    = "gift"
	GATEWAY_TYPE = "gateway"
	ADMIN_TYPE   = "admin"
	DEFAULT_TYPE = "default"
	OPEN_STATUS  = "open"
	PAID_STATUS  = "paid"
)

type TransactionType string

func (t TransactionType) String() string {
	str := string(t)
	if str != GIFT_TYPE && str != GATEWAY_TYPE && str != ADMIN_TYPE {
		str = DEFAULT_TYPE
	}
	return str
}

type TransactionStatus string

func (t TransactionStatus) String() string {
	str := string(t)
	if str != PAID_STATUS {
		str = OPEN_STATUS
	}

	return str
}

type Transaction struct {
	gorm.Model
	Title          *string
	Amount         int64             `gorm:"not null"`
	Type           TransactionType   `gorm:"not null"`
	PhoneNumber    string            `gorm:"not null"`
	Status         TransactionStatus `gorm:"default:open; not null"`
	DiscountCodeID *uint
	DiscountCode   *DiscountCode `gorm:"foreignKey:DiscountCodeID"`
}
