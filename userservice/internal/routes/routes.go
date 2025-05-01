package routes

import (
	"github.com/faruqii/soa/userservice/internal/handler"
	"github.com/faruqii/soa/userservice/internal/service"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router, userService service.UserService) {
	userHandler := handler.NewUserHandler(userService)

	// User routes
	UserRoutes := router.Group("/api/v1/users")
	UserRoutes.Post("/register", userHandler.CreateUser)
	UserRoutes.Get("/:id", userHandler.GetUserByID)
}
