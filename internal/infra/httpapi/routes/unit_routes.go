package routes

import (
	"github.com/fseda/cookbooked-api/internal/domain/repositories"
	"github.com/fseda/cookbooked-api/internal/domain/services"
	"github.com/fseda/cookbooked-api/internal/infra/config"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func loadUnitRoutes(app *fiber.App, db *gorm.DB, env *config.Config) {
	unitRepository := repositories.NewUnitRepository(db)
	unitService := services.NewUnitService(unitRepository)
	unitController := controllers.NewUnitController(unitService)

	unitGroup := app.Group("units")
	unitGroup.Get("", unitController.GetAllUnits)
}
