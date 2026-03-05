package service

import (
	"errors"
	"sport-rental/models"
	"sport-rental/repository"
	"time"
)

type RentalService interface {
	RentEquipment(userID, equipmentID int) (models.RentalHistory, error)
}

type rentalService struct {
	rentalRepo    repository.RentalRepo
	userRepo      repository.UserRepo
	equipmentRepo repository.EquipmentRepo
}

func NewRentalService(r repository.RentalRepo, u repository.UserRepo, e repository.EquipmentRepo) RentalService {
	return &rentalService{r, u, e}
}

func (s *rentalService) RentEquipment(userID, equipmentID int) (models.RentalHistory, error) {
	// Tool & stock validation
	equipment, err := s.equipmentRepo.GetByID(uint(equipmentID))
	if err != nil {
		return models.RentalHistory{}, errors.New("Equipment is not found.")
	}
	if equipment.StockAvailability <= 0 {
		return models.RentalHistory{}, errors.New("Equipment are out of stock.")
	}

	// User and balance (deposit) validation
	user, err := s.userRepo.GetByID(uint(userID))
	if err != nil {
		return models.RentalHistory{}, errors.New("User is not found.")
	}
	if user.DepositAmount < equipment.RentalCosts {
		return models.RentalHistory{}, errors.New("Insufficient deposit balance.")
	}

	// Create rental object
	newRental := models.RentalHistory{
		UserID:      uint(userID),
		EquipmentID: uint(equipmentID),
		RentDate:    time.Now(),
	}

	return s.rentalRepo.CreateRental(newRental)
}
