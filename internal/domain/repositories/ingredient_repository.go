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
	Exists(id uint) (bool, error)
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

// ids must be unique
func (ir *ingredientRepository) ExistsAllIn(ids []uint) (bool, error) {
	var count int64
	res := ir.db.Select("id").Where("id IN ?", ids).Find(&models.Ingredient{}).Count(&count)
	if res.Error != nil {
		return false, res.Error
	}

	return count == int64(len(ids)), nil
}

func (ir *ingredientRepository) Exists(id uint) (bool, error) {
	var count int64
	res := ir.db.Table("ingredients").Select("id").Where("id = ?", id).Count(&count)
	if res.Error != nil {
		return false, res.Error
	}

	return count == 1, nil
}
