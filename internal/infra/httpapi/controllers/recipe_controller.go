package controllers

import (
	"strings"

	globalerrors "github.com/fseda/cookbooked-api/internal/domain/errors"
	"github.com/fseda/cookbooked-api/internal/domain/services"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/validation"
	"github.com/gofiber/fiber/v2"
)

type RecipeController interface {
	CreateRecipe(c *fiber.Ctx) error
	GetRecipesByUserID(c *fiber.Ctx) error
	GetRecipesByTitleSubstring(c *fiber.Ctx) error
}

type recipeController struct {
	recipeService           services.RecipeService
	// recipeIngredientService services.RecipeIngredientService
}

func NewRecipeController(recipeService services.RecipeService,
	//  recipeIngredientService services.RecipeIngredientService,
) RecipeController {
	return &recipeController{
		recipeService,
		// recipeIngredientService,
	}
}

type recipeIngredientRequest struct {
	IngredientID uint    `json:"ingredient_id" validate:"required=true,"`
	UnitID       uint    `json:"unit_id" validate:"required=true,number=true"`
	Quantity     float32 `json:"quantity" validate:"required=true"`
}

type createRecipeRequest struct {
	Title             string                            `json:"title" validate:"required=true,min=3,max=255"`
	Description       string                            `json:"description" validate:"required=true"`
	Body              string                            `json:"body" validate:"required=true"`
	Link              string                            `json:"link" validate:"url=true"`
	RecipeIngredients []*services.RecipeIngredientInput `json:"recipe_ingredients" validate:"required=true"`
}

func (rc *recipeController) CreateRecipe(c *fiber.Ctx) error {
	// userClaims := c.Locals("user").(*jwtutil.CustomClaims)
	// userID = userClaims.UserID

	var req createRecipeRequest
	if err := c.BodyParser(&req); err != nil {
		return httpstatus.UnprocessableEntityError(globalerrors.GlobalUnableToParseBody.Error())
	}

	if errMsgs := validation.MyValidator.CreateErrorResponse(req); len(errMsgs) > 0 {
		return httpstatus.BadRequestError(strings.Join(errMsgs, " and "))
	}

	return c.SendString("ok!")

}

func (rc *recipeController) GetRecipesByUserID(c *fiber.Ctx) error { return nil }

func (rc *recipeController) GetRecipesByTitleSubstring(c *fiber.Ctx) error { return nil }
