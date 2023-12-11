package services

import (
	"errors"

	globalerrors "github.com/fseda/cookbooked-api/internal/domain/errors"
	"github.com/fseda/cookbooked-api/internal/domain/models"
	"github.com/fseda/cookbooked-api/internal/domain/repositories"
	validationPkg "github.com/fseda/cookbooked-api/internal/infra/httpapi/validation"
)

type RecipeService interface {
	CreateRecipe(
		title string,
		description string,
		body string,
		tagsIDs []uint,
		link string,
		userID uint,
	) (*models.Recipe, validationPkg.Validation, error)
	FindRecipesByUserID(userID uint) ([]models.Recipe, error)
	FindRecipeByID(recipeID uint) (*models.Recipe, error)
	FindUserRecipesTitleBySubstring(userID uint, titleSubstring string) ([]models.Recipe, error)
	UpdateRecipe(recipeID, userID uint, title, description, body, link string) (recipe *models.Recipe, validation validationPkg.Validation, err error)
	DeleteRecipe(recipeID, userID uint) (int64, error)
}

type recipeService struct {
	recipeRepository repositories.RecipeRepository
}

func NewRecipeService(recipeRepository repositories.RecipeRepository) RecipeService {
	return &recipeService{recipeRepository}
}

func (rs *recipeService) CreateRecipe(
	title string,
	description string,
	body string,
	tagsIDs []uint,
	link string,
	userID uint,
) (recipe *models.Recipe, validation validationPkg.Validation, err error) {
	validation = rs.validateRecipe(models.Recipe{
		Base: models.Base{
			ID: 0,
		},
		Title:       title,
		Description: description,
		Body:        body,
		UserID:      &userID,
	})

	if validation.HasErrors() {
		return
	}

	recipe = &models.Recipe{
		Title:       title,
		Description: description,
		Body:        body,
		Link:        link,
		UserID:      &userID,
	}
	err = rs.recipeRepository.Create(recipe)

	return recipe, validation, err
}

func (rs *recipeService) FindRecipesByUserID(userID uint) ([]models.Recipe, error) {
	recipes, err := rs.recipeRepository.FindAllFromUser(userID)
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

func (rs *recipeService) FindRecipeByID(recipeID uint) (*models.Recipe, error) {
	recipe, err := rs.recipeRepository.FindByID(recipeID)
	if err != nil {
		return nil, err
	}
	return recipe, nil
}

func (rs *recipeService) FindUserRecipesTitleBySubstring(userID uint, titleSubstring string) ([]models.Recipe, error) {
	recipes, err := rs.recipeRepository.FindUserRecipesByTitleSubstring(userID, titleSubstring)
	if err != nil {
		return nil, globalerrors.GlobalInternalServerError
	}

	return recipes, nil
}

func (rs *recipeService) UpdateRecipe(recipeID, userID uint, title, description, body, link string) (recipe *models.Recipe, validation validationPkg.Validation, err error) {
	recipe = &models.Recipe{
		Base: models.Base{
			ID: recipeID,
		},
		UserID:      &userID,
		Title:       title,
		Description: description,
		Body:        body,
		Link:        link,
	}

	exists, err := rs.recipeRepository.UserRecipeExists(userID, recipeID)
	if err != nil {
		return nil, validation, err
	}
	if !exists {
		return nil, validation, globalerrors.RecipeNotFound
	}

	validation = rs.validateRecipe(*recipe)
	if validation.HasErrors() {
		return nil, validation, nil
	}

	err = rs.recipeRepository.Update(recipe)
	if err != nil {
		return nil, validation, err
	}

	updatedRecipe, err := rs.recipeRepository.FindByID(recipeID)

	return updatedRecipe, validation, err
}

func (rs *recipeService) DeleteRecipe(recipeID, userID uint) (rowsAff int64, err error) {
	rowsAff, err = rs.recipeRepository.Delete(recipeID, userID)
	if err != nil {
		return 0, globalerrors.GlobalInternalServerError
	}

	return rowsAff, nil
}

func (rs *recipeService) validateRecipe(recipe models.Recipe) validationPkg.Validation {
	validation := validationPkg.NewValidation()

	if recipe.Title == "" {
		validation.AddError("title", errors.New("title is required"))
	} else {
		if len(recipe.Title) < 3 {
			validation.AddError("title", errors.New("title must be longer than 2 characters"))
		}
		if len(recipe.Title) >= 255 {
			validation.AddError("title", errors.New("title must be shorter than 255 characters"))
		}

		isRecipeTitleTakenByUser, _ := rs.recipeRepository.IsRecipeTitleTakenByUser(*recipe.UserID, recipe.ID, recipe.Title)
		if isRecipeTitleTakenByUser {
			validation.AddError("title", errors.New("title already taken"))
		}
	}

	if recipe.Description == "" {
		validation.AddError("description", errors.New("description is required"))
	}

	if recipe.Body == "" {
		validation.AddError("body", errors.New("body is required"))
	}

	return validation
}
