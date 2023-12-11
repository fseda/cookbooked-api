package services

import (
	"errors"
	"fmt"

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
		recipeIngredients []*models.RecipeIngredient,
		tagsIDs []uint,
		link string,
		userID uint,
	) (*models.Recipe, validationPkg.Validation, error)
	AddRecipeIngredient(
		userID uint,
		recipeID uint,
		ingredientID uint,
		unitID uint,
		quantity float32,
	) (int64, error)
	AddRecipeIngredients(
		userID uint,
		recipeID uint,
		recipeIngredients []*models.RecipeIngredient,
	) (int64, error)
	SetRecipeIngredients(
		userID uint,
		recipeID uint,
		recipeIngredients []*models.RecipeIngredient,
	) (rowsAff int64, err error)
	RemoveRecipeIngredient(userID, recipeID, ingredientID uint) (int64, error)
	FindRecipesByUserID(userID uint) ([]models.Recipe, error)
	FindRecipeByID(recipeID uint) (*models.Recipe, error)
	FindUserRecipesTitleBySubstring(userID uint, titleSubstring string) ([]models.Recipe, error)
	UpdateRecipe(recipeID, userID uint, title, description, body, link string) (recipe *models.Recipe, validation validationPkg.Validation, err error)
	DeleteRecipe(recipeID, userID uint) (int64, error)
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
) (recipe *models.Recipe, validation validationPkg.Validation, err error) {
	validation = validationPkg.NewValidation()

	validation = rs.validateRecipe(models.Recipe{
		Base: models.Base{
			ID: 0,
		},
		Title:       title,
		Description: description,
		Body:        body,
		UserID:      &userID,
	})

	ingredientsIDs, unitsIDs := rs.getIDs(recipeIngredients)

	// check if ingredients are unique
	if !rs.ingredientsAreUnique(ingredientsIDs) {
		validation.AddError("ingredients", errors.New("ingredients must be unique"))
	}

	exists, err := rs.ingredientRepository.ExistsAllIn(ingredientsIDs)
	if err != nil {
		return
	}
	if !exists {
		validation.AddError("ingredients", errors.New("invalid ingredients"))
	}
	invalidIDs, err := rs.unitRepository.InvalidIDs(unitsIDs)
	if err != nil {
		return 
	}
	if invalidIDs != nil {
		validation.AddError("units", fmt.Errorf("invalid units (%v)", invalidIDs))
	}

	// check if quantity is valid
	if !rs.quantitiesAreValid(recipeIngredients) {
		validation.AddError("quantity", errors.New("invalid quantity"))
	}

	if validation.HasErrors() {
		return
	}

	recipe = &models.Recipe{
		Title:             title,
		Description:       description,
		Body:              body,
		Link:              link,
		UserID:            &userID,
		RecipeIngredients: recipeIngredients,
	}
	err = rs.recipeRepository.Create(recipe)
	
	return recipe, validation, err
}

func (rs *recipeService) AddRecipeIngredient(
	userID uint,
	recipeID uint,
	ingredientID uint,
	unitID uint,
	quantity float32,
) (rowsAff int64, err error) {
	exists, err := rs.recipeRepository.UserRecipeExists(userID, recipeID)
	if err != nil {
		err = globalerrors.GlobalInternalServerError
		return
	}
	if !exists {
		err = globalerrors.RecipeNotFound
		return
	}

	// check if ingredient exists
	exists, err = rs.ingredientRepository.Exists(ingredientID)
	if err != nil {
		err = globalerrors.GlobalInternalServerError
		return
	}
	if !exists {
		err = globalerrors.RecipeInvalidIngredient
		return
	}

	// check if unit exists
	exists, err = rs.unitRepository.Exists(unitID)
	if err != nil {
		err = globalerrors.GlobalInternalServerError
		return
	}
	if !exists {
		err = globalerrors.RecipeInvalidUnit
		return
	}

	// check if quantity is valid
	if quantity <= 0 {
		err = globalerrors.RecipeInvalidQuantity
		return
	}

	recipeIngredient := &models.RecipeIngredient{
		RecipeID:     recipeID,
		IngredientID: ingredientID,
		UnitID:       unitID,
		Quantity:     quantity,
	}

	rowsAff, err = rs.recipeIngredientRepository.Link(recipeIngredient)
	if err != nil {
		if errors.Is(err, globalerrors.RecipeIngredientsMustBeUnique) {
			err = globalerrors.RecipeIngredientsMustBeUnique
			return
		}
		err = globalerrors.GlobalInternalServerError
		return
	}

	return rowsAff, nil
}

