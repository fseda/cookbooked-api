package repositories

import (
	"errors"

	"github.com/fseda/cookbooked-api/internal/domain/models"
	"gorm.io/gorm"
)

type RecipeRepository interface {
	Create(recipe *models.Recipe) error
	FindAllFromUser(userID uint) ([]models.Recipe, error)
	FindByID(id uint) (*models.Recipe, error)
	FindRecipesByTitleSubstring(titleSubstring string) ([]models.Recipe, error)
	FindUserRecipesByTitleSubstring(userID uint, titleSubstring string) ([]models.Recipe, error)
	IsRecipeTitleTakenByUser(id uint, title string) (bool, error)
}

type recipeRepository struct {
	db *gorm.DB
}

func NewRecipeRepository(db *gorm.DB) RecipeRepository {
	return &recipeRepository{db}
}

func (r *recipeRepository) Create(recipe *models.Recipe) error {
	return r.db.Create(recipe).Error
}

func (r *recipeRepository) FindAllFromUser(userID uint) ([]models.Recipe, error) {
	var recipes []models.Recipe
	err := r.db.
		Preload("RecipeTags").
		Where("user_id = ?", userID).
		Find(&recipes).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return recipes, nil
}

func (r *recipeRepository) FindByID(id uint) (*models.Recipe, error) {
	var recipe models.Recipe
	if err := r.db.
		Preload("RecipeTags").
		Preload("RecipeIngredients").
		Preload("RecipeIngredients.Ingredient").
		Preload("RecipeIngredients.Ingredient.Category").
		Preload("RecipeIngredients.Unit").
		First(&recipe, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &recipe, nil
}

func (r *recipeRepository) FindRecipesByTitleSubstring(title string) ([]models.Recipe, error) {
	var recipes []models.Recipe

	// Find by substring of the title
	if err := r.db.Where("title ILIKE ?", "%"+title+"%").Find(&recipes).Error; err != nil {
		return nil, err
	}
	return recipes, nil
}

func (r *recipeRepository) FindUserRecipesByTitleSubstring(userID uint, title string) ([]models.Recipe, error) {
	var recipes []models.Recipe

	// Find by substring of the title
	if err := r.db.
		Where("title ILIKE ? AND user_id = ?", "%"+title+"%", userID).
		Find(&recipes).Error; err != nil {
		return nil, err
	}
	return recipes, nil
}

func (r *recipeRepository) IsRecipeTitleTakenByUser(userID uint, title string) (bool, error) {
	var recipe models.Recipe
	if err := r.db.Where("title ILIKE ?", title).Where("user_id = ?", userID).First(&recipe).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
