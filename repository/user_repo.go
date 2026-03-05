package repository

import (
	"sport-rental/models"

	"gorm.io/gorm"
)

// Contract module
type UserRepo interface {
	Create(user models.User) (models.User, error)
	GetByEmail(email string) (models.User, error)
	UpdateBalance(userID int, newAmount float64) error
	GetByID(id uint) (models.User, error)
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

// UpdateBalance - Update user balance
func (r *userRepo) UpdateBalance(userID int, newAmount float64) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).
		Update("deposit_amount", gorm.Expr("deposit_amount + ?", newAmount)).Error
}

// GetByID - Find user based on ID
func (r *userRepo) GetByID(id uint) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return user, err
}
