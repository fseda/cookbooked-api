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

func (ir *unitRepository) ExistsAllIn(ids []uint) (bool, error) {
	res := ir.db.Model(&models.Unit{}).Where("id IN ?", ids)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, res.Error
	}

	return true, nil	
}