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

func (c *unitController) GetAllUnits(ctx *fiber.Ctx) error {
	units, err := c.unitService.GetAllUnits()
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"units": units})
}