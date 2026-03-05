package models

import (
	"time"

	"gorm.io/gorm"
)

// User
type User struct {
	gorm.Model
	Username      string  `gorm:"unique;not null" json:"username"`
	Email         string  `gorm:"unique;not null" json:"email"`
	Password      string  `gorm:"not null" json:"-"`
	DepositAmount float64 `gorm:"default:0" json:"deposit_amount"`
}

// Equipment
type Equipment struct {
	gorm.Model
	Name              string  `gorm:"not null" json:"name"`
	Category          string  `gorm:"not null" json:"category"`
	RentalCosts       float64 `gorm:"not null" json:"rental_costs"`
	StockAvailability int     `gorm:"not null" json:"stock_availability"`
}

// Rental history
type RentalHistory struct {
	gorm.Model
	UserID      uint       `json:"user_id"`
	EquipmentID uint       `json:"equipment_id"`
	DueDate     time.Time  `json:"due_date"`
	RentDate    time.Time  `json:"rent_date"`
	ReturnDate  *time.Time `json:"return_date"`

	User      User      `gorm:"foreignKey:UserID"`
	Equipment Equipment `gorm:"foreignKey:EquipmentID"`
}
