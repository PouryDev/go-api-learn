package services

import (
	"gorm.io/gorm"
	"video/internal/models"
)

type TransactionService struct {
	db *gorm.DB
}

func (ts TransactionService) Create(amount int64, tType models.TransactionType, phone string, status models.TransactionStatus, title *string) (*models.Transaction, error) {
	t := models.Transaction{
		Amount:      amount,
		Type:        tType,
		PhoneNumber: phone,
		Status:      status,
		Title:       title,
	}
	res := ts.db.Create(&t)
	if res.Error != nil {
		return nil, res.Error
	}

	return &t, nil
}

func (ts TransactionService) Pay(t models.Transaction) error {
	t.Status = models.PAID_STATUS
	res := ts.db.Save(t)
	return res.Error
}
