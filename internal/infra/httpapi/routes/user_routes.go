package routes

import (
	"github.com/fseda/cookbooked-api/internal/domain/repositories"
	"github.com/fseda/cookbooked-api/internal/domain/services"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func addUserRoutes(app *fiber.App, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	users := app.Group("/users")
	users.Post("/", userController.Create)
	users.Get("/:id", userController.FindOne)
}
