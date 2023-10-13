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
	GetUserRecipesTitleBySubstring(c *fiber.Ctx) error
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
		case errors.Is(err, globalerrors.RecipeInvalidIngredient):
		case errors.Is(err, globalerrors.RecipeInvalidUnit):
			return httpstatus.NotFoundError(err.Error())
		default:
			return httpstatus.BadRequestError(err.Error())
		}
	}

	return c.Status(fiber.StatusCreated).JSON(newRecipe)
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
		return httpstatus.InternalServerError(globalerrors.GlobalInternalServerError.Error())
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

// TODO: Implement this
func (rc *recipeController) GetUserRecipesTitleBySubstring(c *fiber.Ctx) error { return nil }
