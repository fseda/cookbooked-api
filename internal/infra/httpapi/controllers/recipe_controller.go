package controllers

import (
	"errors"
	"strings"

	globalerrors "github.com/fseda/cookbooked-api/internal/domain/errors"
	"github.com/fseda/cookbooked-api/internal/domain/models"
	"github.com/fseda/cookbooked-api/internal/domain/services"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/validation"
	jwtutil "github.com/fseda/cookbooked-api/internal/infra/jwt"
	"github.com/gofiber/fiber/v2"
)

type RecipeController interface {
	CreateRecipe(c *fiber.Ctx) error
	GetAllRecipesByUserID(c *fiber.Ctx) error
	GetRecipeDetails(c *fiber.Ctx) error
	AddRecipeIngredient(c *fiber.Ctx) error
	AddRecipeIngredients(c *fiber.Ctx) error
	RemoveRecipeIngredient(c *fiber.Ctx) error
}

type recipeController struct {
	recipeService services.RecipeService
}

func NewRecipeController(recipeService services.RecipeService) RecipeController {
	return &recipeController{recipeService}
}

type recipeIngredientRequest struct {
	IngredientID uint    `json:"ingredient_id" validate:"required=true,number=true"`
	UnitID       uint    `json:"unit_id" validate:"required=true,number=true"`
	Quantity     float32 `json:"quantity" validate:"required=true"`
}

type recipeIngredientsRequest struct {
	RecipeIngredients []*recipeIngredientRequest `json:"recipe_ingredients" validate:"required=true"`
}

type createRecipeRequest struct {
	Title             string                     `json:"title" validate:"required=true,min=3,max=255"`
	Description       string                     `json:"description" validate:"required=true"`
	Body              string                     `json:"body" validate:"required=true"`
	Link              string                     `json:"link"`
	RecipeIngredients []*recipeIngredientRequest `json:"recipe_ingredients"`
	TagIDs            []uint                     `json:"tag_ids" validate:"dive,number=true"`
}

type getRecipeDetailsResponse struct {
	ID                uint                       `json:"id"`
	Title             string                     `json:"title"`
	Description       string                     `json:"description"`
	Body              string                     `json:"body"`
	Link              string                     `json:"link"`
	RecipeIngredients []*models.RecipeIngredient `json:"recipe_ingredients"`
	RecipeTags        []*models.RecipeTag        `json:"recipe_tags"`
}

