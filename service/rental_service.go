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
	// Tool validation
	equipment, err := s.equipmentRepo.GetByID(uint(equipmentID))
	if err != nil || equipment.StockAvailability <= 0 {
		return models.RentalHistory{}, errors.New("Tool are out of stock.")
	}

	// Create rental object
	newRental := models.RentalHistory{
		UserID:      uint(userID),
		EquipmentID: uint(equipmentID),
		RentDate:    time.Now(),
	}

	return s.rentalRepo.CreateRental(newRental)
}
