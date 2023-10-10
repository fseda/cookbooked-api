package services

import (
	globalerrors "github.com/fseda/cookbooked-api/internal/domain/errors"
	"github.com/fseda/cookbooked-api/internal/domain/models"
	"github.com/fseda/cookbooked-api/internal/domain/repositories"
)

type RecipeService interface {
	Create(
		title string,
		description string,
		body string,
		recipeIngredients []*RecipeIngredientInput,
		tagsIDs []uint,
		link string,
		userID uint,
	) (*models.Recipe, error)
	FindManyByUserID(userID uint) ([]models.Recipe, error)
	FindUserRecipesTitleBySubstring(userID uint, titleSubstring string) ([]models.Recipe, error)
}

type recipeService struct {
	repository repositories.RecipeRepository
}

func NewRecipeService(repository repositories.RecipeRepository) RecipeService {
	return &recipeService{repository}
}

type RecipeIngredientInput struct {
	IngredientID uint
	UnitID       uint
	Quantity     float32
}

func (rs *recipeService) Create(
	title string,
	description string,
	body string,
	recipeIngredients []*RecipeIngredientInput, 
	tagsIDs []uint,
	link string,
	userID uint,
) (*models.Recipe, error) {
	var err error

	isRecipeTitleTakenByUser, err := rs.repository.IsRecipeTitleTakenByUser(userID, title)
	if err != nil {
		return nil, globalerrors.GlobalInternalServerError
	}
	if isRecipeTitleTakenByUser {
		return nil, globalerrors.RecipeTitleOfUserExists
	}

	recipeIngredientsModel := make([]*models.RecipeIngredient, len(recipeIngredients))
	for i, recipeIngredient := range recipeIngredients {
		// TODO: check if ingredient exists
		// TODO: check if unit exists
		// TODO: check if quantity is valid
		// TODO: check for duplicate ingredients

		recipeIngredientsModel[i] = &models.RecipeIngredient{
			IngredientID: recipeIngredient.IngredientID,
			UnitID:       recipeIngredient.UnitID,
			Quantity:     recipeIngredient.Quantity,
		}
	}

	recipeTagsModel := make([]*models.RecipeTag, len(tagsIDs))
	for i, tagID := range tagsIDs {
		// TODO: check if tag exists
		// TODO: check for duplicate tags

		recipeTagsModel[i] = &models.RecipeTag{
			TagID: tagID,
		}
	}
	
	recipe := &models.Recipe{
		Title:       title,
		Description: description,
		Body:        body,
		Link:        link,
		UserID:      &userID,
		RecipeIngredients: recipeIngredientsModel,
		RecipeTags: recipeTagsModel,
	}
	err = rs.repository.Create(recipe)
	if err != nil {
		return nil, globalerrors.GlobalInternalServerError
	}
	return recipe, nil
}

func (rs *recipeService) FindManyByUserID(userID uint) ([]models.Recipe, error) {
	recipes, err := rs.repository.FindAllUserRecipes(userID)
	if err != nil {
		return nil, globalerrors.GlobalInternalServerError
	}

	return recipes, nil
}

func (rs *recipeService) FindUserRecipesTitleBySubstring(userID uint, titleSubstring string) ([]models.Recipe, error) {
	recipes, err := rs.repository.FindUserRecipesByTitleSubstring(userID, titleSubstring)
	if err != nil {
		return nil, globalerrors.GlobalInternalServerError
	}

	return recipes, nil
}
