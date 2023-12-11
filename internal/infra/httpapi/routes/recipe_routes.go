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
	recipeService := services.NewRecipeService(recipeRepository)
	recipeController := controllers.NewRecipeController(recipeService)

	recipeGroup := app.Group("recipes", middlewares.JWTAuthMiddleware(env.Http.JWTSecretKey))
	recipeGroup.Post("new", recipeController.CreateRecipe)
	recipeGroup.Get("", recipeController.GetAllRecipesByUserID)
	recipeGroup.Get(
		":recipe_id",
		middlewares.ValidateID("recipe_id"),
		recipeController.GetRecipeDetails,
	)
	recipeGroup.Patch(
		":recipe_id",
		middlewares.ValidateID("recipe_id"),
		recipeController.UpdateRecipe,
	)
	recipeGroup.Delete(
		":recipe_id",
		middlewares.ValidateID("recipe_id"),
		recipeController.DeleteRecipe,
	)
}
