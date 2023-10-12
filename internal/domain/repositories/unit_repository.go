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
	var count int64
	res := ir.db.Where("id IN ?", ids).Find(&models.Unit{}).Count(&count)
	if res.Error != nil {
		return false, res.Error
	}

	return count == int64(len(ids)), nil	
}
