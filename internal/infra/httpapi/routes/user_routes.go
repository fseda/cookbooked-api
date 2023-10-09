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

func loadUserRoutes(app *fiber.App, db *gorm.DB, env *config.Config) {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	userGroup := app.Group("/user")
	// userGroup.Get("/:id", userController.FindOne)
	// userGroup.Delete("/:id", userController.Delete)

	userGroup.Get("/me", middlewares.JWTAuthMiddleware(env.Http.JWTSecretKey), userController.Profile)
}
