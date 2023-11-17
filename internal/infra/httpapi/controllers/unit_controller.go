package controllers

import (
	"github.com/fseda/cookbooked-api/internal/domain/models"
	"github.com/fseda/cookbooked-api/internal/domain/services"
	"github.com/gofiber/fiber/v2"
)

type UnitController interface {
	GetAllUnits(ctx *fiber.Ctx) error
}

type unitController struct {
	unitService services.UnitService
}

func NewUnitController(unitService services.UnitService) UnitController {
	return &unitController{unitService}
}

type unitResponse struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	Type models.Type `json:"type"`
}

type unitsResponse struct {
	Units []unitResponse `json:"units"`
}

//	@Summary		Get all units
//	@Description	Get all units
//	@Tags			Units
//	@Success		200	{object}	unitsResponse
//	@Router			/units [get]
func (c *unitController) GetAllUnits(ctx *fiber.Ctx) error {
	units, err := c.unitService.GetAllUnits()
	if err != nil {
		return err
	}

	unitsResponse := &unitsResponse{}

	for _, unit := range units {
		unitsResponse.Units = append(unitsResponse.Units, unitResponse{
			ID: unit.ID,
			Name: unit.Name,
			Symbol: unit.Symbol,
			Type: unit.Type,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(unitsResponse.Units)
}