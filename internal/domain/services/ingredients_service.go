package services

import (
	"github.com/fseda/cookbooked-api/internal/domain/models"
	"github.com/fseda/cookbooked-api/internal/domain/repositories"
)

type IngredientService interface {
	GetAllIngredients() ([]models.Ingredient, error)
}

type ingredientService struct {
	ingredientRepository repositories.IngredientRepository
}

func NewIngredientService(ingredientRepository repositories.IngredientRepository) IngredientService {
	return &ingredientService{ingredientRepository}
}

func (s *ingredientService) GetAllIngredients() ([]models.Ingredient, error) {
	return s.ingredientRepository.GetAllIngredients()
}