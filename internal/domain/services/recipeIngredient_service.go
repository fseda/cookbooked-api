package services

import "github.com/fseda/cookbooked-api/internal/domain/models"

type RecipeIngredientService interface {
	AddIngredientToRecipe(recipeID, ingredientID, unitID uint, quantity float32) error
	RemoveIngredientFromRecipe(recipeID, ingredientID uint) error
	GetIngredientsByRecipeID(recipeID uint) ([]*models.Ingredient, error)
	GetRecipesByIngredientID(ingredientID uint) ([]*models.Recipe, error)
}
