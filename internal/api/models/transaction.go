package models

import "gorm.io/gorm"

const (
	GiftType    = "gift"
	GatewayType = "gateway"
	AdminType   = "admin"
	DefaultType = "default"
)

type TransactionType string

func (t TransactionType) String() string {
	str := string(t)
	if str != GiftType && str != GatewayType && str != AdminType {
		str = DefaultType
	}
	return str
}

type Transaction struct {
	gorm.Model
	Title     *string
	Amount    int64
	Type      TransactionType
	UserPhone string
}
