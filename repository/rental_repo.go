package repository

import (
	"sport-rental/models"

	"gorm.io/gorm"
)

type RentalRepo interface {
	CreateRental(rental models.RentalHistory) (models.RentalHistory, error)
	GetRentalByUserID(userID int) ([]models.RentalHistory, error)
}

type rentalRepo struct {
	db *gorm.DB
}

// Constructor for repo init
func NewRentalRepo(db *gorm.DB) RentalRepo {
	return &rentalRepo{db: db}
}

func (r *rentalRepo) CreateRental(rental models.RentalHistory) (models.RentalHistory, error) {
	// Using transaction for data consistency
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// Save rental history
		if err := tx.Create(&rental).Error; err != nil {
			return err
		}

		// Reducing stock
		if err := tx.Model(&models.Equipment{}).Where("id = ?", rental.EquipmentID).
			Update("stock_availability", gorm.Expr("stock_availability - ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
	return rental, err
}

func (r *rentalRepo) GetRentalByUserID(userID int) ([]models.RentalHistory, error) {
	var history []models.RentalHistory
	// Preload is used to retrieve Equipment data simultaneously
	err := r.db.Preload("Equipment").Where("user_id = ?", userID).Find(&history).Error
	return history, err
}
