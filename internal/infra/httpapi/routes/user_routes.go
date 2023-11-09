package routes

import (
	// "github.com/fseda/cookbooked-api/internal/domain/models"
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
	recipeRepository := repositories.NewRecipeRepository(db)
	recipeIngredientRepository := repositories.NewRecipeIngredientRepository(db)
	ingredientRepository := repositories.NewIngredientRepository(db)
	unitRepository := repositories.NewUnitRepository(db)

	userService := services.NewUserService(userRepository)
	recipeService := services.NewRecipeService(recipeRepository,
		recipeIngredientRepository,
		ingredientRepository,
		unitRepository,
	)

	userController := controllers.NewUserController(userService)
	recipeController := controllers.NewRecipeController(recipeService)

	userGroup := app.Group("users")
	userGroup.Get("exists", userController.UserExists)
	userGroup.Get(":id",
		middlewares.ValidateID("id"),
		middlewares.JWTAuthMiddleware(env.Http.JWTSecretKey),
		middlewares.RoleRequired(models.ADMIN),
		userController.GetOneByID,
	)

	meGroup := app.Group("me", middlewares.JWTAuthMiddleware(env.Http.JWTSecretKey))
	meGroup.Get("", userController.Profile)
	meGroup.Get("", userController.Delete)

	meRecipeGroup := meGroup.Group("recipes")
	meRecipeGroup.Get("",
		recipeController.GetRecipesByUserID,
	)
	meRecipeGroup.Get(":recipe_id",
		middlewares.ValidateID("recipe_id"),
		recipeController.GetRecipeDetails,
	)
	meRecipeGroup.Post(":recipe_id/ingredients",
		middlewares.ValidateID("recipe_id"),
		recipeController.AddRecipeIngredient,
	)
	meRecipeGroup.Post(":recipe_id/ingredients/bulk",
		middlewares.ValidateID("recipe_id"),
		recipeController.AddRecipeIngredients,
	)
	meRecipeGroup.Delete(":recipe_id/ingredients/:recipe_ingredient_id",
		middlewares.ValidateID("recipe_id"),
		middlewares.ValidateID("recipe_ingredient_id"),
		recipeController.RemoveRecipeIngredient,
	)
}
