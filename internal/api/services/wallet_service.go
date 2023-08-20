package services

import (
	"gorm.io/gorm"
	"video/internal/api/models"
	"video/pkg/database"
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
	wallet := models.Wallet{PhoneNumber: userPhone, Balance: 0}
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

func (ws WalletService) GetUserBalance(userPhone string) (*int64, error) {
	var wallet models.Wallet
	res := ws.DB.Where("phone_number = ?", userPhone).First(&wallet)
	if res.Error != nil {
		return nil, res.Error
	}

	return &wallet.Balance, nil
}

func (ws WalletService) Find(userPhone string) (*models.Wallet, error) {
	var wallet models.Wallet
	res := ws.DB.Model(&models.Wallet{}).Where("phone_number = ?", userPhone).First(&wallet)
	if res.Error != nil {
		return nil, res.Error
	}

	return &wallet, nil
}

// TODO add statistics
