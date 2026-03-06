package repository

import (
	"sport-rental/models"

	"gorm.io/gorm"
)

type TopUpRepo interface {
	Create(topup models.TopUp) (models.TopUp, error)
	GetByOrderID(orderID string) (models.TopUp, error)
	UpdateStatus(orderID string, status string) error
}

type topUpRepo struct {
	db *gorm.DB
}

func NewTopUpRepo(db *gorm.DB) TopUpRepo {
	return &topUpRepo{db}
}

func (r *topUpRepo) Create(topup models.TopUp) (models.TopUp, error) {
	err := r.db.Create(&topup).Error
	return topup, err
}

func (r *topUpRepo) GetByOrderID(orderID string) (models.TopUp, error) {
	var topup models.TopUp
	err := r.db.Where("order_id = ?", orderID).First(&topup).Error
	return topup, err
}

func (r *topUpRepo) UpdateStatus(orderID string, status string) error {
	return r.db.Model(&models.TopUp{}).Where("order_id = ?", orderID).Update("status", status).Error
}