type getRecipeResponse struct {
	ID          uint                `json:"id"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Tags        []*models.RecipeTag `json:"tags"`
}

type getAllRecipesResponse struct {
	Recipes []*getRecipeResponse `json:"recipes"`
}

//	@Summary		Create a new recipe
//	@Description	Create a new recipe with the given input data
//	@Tags			Recipes
//	@Accept			json
//	@Produce		json
//	@Param			input	body	createRecipeRequest	true	"Recipe creation data"
//	@Security		ApiKeyAuth
//	@Success		201	{object}	createRecipeRequest
//	@Failure		400	{object}	httpstatus.GlobalErrorHandlerResp
//	@Router			/recipes/new   [post]
func (rc *recipeController) CreateRecipe(c *fiber.Ctx) error {
	userClaims := c.Locals("user").(*jwtutil.CustomClaims)
	userID := userClaims.UserID

	var req createRecipeRequest
	if err := c.BodyParser(&req); err != nil {
		return httpstatus.UnprocessableEntityError(globalerrors.GlobalUnableToParseBody.Error())
	}

	recipeIngredientInput := make([]*models.RecipeIngredient, len(req.RecipeIngredients))

	errMsgs := validation.MyValidator.CreateErrorResponse(req)
	for i, ri := range req.RecipeIngredients {
		riErrMsgs := validation.MyValidator.CreateErrorResponse(ri)
		if len(riErrMsgs) > 0 {
			errMsgs = append(errMsgs, riErrMsgs...)
		}

		recipeIngredientInput[i] = &models.RecipeIngredient{
			IngredientID: ri.IngredientID,
			UnitID:       ri.UnitID,
			Quantity:     ri.Quantity,
		}
	}
	if len(errMsgs) > 0 {
		return httpstatus.BadRequestError(strings.Join(errMsgs, " and "))
	}

	newRecipe, err := rc.recipeService.CreateRecipe(
		req.Title,
		req.Description,
		req.Body,
		recipeIngredientInput,
		req.TagIDs,
		req.Link,
		userID,
	)
	if err != nil {
		switch {
		case errors.Is(err, globalerrors.GlobalInternalServerError):
			return httpstatus.InternalServerError(globalerrors.GlobalInternalServerError.Error())
		case errors.Is(err, globalerrors.RecipeInvalidUnit), errors.Is(err, globalerrors.RecipeInvalidIngredient):
			return httpstatus.NotFoundError(err.Error())
		default:
			return httpstatus.BadRequestError(err.Error())
		}
	}

	return c.Status(fiber.StatusCreated).JSON(newRecipe)
}

//	@Summary		Add an ingredient to a recipe
//	@Description	Add a ingredient to a recipe, if it exists in the recipe update
//	@Tags			Recipe
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id	path	integer					true	"Recipe ID"
//	@Param			input		body	recipeIngredientRequest	true	"Recipe ingredient data"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	recipeIngredientRequest
//	@Failure		400 {object} httpstatus.GlobalErrorHandlerResp
//	@Router			/recipes/{recipe_id}/ingredients [patch]
func (rc *recipeController) AddRecipeIngredient(c *fiber.Ctx) error {
	userClaims := c.Locals("user").(*jwtutil.CustomClaims)
	userID := userClaims.UserID

	recipeID, _ := c.ParamsInt("recipe_id")

	var req recipeIngredientRequest
	if err := c.BodyParser(&req); err != nil {
		return httpstatus.UnprocessableEntityError(globalerrors.GlobalUnableToParseBody.Error())
	}

	errMsgs := validation.MyValidator.CreateErrorResponse(req)
	if len(errMsgs) > 0 {
		return httpstatus.BadRequestError(strings.Join(errMsgs, " and "))
	}

	rowsAff, err := rc.recipeService.AddRecipeIngredient(
		userID,
		uint(recipeID),
		req.IngredientID,
		req.UnitID,
		req.Quantity,
	)
	if err != nil {
		switch {
		case errors.Is(err, globalerrors.GlobalInternalServerError):
			return httpstatus.InternalServerError(globalerrors.GlobalInternalServerError.Error())

		case errors.Is(err, globalerrors.RecipeNotFound):
			return httpstatus.NotFoundError(err.Error())

		case errors.Is(err, globalerrors.RecipeInvalidIngredient), errors.Is(err, globalerrors.RecipeInvalidUnit):
			return httpstatus.NotFoundError(err.Error())

		case errors.Is(err, globalerrors.RecipeInvalidQuantity):
			return httpstatus.BadRequestError(err.Error())

		case errors.Is(err, globalerrors.RecipeIngredientsMustBeUnique):
			return httpstatus.ConflictError(err.Error())

		default:
			return httpstatus.BadRequestError(err.Error())
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"rows_affected": rowsAff,
	})
}

//	@Summary		Add multiple ingredients to a recipe
//	@Description	Add multiple ingredients to a recipe, if it exists in the recipe update
//	@Tags			Recipes
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id	path	integer						true	"Recipe ID"
//	@Param			input		body	recipeIngredientsRequest	true	"Recipe ingredients data"
//	@Security		ApiKeyAuth
//	@Success		200
//	@Failure		400
//	@Router			/recipes/{recipe_id}/ingredients [patch]
func (rs *recipeController) AddRecipeIngredients(c *fiber.Ctx) error {
	userClaims := c.Locals("user").(*jwtutil.CustomClaims)
	userID := userClaims.UserID

	recipeID, _ := c.ParamsInt("recipe_id")

	var req recipeIngredientsRequest
	if err := c.BodyParser(&req); err != nil {
		return httpstatus.UnprocessableEntityError(globalerrors.GlobalUnableToParseBody.Error())
	}

	recipeIngredientInput := make([]*models.RecipeIngredient, len(req.RecipeIngredients))

	errMsgs := validation.MyValidator.CreateErrorResponse(req)
	for i, ri := range req.RecipeIngredients {
		riErrMsgs := validation.MyValidator.CreateErrorResponse(ri)
		if len(riErrMsgs) > 0 {
			errMsgs = append(errMsgs, riErrMsgs...)
		}

		recipeIngredientInput[i] = &models.RecipeIngredient{
			RecipeID:     uint(recipeID),
			IngredientID: ri.IngredientID,
			UnitID:       ri.UnitID,
			Quantity:     ri.Quantity,
		}
	}
	if len(errMsgs) > 0 {
		return httpstatus.BadRequestError(strings.Join(errMsgs, " and "))
	}

	rowsAff, err := rs.recipeService.AddRecipeIngredients(userID, uint(recipeID), recipeIngredientInput)
	if err != nil {
		switch {
		case errors.Is(err, globalerrors.GlobalInternalServerError):
			return httpstatus.InternalServerError(globalerrors.GlobalInternalServerError.Error())

		case errors.Is(err, globalerrors.RecipeNotFound):
			return httpstatus.NotFoundError(err.Error())

		case errors.Is(err, globalerrors.RecipeInvalidIngredient), errors.Is(err, globalerrors.RecipeInvalidUnit):
			return httpstatus.NotFoundError(err.Error())

		case errors.Is(err, globalerrors.RecipeInvalidQuantity):
			return httpstatus.BadRequestError(err.Error())

		case errors.Is(err, globalerrors.RecipeIngredientsMustBeUnique):
			return httpstatus.ConflictError(err.Error())

		default:
			return httpstatus.BadRequestError(err.Error())
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"rows_affected": rowsAff,
	})
}

//	@Summary		Remove an ingredient from a recipe
//	@Description	Remove an ingredient from a recipe
//	@Tags			Recipes
//	@Param			recipe_id				path	integer	true	"Recipe ID"
//	@Param			recipe_ingredient_id	path	integer	true	"Recipe ingredient ID"
//	@Security		ApiKeyAuth
//	@Success		200
//	@Failure		400
//	@Router			/recipes/{recipe_id}/ingredients/{recipe_ingredient_id} [delete]
func (rc *recipeController) RemoveRecipeIngredient(c *fiber.Ctx) error {
	userClaims := c.Locals("user").(*jwtutil.CustomClaims)
	userID := userClaims.UserID

	recipeID, _ := c.ParamsInt("recipe_id")
	ingredientID, _ := c.ParamsInt("recipe_ingredient_id")

	rowsAff, err := rc.recipeService.RemoveRecipeIngredient(userID, uint(recipeID), uint(ingredientID))
	if err != nil {
		switch err {
		case globalerrors.GlobalInternalServerError:
			return httpstatus.InternalServerError(globalerrors.GlobalInternalServerError.Error())

		case globalerrors.RecipeNotFound:
			return httpstatus.NotFoundError(globalerrors.RecipeNotFound.Error())

		case globalerrors.RecipeIngredientNotFound:
			return httpstatus.NotFoundError(globalerrors.RecipeIngredientNotFound.Error())

		default:
			return httpstatus.BadRequestError(err.Error())
		}
	}

	if rowsAff == 0 {
		return httpstatus.NoContent("no rows affected")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"rows_affected": rowsAff,
	})
}

//	@Summary		Get all recipes from a user
//	@Description	Get all recipes from a user, by user id
//	@Tags			Recipes
//	@Produces		json
//	@Security		ApiKeyAuth
//	@Success		200
//	@Failure		400
//	@Router			/recipes [get]
func (rc *recipeController) GetAllRecipesByUserID(c *fiber.Ctx) error {
	userClaims := c.Locals("user").(*jwtutil.CustomClaims)
	userID := userClaims.UserID

	recipes, err := rc.recipeService.FindRecipesByUserID(userID)
	if err != nil {
		return httpstatus.InternalServerError(globalerrors.GlobalInternalServerError.Error())
	}

	recipesResponse := make([]*getRecipeResponse, len(recipes))
	for i, recipe := range recipes {
		recipesResponse[i] = &getRecipeResponse{
			ID:          recipe.ID,
			Title:       recipe.Title,
			Description: recipe.Description,
			Tags:        recipe.RecipeTags,
		}
	}

	return c.Status(fiber.StatusOK).JSON(recipesResponse)
}


//	@Summary		Get a recipe details
//	@Description	Get a recipe details, by recipe id
//	@Tags			Recipes
//	@Produces		json
//	@Param			recipe_id	path	integer	true	"Recipe ID"
//	@Security		ApiKeyAuth
//	@Success		200	
//	@Failure		400
//	@Router			/recipes/{recipe_id} [get]
func (rc *recipeController) GetRecipeDetails(c *fiber.Ctx) error {
	claims := c.Locals("user").(*jwtutil.CustomClaims)
	userID := claims.UserID

	recipeID, _ := c.ParamsInt("recipe_id")

	recipe, err := rc.recipeService.FindUserRecipeByID(userID, uint(recipeID))
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

	recipeDetailsResponse := &getRecipeDetailsResponse{
		ID:                recipe.ID,
		Title:             recipe.Title,
		Description:       recipe.Description,
		Body:              recipe.Body,
		Link:              recipe.Link,
		RecipeIngredients: recipe.RecipeIngredients,
		RecipeTags:        recipe.RecipeTags,
	}

	return c.Status(fiber.StatusOK).JSON(recipeDetailsResponse)
}
