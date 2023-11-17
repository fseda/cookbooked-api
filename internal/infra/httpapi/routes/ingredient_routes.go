package routes

import (
	"github.com/fseda/cookbooked-api/internal/domain/repositories"
	"github.com/fseda/cookbooked-api/internal/domain/services"
	"github.com/fseda/cookbooked-api/internal/infra/config"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func loadIngredientRoutes(app *fiber.App, db *gorm.DB, env *config.Config) {
	ingredientRepository := repositories.NewIngredientRepository(db)
	ingredientService := services.NewIngredientService(ingredientRepository)
	ingredientController := controllers.NewIngredientController(ingredientService)

	ingredientGroup := app.Group("ingredients")
	ingredientGroup.Get("", ingredientController.GetAllIngredients)
}
