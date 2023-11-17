package controllers

import (
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

// @Summary Get all units
// @Description Get all units
// @Tags Units
// @Success 200 {object} []models.Unit
// @Router /units [get]
func (c *unitController) GetAllUnits(ctx *fiber.Ctx) error {
	units, err := c.unitService.GetAllUnits()
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"units": units})
}