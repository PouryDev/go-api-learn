package services

import (
	"gorm.io/gorm"
	"video/internal/database"
	"video/internal/models"
)

type DiscountCodeService struct {
	DB *gorm.DB
}

func NewDiscountCodeService() *DiscountCodeService {
	return &DiscountCodeService{DB: database.NewDB()}
}

func (dcs DiscountCodeService) Create(code string, maxUsers int, value int64, dcType models.DiscountCodeType) (*models.DiscountCode, error) {
	dc := models.DiscountCode{
		Code:     code,
		MaxUsers: maxUsers,
		Value:    value,
		Type:     dcType,
	}
	res := dcs.DB.Create(&dc)
	if res.Error != nil {
		return nil, res.Error
	}

	return &dc, nil
}

func (dcs DiscountCodeService) Update(dc models.DiscountCode) error {
	res := dcs.DB.Save(dc)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (dcs DiscountCodeService) Delete(dc models.DiscountCode) error {
	res := dcs.DB.Delete(&dc)
	return res.Error
}

func (dcs DiscountCodeService) IncreaseUsage(dc *models.DiscountCode) error {
	usage := dc.UsedCount + 1
	res := dcs.DB.Save(dc)
	if res.Error != nil {
		return res.Error
	}
	dc.UsedCount = usage
	return nil
}

func (dcs DiscountCodeService) GetByCode(code string) (*models.DiscountCode, error) {
	var dc models.DiscountCode
	res := dcs.DB.Model(&models.DiscountCode{}).Where("`code` = ?", code).First(&dc)
	if res.Error != nil {
		return nil, res.Error
	}

	return &dc, nil
}

func (dcs DiscountCodeService) Find(id uint) (*models.DiscountCode, error) {
	var dc models.DiscountCode
	res := dcs.DB.Model(&models.DiscountCode{}).Where("`id` = ?", id).First(&dc)
	if res.Error != nil {
		return nil, res.Error
	}

	return &dc, nil
}
