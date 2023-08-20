package models

import "gorm.io/gorm"

const (
	GIFT_TYPE    = "gift"
	GATEWAY_TYPE = "gateway"
	ADMIN_TYPE   = "admin"
	DEFAULT_TYPE = "default"
)

type TransactionType string

func (t TransactionType) String() string {
	str := string(t)
	if str != GIFT_TYPE && str != GATEWAY_TYPE && str != ADMIN_TYPE {
		str = DEFAULT_TYPE
	}
	return str
}

type Transaction struct {
	gorm.Model
	Title     *string
	Amount    int64           `gorm:"not null"`
	Type      TransactionType `gorm:"not null"`
	UserPhone string          `gorm:"not null"`
}
