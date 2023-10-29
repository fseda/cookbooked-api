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
	GetRecipesByUserID(c *fiber.Ctx) error
	GetRecipeDetails(c *fiber.Ctx) error
	AddRecipeIngredient(c *fiber.Ctx) error
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

// CreateRecipe godoc
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

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"rows_affected": rowsAff,
	})
}

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

func (rc *recipeController) GetRecipesByUserID(c *fiber.Ctx) error {
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

func (rc *recipeController) GetRecipeDetails(c *fiber.Ctx) error {
	claims := c.Locals("user").(*jwtutil.CustomClaims)
	userID := claims.UserID

	recipeID, _ := c.ParamsInt("id")

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
