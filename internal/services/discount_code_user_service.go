package services

import (
	"errors"
	"gorm.io/gorm"
	"video/internal/database"
	"video/internal/models"
)

type DiscountCodeUserService struct {
	db *gorm.DB
}

func NewDiscountCodeUserService() *DiscountCodeUserService {
	return &DiscountCodeUserService{db: database.NewDB()}
}

func (ds DiscountCodeUserService) Create(codeID uint, phoneNumber string) error {
	du := models.DiscountCodeUser{
		DiscountCodeID: codeID,
		PhoneNumber:    phoneNumber,
	}
	res := ds.db.Model(&models.DiscountCodeUser{}).Create(&du)
	return res.Error
}

func (ds DiscountCodeUserService) CheckCodeForUser(codeID uint, phoneNumber string) (bool, error) {
	var count int64
	res := ds.db.Model(&models.DiscountCodeUser{}).
		Where("discount_code_id = ?", codeID).
		Where("phone_number = ?", phoneNumber).
		Count(&count)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, res.Error
	}

	if count != 0 {
		return false, errors.New("code has been used by the user before")
	}

	return true, nil
}
