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

func addAuthRoutes(app *fiber.App, db *gorm.DB, env *config.Config) {
	userRepository := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepository, env)
	authController := controllers.NewAuthController(authService)

	auth := app.Group("/auth")

	auth.Post("/login", authController.Login)
	auth.Post("/signup", authController.RegisterUser)

	auth.Get(
		"/profile",
	 	middlewares.JWTAuthMiddleware(env.Http.JWTSecretKey),
		authController.Profile,
	)
}