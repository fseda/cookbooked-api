package controllers

import (
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

// @Summary Get all ingredients
// @Description Get all ingredients
// @Tags Ingredients
// @Success 200 {object} []models.Ingredient
// @Router /ingredients [get]
func (c *ingredientController) GetAllIngredients(ctx *fiber.Ctx) error {
	ingredients, err := c.ingredientService.GetAllIngredients()
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"ingredients": ingredients})
}