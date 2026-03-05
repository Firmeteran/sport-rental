package repository

import (
	"sport-rental/models"
	"time"

	"gorm.io/gorm"
)

type RentalRepo interface {
	CreateRental(rental models.RentalHistory, cost float64) (models.RentalHistory, error)
	GetRentalByUserID(userID int) ([]models.RentalHistory, error)
	UpdateReturn(rentalID uint, equipmentID uint) error
	GetByID(id uint) (models.RentalHistory, error)
}

type rentalRepo struct {
	db *gorm.DB
}

// Constructor for repo init
func NewRentalRepo(db *gorm.DB) RentalRepo {
	return &rentalRepo{db: db}
}

func (r *rentalRepo) CreateRental(rental models.RentalHistory, cost float64) (models.RentalHistory, error) {
	// Using transaction for data consistency
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// Reducing user deposit balance
		if err := tx.Model(&models.User{}).Where("id = ?", rental.UserID).
			Update("deposit_amount", gorm.Expr("deposit_amount - ?", cost)).Error; err != nil {
			return err
		}

		// Reducing stock
		if err := tx.Model(&models.Equipment{}).Where("id = ?", rental.EquipmentID).
			Update("stock_availability", gorm.Expr("stock_availability - ?", 1)).Error; err != nil {
			return err
		}

		// Save rental history
		if err := tx.Create(&rental).Error; err != nil {
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

func (r *rentalRepo) UpdateReturn(rentalID uint, equipmentID uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Update ReturnDate in rental_histories table
		now := time.Now()
		if err := tx.Model(&models.RentalHistory{}).Where("id = ?", rentalID).
			Update("return_date", &now).Error; err != nil {
			return err
		}

		// Adding back stock in equipments table (+ 1)
		if err := tx.Model(&models.Equipment{}).Where("id = ?", equipmentID).
			Update("stock_availability", gorm.Expr("stock_availability + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
}

// Search ID for validation in the service level
func (r *rentalRepo) GetByID(id uint) (models.RentalHistory, error) {
	var rental models.RentalHistory
	err := r.db.First(&rental, id).Error
	return rental, err
}
