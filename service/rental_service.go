package service

import (
	"errors"
	"sport-rental/models"
	"sport-rental/repository"
	"time"
)

type RentalService interface {
	RentEquipment(userID, equipmentID int) (models.RentalHistory, error)
	ReturnEquipment(rentalID int) error
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

	// Due date (24 hour) from now
	dueDate := time.Now().Add(24 * time.Hour)

	// Create rental object
	newRental := models.RentalHistory{
		UserID:      uint(userID),
		EquipmentID: uint(equipmentID),
		RentDate:    time.Now(),
		DueDate:     dueDate,
	}

	return s.rentalRepo.CreateRental(newRental, equipment.RentalCosts)
}

func (s *rentalService) ReturnEquipment(rentalID int) error {
	// Search for rental data
	rental, err := s.rentalRepo.GetByID(uint(rentalID))
	if err != nil {
		return errors.New("Rental data cannot be found.")
	}

	// Validation: has it ever been returned?
	if rental.ReturnDate != nil {
		return errors.New("This tool has been returned previously.")
	}

	// Overdue logic
	now := time.Now()
	if now.After(rental.DueDate) {
		// Calculate the time in hours
		overdueHours := now.Sub(rental.DueDate).Hours()

		// Charge IDR 10.000 per hour of delay
		fineAmount := int(overdueHours) * 10000

		// Deduct user's balance for fines
		err := s.userRepo.UpdateBalance(int(rental.UserID), -float64(fineAmount))
		if err != nil {
			return errors.New("Failed to deduct fine from balance.")
		}
	}

	// Return execution through repo
	return s.rentalRepo.UpdateReturn(rental.ID, rental.EquipmentID)
}
