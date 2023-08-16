package services

import (
	"api-learn/internal/api/models"
	"api-learn/pkg/database"
	"gorm.io/gorm"
)

type WalletService struct {
	DB *gorm.DB
}

func NewWalletService() *WalletService {
	svc := database.Service{}
	svc.NewService()
	return &WalletService{DB: svc.DB}
}

func (ws WalletService) Create(userPhone string) (*models.Wallet, error) {
	wallet := models.Wallet{UserPhone: userPhone, Balance: 0}
	res := ws.DB.Create(&wallet)
	if res.Error != nil {
		return nil, res.Error
	}

	return &wallet, nil
}

func (ws WalletService) IncreaseBalance(wallet models.Wallet, amount int) error {
	balance := int64(amount) + wallet.Balance
	res := ws.DB.Model(&models.Wallet{}).Where("id = ?", wallet.ID).Update("balance", balance)

	return res.Error
}

func (ws WalletService) GetUserBalance(userPhone string) *int64 {
	var wallet models.Wallet
	res := ws.DB.Where("user_phone = ?", userPhone).First(&wallet)
	if res.Error != nil {
		return nil
	}

	return &wallet.Balance
}

// TODO add statistics
