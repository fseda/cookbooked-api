package repositories

import (
	"github.com/fseda/cookbooked-api/internal/domain/models"
	"gorm.io/gorm"
)

func NewIngredientRepository(db *gorm.DB) IngredientRepository {
	return &ingredientRepository{db}
}

type IngredientRepository interface {
	FindByID(id uint) (*models.Ingredient, error)
	ExistsAllIn(ids []uint) (bool, error)
}

type ingredientRepository struct {
	db *gorm.DB
}

func (ir *ingredientRepository) FindByID(id uint) (*models.Ingredient, error) {
	var ingredient models.Ingredient
	result := ir.db.First(&ingredient, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &ingredient, nil
}

func (ir *ingredientRepository) ExistsAllIn(ids []uint) (bool, error) {
	res := ir.db.Model(&models.Ingredient{}).Where("id IN ?", ids)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, res.Error
	}

	return true, nil	
}

