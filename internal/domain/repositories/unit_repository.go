package repositories

import (
	"github.com/fseda/cookbooked-api/internal/domain/models"
	"gorm.io/gorm"
)

func NewUnitRepository(db *gorm.DB) UnitRepository {
	return &unitRepository{db}
}

type UnitRepository interface {
	FindByID(id uint) (*models.Unit, error)
	ExistsAllIn(ids []uint) (bool, error)
	InvalidIDs(ids []uint) (invalidIDs []uint, err error)
}

type unitRepository struct {
	db *gorm.DB
}

func (ir *unitRepository) FindByID(id uint) (*models.Unit, error) {
	var unit models.Unit
	result := ir.db.First(&unit, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &unit, nil
}

// ids must be unique
func (ir *unitRepository) ExistsAllIn(ids []uint) (bool, error) {
	for _, id := range ids {
		err := ir.db.First(&models.Unit{}, id).Error
		if err == nil {
			continue
		} 
		if err == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, err
		}
	}
	
	return true, nil
}

func (ir *unitRepository) InvalidIDs(ids []uint) (invalidIDs []uint, err error) {
	for _, id := range ids {
		err := ir.db.First(&models.Unit{}, id).Error
		if err == nil {
			continue
		} 
		if err == gorm.ErrRecordNotFound {
			invalidIDs = append(invalidIDs, id)
		} else {
			return nil, err
		}
	}
	
	return invalidIDs, nil
}
