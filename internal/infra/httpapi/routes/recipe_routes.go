package routes

import (
	"github.com/fseda/cookbooked-api/internal/domain/repositories"
	"github.com/fseda/cookbooked-api/internal/domain/services"
	"github.com/fseda/cookbooked-api/internal/infra/config"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/controllers"
	middlewares "github.com/fseda/cookbooked-api/internal/infra/httpapi/middleware/auth"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func loadRecipeRoutes(app *fiber.App, db *gorm.DB, env *config.Config) {
	recipeRepository := repositories.NewRecipeRepository(db)
	ingredientRepository := repositories.NewIngredientRepository(db)
	unitRepository := repositories.NewUnitRepository(db)
	recipeService := services.NewRecipeService(recipeRepository, ingredientRepository, unitRepository)
	recipeController := controllers.NewRecipeController(recipeService)

	recipeGroup := app.Group("/recipes")
	recipeGroup.Post("/new", middlewares.JWTAuthMiddleware(env.Http.JWTSecretKey), recipeController.CreateRecipe)
	recipeGroup.Get("/mine", middlewares.JWTAuthMiddleware(env.Http.JWTSecretKey), recipeController.GetRecipesByUserID)
}
