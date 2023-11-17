package controllers

import (
	"github.com/fseda/cookbooked-api/internal/domain/models"
	"github.com/fseda/cookbooked-api/internal/domain/services"
	"github.com/gofiber/fiber/v2"
)

type IngredientController interface {
	GetAllIngredients(ctx *fiber.Ctx) error
}

type ingredientController struct {
	ingredientService services.IngredientService
}

func NewIngredientController(ingredientService services.IngredientService) IngredientController {
	return &ingredientController{ingredientService}
}

type ingredientResponse struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	CategoryID *uint `json:"category_id"`
	Category models.IngredientsCategory `json:"category"`
}

type ingredientsResponse struct {
	Ingredients []ingredientResponse `json:"ingredients"`
}

//	@Summary		Get all ingredients
//	@Description	Get all ingredients
//	@Tags			Ingredients
//	@Success		200	{object}	ingredientsResponse.Ingredients
//	@Router			/ingredients [get]
func (c *ingredientController) GetAllIngredients(ctx *fiber.Ctx) error {
	ingredients, err := c.ingredientService.GetAllIngredients()
	if err != nil {
		return err
	}

	ingredientsResponse := &ingredientsResponse{}

	for _, ingredient := range ingredients {
		ingredientsResponse.Ingredients = append(ingredientsResponse.Ingredients, ingredientResponse{
			ID: ingredient.ID,
			Name: ingredient.Name,
			Icon: ingredient.Icon,
			CategoryID: ingredient.CategoryID,
			Category: ingredient.Category,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(ingredientsResponse.Ingredients)
}