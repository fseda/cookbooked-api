package routes

import (
	"github.com/fseda/cookbooked-api/internal/domain/repositories"
	"github.com/fseda/cookbooked-api/internal/domain/services"
	"github.com/fseda/cookbooked-api/internal/infra/config"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/controllers"
	middlewares "github.com/fseda/cookbooked-api/internal/infra/httpapi/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func loadRecipeRoutes(app *fiber.App, db *gorm.DB, env *config.Config) {
	recipeRepository := repositories.NewRecipeRepository(db)
	ingredientRepository := repositories.NewIngredientRepository(db)
	recipeIngredientRepository := repositories.NewRecipeIngredientRepository(db)
	unitRepository := repositories.NewUnitRepository(db)
	recipeService := services.NewRecipeService(recipeRepository, recipeIngredientRepository, ingredientRepository, unitRepository)
	recipeController := controllers.NewRecipeController(recipeService)

	recipeGroup := app.Group("recipes", middlewares.JWTAuthMiddleware(env.Http.JWTSecretKey))
	recipeGroup.Post("new", recipeController.CreateRecipe)
	recipeGroup.Get(
		":recipe_id",
		middlewares.ValidateID("recipe_id"),
		recipeController.GetRecipeDetails,
	)
	recipeGroup.Patch(
		":recipe_id/ingredients/single",
		middlewares.ValidateID("recipe_id"),
		recipeController.AddRecipeIngredient,
	)
	recipeGroup.Patch(
		":recipe_id/ingredients",
		middlewares.ValidateID("recipe_id"),
		recipeController.AddRecipeIngredients,
	)
	recipeGroup.Delete(
		":recipe_id/ingredients/:recipe_ingredient_id",
		middlewares.ValidateID("recipe_id"),
		middlewares.ValidateID("recipe_ingredient_id"),
		recipeController.RemoveRecipeIngredient,
	)
}
