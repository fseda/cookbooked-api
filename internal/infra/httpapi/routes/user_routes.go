package routes

import (
	"github.com/fseda/cookbooked-api/internal/domain/models"
	"github.com/fseda/cookbooked-api/internal/domain/repositories"
	"github.com/fseda/cookbooked-api/internal/domain/services"
	"github.com/fseda/cookbooked-api/internal/infra/config"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/controllers"
	middlewares "github.com/fseda/cookbooked-api/internal/infra/httpapi/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func loadUserRoutes(app *fiber.App, db *gorm.DB, env *config.Config) {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	recipeRepository := repositories.NewRecipeRepository(db)
	ingredientRepository := repositories.NewIngredientRepository(db)
	unitRepository := repositories.NewUnitRepository(db)
	recipeService := services.NewRecipeService(recipeRepository, ingredientRepository, unitRepository)
	recipeController := controllers.NewRecipeController(recipeService)

	userGroup := app.Group("users")
	userGroup.Get(
		":id",
		middlewares.ValidateID(),
		middlewares.JWTAuthMiddleware(env.Http.JWTSecretKey),
		middlewares.RoleRequired(models.ADMIN),
		userController.GetOneByID,
	)

	meGroup := userGroup.Group("me")
	meGroup.Get("", middlewares.JWTAuthMiddleware(env.Http.JWTSecretKey), userController.Profile)
	meGroup.Get("recipes", middlewares.JWTAuthMiddleware(env.Http.JWTSecretKey), recipeController.GetRecipesByUserID)
	meGroup.Get(
		"/recipes/:id",
		middlewares.JWTAuthMiddleware(env.Http.JWTSecretKey),
		middlewares.ValidateID(),
		recipeController.GetRecipeDetails,
	)
}
