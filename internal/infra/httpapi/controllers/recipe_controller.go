package controllers

import (
	"errors"

	globalerrors "github.com/fseda/cookbooked-api/internal/domain/errors"
	"github.com/fseda/cookbooked-api/internal/domain/services"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	jwtutil "github.com/fseda/cookbooked-api/internal/infra/jwt"
	"github.com/gofiber/fiber/v2"
)

type RecipeController interface {
	CreateRecipe(c *fiber.Ctx) error
	GetAllRecipesByUserID(c *fiber.Ctx) error
	GetRecipeDetails(c *fiber.Ctx) error
	UpdateRecipe(c *fiber.Ctx) error
	DeleteRecipe(c *fiber.Ctx) error
}

type recipeController struct {
	recipeService services.RecipeService
}

func NewRecipeController(recipeService services.RecipeService) RecipeController {
	return &recipeController{recipeService}
}

type createRecipeRequest struct {
	Title             string                     `json:"title" validate:"required=true,min=3,max=255"`
	Description       string                     `json:"description" validate:"required=true"`
	Body              string                     `json:"body" validate:"required=true"`
	Link              string                     `json:"link"`
	TagIDs            []uint                     `json:"tag_ids" validate:"dive,number=true"`
}

type updateRecipeRequest struct {
	Title       string `json:"title" validate:"required=true,min=3,max=255"`
	Description string `json:"description" validate:"required=true"`
	Body        string `json:"body" validate:"required=true"`
	Link        string `json:"link"`
}

type getRecipeDetailsResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
	Link        string `json:"link"`
	// RecipeTags        []*models.RecipeTag        `json:"recipe_tags"`
}

type getRecipeResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	// Tags        []*models.RecipeTag `json:"tags"`
}

type getAllRecipesResponse struct {
	Recipes []*getRecipeResponse `json:"recipes"`
}

// @Summary		Create a new recipe
// @Description	Create a new recipe with the given input data
// @Tags			Recipes
// @Accept			json
// @Produce		json
// @Param			input	body	createRecipeRequest	true	"Recipe creation data"
// @Security		ApiKeyAuth
// @Success		201	{object}	createRecipeRequest
// @Failure		400	{object}	httpstatus.GlobalErrorHandlerResp
// @Router			/recipes/new   [post]
func (rc *recipeController) CreateRecipe(c *fiber.Ctx) error {
	userClaims := c.Locals("user").(*jwtutil.CustomClaims)
	userID := userClaims.UserID

	var req createRecipeRequest
	if err := c.BodyParser(&req); err != nil {
		return httpstatus.UnprocessableEntityError(globalerrors.GlobalUnableToParseBody.Error())
	}

	newRecipe, validation, err := rc.recipeService.CreateRecipe(
		req.Title,
		req.Description,
		req.Body,
		req.TagIDs,
		req.Link,
		userID,
	)
	if err != nil {
		return httpstatus.InternalServerError(err.Error())
	}
	if validation.HasErrors() {
		return c.Status(fiber.StatusBadRequest).JSON(validation)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"recipe": newRecipe,
	})
}

// @Summary		Get all recipes from a user
// @Description	Get all recipes from a user, by user id
// @Tags			Recipes
// @Produces		json
// @Security		ApiKeyAuth
// @Success		200	{object}	getAllRecipesResponse
// @Failure		400
// @Router			/recipes [get]
func (rc *recipeController) GetAllRecipesByUserID(c *fiber.Ctx) error {
	userClaims := c.Locals("user").(*jwtutil.CustomClaims)
	userID := userClaims.UserID

	recipes, err := rc.recipeService.FindRecipesByUserID(userID)
	if err != nil {
		return httpstatus.InternalServerError(err.Error())
	}

	recipesResponse := make([]*getRecipeResponse, len(recipes))
	for i, recipe := range recipes {
		recipesResponse[i] = &getRecipeResponse{
			ID:          recipe.ID,
			Title:       recipe.Title,
			Description: recipe.Description,
			// Tags:        recipe.RecipeTags,
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"recipes": recipesResponse,
	})
}

