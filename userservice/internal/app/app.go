package app

import (
	"fmt"
	"log"

	"github.com/faruqii/soa/userservice/internal/repository"
	"github.com/faruqii/soa/userservice/internal/routes"
	"github.com/faruqii/soa/userservice/internal/service"
	"github.com/faruqii/soa/userservice/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start() {
	app := fiber.New()

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	repository := repository.NewUserRepository(db)
	userService := service.NewUserService(repository)
	routes.UserRoutes(app, userService)

	if err := app.Listen(":3001"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	fmt.Println("User service is running on port 3001")

	app.Use(logger.New(logger.Config{
		Format: "${time} | ${ip} | ${status} | ${method} ${path} | ${latency}\n",
	}))

}
