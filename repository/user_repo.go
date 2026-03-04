package repository

import (
	"sport-rental/models"

	"gorm.io/gorm"
)

// Contract module
type UserRepo interface {
	Create(user models.User) (models.User, error)
	GetByEmail(email string) (models.User, error)
	UpdateDeposit(userID int, newAmount float64) error
}

type userRepo struct {
	db *gorm.DB
}

// Constructor for repo
func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

// Create - Save new user to database
func (r *userRepo) Create(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

// GetByEmail - Find user based on email (for login/validation)
func (r *userRepo) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

// UpdateDeposit - Update user balance
func (r *userRepo) UpdateDeposit(userID int, newAmount float64) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("deposit_amount", newAmount).Error
}
