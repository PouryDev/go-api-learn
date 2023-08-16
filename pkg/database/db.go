package database

import (
	"api-learn/internal/api/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func (s *Service) NewService() {
	dsn := "root:chaghalnameh@tcp(127.0.0.1:3306)/video_challenge_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	err = db.AutoMigrate(&models.Wallet{})
	if err != nil {
		logrus.Warn(err)
	}
	err = db.AutoMigrate(&models.Transaction{})
	if err != nil {
		logrus.Warn(err)
	}

	s.DB = db
}
