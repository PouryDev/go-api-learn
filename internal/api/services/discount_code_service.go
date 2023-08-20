package services

import (
	"gorm.io/gorm"
	"video/internal/api/models"
	"video/pkg/database"
)

type DiscountCodeService struct {
	DB *gorm.DB
}

func NewDiscountCodeService() *DiscountCodeService {
	svc := database.Service{}
	svc.NewService()
	return &DiscountCodeService{DB: svc.DB}
}

func (dcs DiscountCodeService) Create(code string, maxUsers int, userCount int, value int64, dcType models.DiscountCodeType) (*models.DiscountCode, error) {
	dc := models.DiscountCode{
		Code:      code,
		MaxUsers:  maxUsers,
		UsedCount: userCount,
		Value:     value,
		Type:      dcType,
	}
	res := dcs.DB.Create(&dc)
	if res.Error != nil {
		return nil, res.Error
	}

	return &dc, nil
}

func (dcs DiscountCodeService) Update(dc *models.DiscountCode, maxUsers int, usedCount int, value int64) error {
	res := dcs.DB.Model(&models.DiscountCode{}).Where("id = ?", dc.ID).Updates(map[string]interface{}{
		"max_users":   maxUsers,
		"users_count": usedCount,
		"value":       value,
	})
	if res.Error != nil {
		return res.Error
	}

	dc.MaxUsers = maxUsers
	dc.UsedCount = usedCount
	dc.Value = value
	return nil
}

func (dcs DiscountCodeService) Delete(dc models.DiscountCode) error {
	res := dcs.DB.Delete(&dc)
	return res.Error
}

func (dcs DiscountCodeService) IncreaseUsage(dc *models.DiscountCode) error {
	usage := dc.UsedCount + 1
	res := dcs.DB.Model(models.DiscountCode{}).Where("id = ?", dc.ID).Update("used_count", usage)
	if res.Error != nil {
		return res.Error
	}
	dc.UsedCount = usage
	return nil
}
