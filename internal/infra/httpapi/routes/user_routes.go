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
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

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
}
