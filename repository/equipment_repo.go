package repository

import (
	"sport-rental/models"

	"gorm.io/gorm"
)

// Contract module
type EquipmentRepo interface {
	Create(equipment models.Equipment) (models.Equipment, error)
	GetAll() ([]models.Equipment, error)
	GetByID(id uint) (models.Equipment, error)
	Update(id uint, equipment models.Equipment) (models.Equipment, error)
	Delete(id uint) error
}

type equipmentRepo struct {
	db *gorm.DB
}

// Constructor for repo init
func NewEquipmentRepo(db *gorm.DB) EquipmentRepo {
	return &equipmentRepo{db: db}
}

func (r *equipmentRepo) Create(eq models.Equipment) (models.Equipment, error) {
	err := r.db.Create(&eq).Error
	return eq, err
}

func (r *equipmentRepo) GetAll() ([]models.Equipment, error) {
	var equipments []models.Equipment
	err := r.db.Find(&equipments).Error
	return equipments, err
}

func (r *equipmentRepo) GetByID(id uint) (models.Equipment, error) {
	var eq models.Equipment
	err := r.db.First(&eq, id).Error
	return eq, err
}

func (r *equipmentRepo) Update(id uint, eq models.Equipment) (models.Equipment, error) {
	var existing models.Equipment
	if err := r.db.First(&existing, id).Error; err != nil {
		return models.Equipment{}, err
	}
	err := r.db.Model(&existing).Updates(eq).Error
	return existing, err
}

func (r *equipmentRepo) Delete(id uint) error {
	return r.db.Delete(&models.Equipment{}, id).Error
}
