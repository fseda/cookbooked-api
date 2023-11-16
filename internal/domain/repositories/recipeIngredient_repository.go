package repositories

import (
	"errors"
	"fmt"

	globalerrors "github.com/fseda/cookbooked-api/internal/domain/errors"
	"github.com/fseda/cookbooked-api/internal/domain/models"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type RecipeIngredientRepository interface {
	Link(recipeIngredient *models.RecipeIngredient) (int64, error)
	LinkAll(recipeIngredients []*models.RecipeIngredient) (int64, error)
	UpdateAll(recipeIngredients []*models.RecipeIngredient) (int64, error)
	Unlink(recipeID, ingredientID uint) (int64, error)
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

func (r *recipeIngredientRepository) Link(recipeIngredient *models.RecipeIngredient) (int64, error) {
	res := r.db.Create(recipeIngredient)
	err := res.Error
	rowsAff := res.RowsAffected

	if err != nil {
		if pgError := err.(*pgconn.PgError); pgError != nil {
			if pgError.Code == "23505" {
				return rowsAff, globalerrors.RecipeIngredientsMustBeUnique
			}
		}
		return rowsAff, fmt.Errorf("error linking recipe ingredient: %w", err)
	}

	return rowsAff, nil
}

func (r *recipeIngredientRepository) LinkAll(recipeIngredients []*models.RecipeIngredient) (int64, error) {
	res := r.db.Create(&recipeIngredients)
	err := res.Error
	rowsAff := res.RowsAffected

	if err != nil {
		return rowsAff, fmt.Errorf("error linking recipe ingredient: %w", err)
	}

	return rowsAff, nil
}

func (r *recipeIngredientRepository) UpdateAll(recipeIngredients []*models.RecipeIngredient) (int64, error) {
	fmt.Println(recipeIngredients[0])
	res := r.db.Save(&recipeIngredients)
	err := res.Error
	rowsAff := res.RowsAffected

	if err != nil {
		return rowsAff, fmt.Errorf("error updating recipe ingredients: %w", err)
	}

	return rowsAff, nil
}

func (r *recipeIngredientRepository) Unlink(recipeID, ingredientID uint) (int64, error) {
	res := r.db.Unscoped().Where("recipe_id = ? AND ingredient_id = ?", recipeID, ingredientID).Delete(&models.RecipeIngredient{})
	err := res.Error
	rowsAff := res.RowsAffected

	if res.Error != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return rowsAff, globalerrors.RecipeIngredientNotFound
		}
		return rowsAff, fmt.Errorf("error unlinking recipe ingredient: %w", err)
	}

	return rowsAff, nil
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