// @Summary		Get a recipe details
// @Description	Get a recipe details, by recipe id
// @Tags			Recipes
// @Produces		json
// @Param			recipe_id	path	integer	true	"Recipe ID"
// @Security		ApiKeyAuth
// @Success		200	{object}	getRecipeDetailsResponse
// @Failure		400
// @Router			/recipes/{recipe_id} [get]
func (rc *recipeController) GetRecipeDetails(c *fiber.Ctx) error {
	claims := c.Locals("user").(*jwtutil.CustomClaims)
	userID := claims.UserID

	recipeID, _ := c.ParamsInt("recipe_id")

	recipe, err := rc.recipeService.FindRecipeByID(uint(recipeID))
	if err != nil {
		return httpstatus.InternalServerError(err.Error())
	}
	if recipe == nil {
		return httpstatus.NotFoundError(globalerrors.RecipeNotFound.Error())
	}

	canEdit := userID == *recipe.UserID
	recipeDetailsResponse := &getRecipeDetailsResponse{
		ID:                recipe.ID,
		Title:             recipe.Title,
		Description:       recipe.Description,
		Body:              recipe.Body,
		Link:              recipe.Link,
		// RecipeTags:        recipe.RecipeTags,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"recipe":   recipeDetailsResponse,
		"can_edit": canEdit,
	})
}

// @Summary		Update recipe details
// @Description	Update recipe details, by recipe id
// @Tags			Recipes
// @Accept			json
// @Produce		json
// @Param			recipe_id	path	integer				true	"Recipe ID"
// @Param			input		body	updateRecipeRequest	true	"Recipe update data"
// @Security		ApiKeyAuth
// @Success		200	{object}	updateRecipeRequest
// @Failure		400	{object}	httpstatus.GlobalErrorHandlerResp
// @Router			/recipes/{recipe_id} [patch]
func (rc *recipeController) UpdateRecipe(c *fiber.Ctx) error {
	claims := c.Locals("user").(*jwtutil.CustomClaims)
	userID := claims.UserID

	recipeID, _ := c.ParamsInt("recipe_id")

	var req updateRecipeRequest
	if err := c.BodyParser(&req); err != nil {
		return httpstatus.UnprocessableEntityError(err.Error())
	}

	updatedRecipe, validation, err := rc.recipeService.UpdateRecipe(
		uint(recipeID),
		userID,
		req.Title,
		req.Description,
		req.Body,
		req.Link,
	)
	if err != nil {
		switch {
		case errors.Is(err, globalerrors.RecipeNotFound):
			return httpstatus.NotFoundError(err.Error())
		default:
			return httpstatus.InternalServerError(err.Error())
		}
	}
	if validation.HasErrors() {
		return c.Status(fiber.StatusBadRequest).JSON(validation)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"recipe": updatedRecipe,
	})
}

// @Summary		Delete a recipe
// @Description	Delete a recipe, by recipe id
// @Tags			Recipes
// @Param			recipe_id	path	integer	true	"Recipe ID"
// @Security		ApiKeyAuth
// @Success		200	{string}	rows_affected
// @Failure		400	{object}	httpstatus.GlobalErrorHandlerResp
// @Failure		404	{object}	httpstatus.GlobalErrorHandlerResp
// @Router			/recipes/{recipe_id} [delete]
func (rc *recipeController) DeleteRecipe(c *fiber.Ctx) error {
	claims := c.Locals("user").(*jwtutil.CustomClaims)
	userID := claims.UserID

	recipeID, _ := c.ParamsInt("recipe_id")

	rowsAff, err := rc.recipeService.DeleteRecipe(uint(recipeID), userID)
	if err != nil {
		switch err {
		case globalerrors.GlobalInternalServerError:
			return httpstatus.InternalServerError(globalerrors.GlobalInternalServerError.Error())
		case globalerrors.RecipeNotFound:
			return httpstatus.NotFoundError(globalerrors.RecipeNotFound.Error())
		default:
			return httpstatus.BadRequestError(err.Error())
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"rows_affected": rowsAff,
	})
}