func (rs *recipeService) AddRecipeIngredients(
	userID uint,
	recipeID uint,
	recipeIngredients []*models.RecipeIngredient,
) (rowsAff int64, err error) {
	exists, err := rs.recipeRepository.UserRecipeExists(userID, recipeID)
	if err != nil {
		err = globalerrors.GlobalInternalServerError
		return
	}
	if !exists {
		err = globalerrors.RecipeNotFound
		return
	}

	ingredientsIDs, unitsIDs := rs.getIDs(recipeIngredients)

	if !rs.ingredientsAreUnique(ingredientsIDs) {
		return 0, globalerrors.RecipeIngredientsMustBeUnique
	}

	exists, err = rs.ingredientRepository.ExistsAllIn(ingredientsIDs)
	if err != nil {
		err = globalerrors.GlobalInternalServerError
		return
	}
	if !exists {
		err = globalerrors.RecipeInvalidIngredient
		return
	}
	invalidIDs, err := rs.unitRepository.InvalidIDs(unitsIDs)
	if err != nil {
		err = globalerrors.GlobalInternalServerError
		return
	}
	if invalidIDs != nil {
		err = fmt.Errorf("%w (%v)", globalerrors.RecipeInvalidUnit, invalidIDs)
		return
	}

	if !rs.quantitiesAreValid(recipeIngredients) {
		err = globalerrors.RecipeInvalidQuantity
		return
	}

	recipe, err := rs.recipeRepository.FindByID(recipeID)
	if err != nil {
		err = globalerrors.GlobalInternalServerError
		return
	}

	for _, ingredient := range recipe.RecipeIngredients {
		for ri, recipeIngredient := range recipeIngredients {
			if ingredient.IngredientID == recipeIngredient.IngredientID {
				recipeIngredients[ri].ID = ingredient.ID
				break
			}
		}
	}

	rowsAff, err = rs.recipeIngredientRepository.UpdateAll(recipeIngredients)
	if err != nil {
		err = globalerrors.GlobalInternalServerError
		return
	}

	return rowsAff, nil
}

func (rs *recipeService) SetRecipeIngredients(
	userID uint,
	recipeID uint,
	recipeIngredients []*models.RecipeIngredient,
) (rowsAff int64, err error) {
	recipe, err := rs.recipeRepository.FindByID(recipeID)
	if err != nil {
		err = globalerrors.GlobalInternalServerError
		return
	}
	if recipe == nil || *recipe.UserID != userID {
		err = globalerrors.RecipeNotFound
		return
	}

	ingredientsIDs, unitsIDs := rs.getIDs(recipeIngredients)
	if !rs.ingredientsAreUnique(ingredientsIDs) {
		return 0, globalerrors.RecipeIngredientsMustBeUnique
	}

	exists, err := rs.ingredientRepository.ExistsAllIn(ingredientsIDs)
	if err != nil {
		err = globalerrors.GlobalInternalServerError
		return
	}
	if !exists {
		err = globalerrors.RecipeInvalidIngredient
		return
	}
	invalidIDs, err := rs.unitRepository.InvalidIDs(unitsIDs)
	if err != nil {
		err = globalerrors.GlobalInternalServerError
		return
	}
	if invalidIDs != nil {
		err = fmt.Errorf("%w (%v)", globalerrors.RecipeInvalidUnit, invalidIDs)
		return
	}

	if !rs.quantitiesAreValid(recipeIngredients) {
		err = globalerrors.RecipeInvalidQuantity
		return
	}

	for _, ingredient := range recipe.RecipeIngredients {
		for ri, recipeIngredient := range recipeIngredients {
			if ingredient.IngredientID == recipeIngredient.IngredientID {
				recipeIngredients[ri].ID = ingredient.ID
				break
			}
			recipeIngredients[ri].RecipeID = recipeID
		}
	}

	rowsAff, err = rs.recipeIngredientRepository.Set(recipeID, recipeIngredients)
	if err != nil {
		err = globalerrors.GlobalInternalServerError
		return
	}

	return rowsAff, nil
}

func (rs *recipeService) RemoveRecipeIngredient(userID, recipeID, ingredientID uint) (rowsAff int64, err error) {
	exists, err := rs.recipeRepository.UserRecipeExists(userID, recipeID)
	if err != nil {
		err = globalerrors.GlobalInternalServerError
		return
	}
	if !exists {
		err = globalerrors.RecipeNotFound
		return
	}
	
	rowsAff, err = rs.recipeIngredientRepository.Unlink(recipeID, ingredientID)
	if err != nil {
		if errors.Is(err, globalerrors.RecipeIngredientNotFound) {
			err = globalerrors.RecipeIngredientNotFound
			return
		}
		err = globalerrors.GlobalInternalServerError
		return
	}

	return rowsAff, nil
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
	if err != nil{
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

		isRecipeTitleTakenByUser, _ := rs.recipeRepository.IsRecipeTitleTakenByUser(*recipe.UserID, recipe.Base.ID, recipe.Title)
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
