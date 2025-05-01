package main

import (
	"fmt"

	"github.com/faruqii/soa/authservice/internal/handler"
	"github.com/faruqii/soa/authservice/internal/repository"
	"github.com/faruqii/soa/authservice/internal/service"
	"github.com/faruqii/soa/authservice/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Connect to the database
	app := fiber.New()

	db, err := database.Connect()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	repo := repository.NewKeyRepository(db)
	service := service.NewKeyService(repo)
	handler := handler.NewKeyHandler(service)

	app.Post("/keys", handler.CreateKey)
	app.Get("/keys/:key", handler.ValidateKey)

	if err := app.Listen(":3000"); err != nil {
		fmt.Println("Error starting the server:", err)
		return
	}
	fmt.Println("Server is running on port 3000")

}
