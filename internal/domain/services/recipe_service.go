package services

import (
	"errors"
	"fmt"

	globalerrors "github.com/fseda/cookbooked-api/internal/domain/errors"
	"github.com/fseda/cookbooked-api/internal/domain/models"
	"github.com/fseda/cookbooked-api/internal/domain/repositories"
)

type RecipeService interface {
	CreateRecipe(
		title string,
		description string,
		body string,
		recipeIngredients []*models.RecipeIngredient,
		tagsIDs []uint,
		link string,
		userID uint,
	) (*models.Recipe, error)
	AddRecipeIngredient(
		userID uint,
		recipeID uint,
		ingredientID uint,
		unitID uint,
		quantity float32,
	) error
	FindRecipesByUserID(userID uint) ([]models.Recipe, error)
	FindUserRecipeByID(userID, recipeID uint) (*models.Recipe, error)
	FindUserRecipesTitleBySubstring(userID uint, titleSubstring string) ([]models.Recipe, error)
}

type recipeService struct {
	recipeRepository           repositories.RecipeRepository
	recipeIngredientRepository repositories.RecipeIngredientRepository
	ingredientRepository       repositories.IngredientRepository
	unitRepository             repositories.UnitRepository
}

func NewRecipeService(
	recipeRepository repositories.RecipeRepository,
	recipeIngredientRepository repositories.RecipeIngredientRepository,
	ingredientRepository repositories.IngredientRepository,
	unitRepository repositories.UnitRepository,
) RecipeService {
	return &recipeService{
		recipeRepository,
		recipeIngredientRepository,
		ingredientRepository,
		unitRepository,
	}
}

func (rs *recipeService) CreateRecipe(
	title string,
	description string,
	body string,
	recipeIngredients []*models.RecipeIngredient,
	tagsIDs []uint,
	link string,
	userID uint,
) (*models.Recipe, error) {
	var err error

	isRecipeTitleTakenByUser, err := rs.recipeRepository.IsRecipeTitleTakenByUser(userID, title)
	if err != nil {
		return nil, globalerrors.GlobalInternalServerError
	}
	if isRecipeTitleTakenByUser {
		return nil, globalerrors.RecipeTitleOfUserExists
	}

	ingredientsIDs, unitsIDs := rs.getIDs(recipeIngredients)

	// check if ingredients are unique
	if !rs.ingredientsAreUnique(ingredientsIDs) {
		return nil, globalerrors.RecipeDuplicateIngredient
	}

	exists, err := rs.ingredientRepository.ExistsAllIn(ingredientsIDs)
	if err != nil {
		return nil, globalerrors.GlobalInternalServerError
	}
	if !exists {
		return nil, globalerrors.RecipeInvalidIngredient
	}
	invalidIDs, err := rs.unitRepository.InvalidIDs(unitsIDs)
	if err != nil {
		return nil, globalerrors.GlobalInternalServerError
	}
	if invalidIDs != nil {
		return nil, fmt.Errorf("%w (%v)", globalerrors.RecipeInvalidUnit, invalidIDs)
	}

	// check if quantity is valid
	if !rs.quantitiesAreValid(recipeIngredients) {
		return nil, globalerrors.RecipeInvalidQuantity
	}

	// region TODO tags
	// recipeTagsModel := make([]*models.RecipeTag, len(tagsIDs))
	// for i, tagID := range tagsIDs {
	// 	// TODO: check if tag exists
	// 	// TODO: check for duplicate tags
	// 	recipeTagsModel[i] = &models.RecipeTag{
	// 		TagID: tagID,
	// 	}
	// }
	// endregion

	recipe := &models.Recipe{
		Title:             title,
		Description:       description,
		Body:              body,
		Link:              link,
		UserID:            &userID,
		RecipeIngredients: recipeIngredients,
		// RecipeTags:        nil,
	}
	err = rs.recipeRepository.Create(recipe)
	if err != nil {
		return nil, globalerrors.GlobalInternalServerError
	}
	return recipe, nil
}

func (rs *recipeService) AddRecipeIngredient(
	userID uint,
	recipeID uint,
	ingredientID uint,
	unitID uint,
	quantity float32,
) error {
	exists, err := rs.recipeRepository.Exists(recipeID)
	if err != nil {
		return globalerrors.GlobalInternalServerError
	}
	if !exists {
		return globalerrors.RecipeNotFound
	}

	// check if ingredient exists
	exists, err = rs.ingredientRepository.Exists(ingredientID)
	if err != nil {
		return globalerrors.GlobalInternalServerError
	}
	if !exists {
		return globalerrors.RecipeInvalidIngredient
	}

	// check if unit exists
	exists, err = rs.unitRepository.Exists(unitID)
	if err != nil {
		return globalerrors.GlobalInternalServerError
	}
	if !exists {
		return globalerrors.RecipeInvalidUnit
	}

	// check if quantity is valid
	if quantity <= 0 {
		return globalerrors.RecipeInvalidQuantity
	}

	recipeIngredient := &models.RecipeIngredient{
		RecipeID:     recipeID,
		IngredientID: ingredientID,
		UnitID:       unitID,
		Quantity:     quantity,
	}

	err = rs.recipeIngredientRepository.Link(recipeIngredient)
	if err != nil {
		if errors.Is(err, globalerrors.RecipeIngredientsMustBeUnique) {
			return globalerrors.RecipeIngredientsMustBeUnique
		}
		return globalerrors.GlobalInternalServerError
	}

	return err
}
func (rs *recipeService) FindRecipesByUserID(userID uint) ([]models.Recipe, error) {
	recipes, err := rs.recipeRepository.FindAllFromUser(userID)
	if err != nil {
		return nil, globalerrors.GlobalInternalServerError
	}
	return recipes, nil
}

func (rs *recipeService) FindUserRecipeByID(userID, recipeID uint) (*models.Recipe, error) {
	recipe, err := rs.recipeRepository.FindByID(recipeID)
	if err != nil {
		return nil, globalerrors.GlobalInternalServerError
	}
	if recipe == nil || *recipe.UserID != userID {
		return nil, globalerrors.RecipeNotFound
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

func (rs *recipeService) getIDs(recipeIngredients []*models.RecipeIngredient) (ingredientsIDs []uint, unitsIDs []uint) {
	ingredientsIDs = make([]uint, len(recipeIngredients))
	unitsIDs = make([]uint, len(recipeIngredients))
	for i, recipeIngredient := range recipeIngredients {
		ingredientsIDs[i] = recipeIngredient.IngredientID
		unitsIDs[i] = recipeIngredient.UnitID
	}
	return ingredientsIDs, unitsIDs
}

func (rs *recipeService) quantitiesAreValid(recipeIngredients []*models.RecipeIngredient) bool {
	for _, recipeIngredient := range recipeIngredients {
		if recipeIngredient.Quantity <= 0 {
			return false
		}
	}
	return true
}

func (rs *recipeService) ingredientsAreUnique(ingredientsIDs []uint) bool {
	uniqueIngredientsIDs := make(map[uint]bool)
	for _, ingredientID := range ingredientsIDs {
		if uniqueIngredientsIDs[ingredientID] {
			return false
		}
		uniqueIngredientsIDs[ingredientID] = true
	}
	return true
}
