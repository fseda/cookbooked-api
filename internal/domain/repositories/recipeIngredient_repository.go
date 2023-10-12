package repositories

import (
	"github.com/fseda/cookbooked-api/internal/domain/models"
	"gorm.io/gorm"
)

type RecipeIngredientRepository interface {
	Link(recipeID, ingredientID, unitID uint, quantity float32) error
	Unlink(recipeID, ingredientID uint) error
	GetIngredientsByRecipeID(recipeID uint) ([]*models.Ingredient, error)
	GetRecipesByIngredientID(ingredientID uint) ([]*models.Recipe, error)
	GetUserRecipesByIngredientID(userID, ingredientID uint) ([]*models.Recipe, error)
}

type recipeIngredientRepository struct {
	db *gorm.DB
}

func NewRecipeIngredientRepository(db *gorm.DB) RecipeIngredientRepository {
	return &recipeIngredientRepository{db}
}

func (r *recipeIngredientRepository) Link(recipeID, ingredientID, unitID uint, quantity float32) error {
	recipeIngredient := &models.RecipeIngredient{
		RecipeID:     recipeID,
		IngredientID: ingredientID,
		UnitID:       unitID,
		Quantity:     quantity,
	}
	return r.db.Create(recipeIngredient).Error
}

func (r *recipeIngredientRepository) Unlink(recipeID, ingredientID uint) error {
	return r.db.Where("recipe_id = ? AND ingredient_id = ?", recipeID, ingredientID).Delete(&models.RecipeIngredient{}).Error
}

func (r *recipeIngredientRepository) GetIngredientsByRecipeID(recipeID uint) ([]*models.Ingredient, error) {
	var ingredients []*models.Ingredient
	err := r.db.
		Table("ingredients as i").
		Select("i.id, i.name, i.description, ic.category").
		Joins("JOIN recipe_ingredients as ri ON ri.ingredient_id = i.id").
		Joins("JOIN ingredients_categories as ic ON ic.id = i.category_id").
		Where("ri.recipe_id = ?", recipeID).
		Find(&ingredients).Error
	if err != nil {
		return nil, err
	}
	return ingredients, err
}

func (r *recipeIngredientRepository) GetRecipesByIngredientID(ingredientID uint) ([]*models.Recipe, error) {
	// join the tables using preload
	
	var recipes []*models.Recipe
	err := r.db.
		Preload("RecipeIngredients").
		Preload("RecipeIngredients.Ingredient").
		Preload("RecipeIngredients.Unit").
		Preload("RecipeTags").
		Preload("RecipeTags.Tag").
		Where("recipe_ingredients.ingredient_id = ?", ingredientID).
		Find(&recipes).Error
	if err != nil {
		return nil, err
	}
	return recipes, err
}

func (r *recipeIngredientRepository) GetUserRecipesByIngredientID(userID, ingredientID uint) ([]*models.Recipe, error) {
	// join the tables using preload
	var recipes []*models.Recipe
	err := r.db.
		Preload("RecipeIngredients").
		Preload("RecipeIngredients.Ingredient").
		Preload("RecipeIngredients.Unit").
		Preload("RecipeTags").
		Preload("RecipeTags.Tag").
		Where("recipe_ingredients.ingredient_id = ? AND recipes.user_id = ?", ingredientID, userID).
		Find(&recipes).Error
	if err != nil {
		return nil, err
	}
	return recipes, err
}