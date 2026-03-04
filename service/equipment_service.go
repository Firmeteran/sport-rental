package service

import (
	"errors"
	"sport-rental/models"
	"sport-rental/repository"
)

type EquipmentService interface {
	AddEquipment(input models.Equipment) (models.Equipment, error)
	FetchAll() ([]models.Equipment, error)
}

type equipmentService struct {
	repo repository.EquipmentRepo
}

func NewEquipmentService(r repository.EquipmentRepo) EquipmentService {
	return &equipmentService{repo: r}
}

func (s *equipmentService) AddEquipment(input models.Equipment) (models.Equipment, error) {
	if input.StockAvailability < 0 {
		return models.Equipment{}, errors.New("Stock cannot be negative.")
	}
	return s.repo.Create(input)
}

func (s *equipmentService) FetchAll() ([]models.Equipment, error) {
	return s.repo.GetAll()
}
