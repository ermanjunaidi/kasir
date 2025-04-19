package main

import (
	"fmt"
	"log"

	"kasir/config"
	"kasir/models"
	"kasir/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()

	// AutoMigrate all models
	config.DB.AutoMigrate(&models.User{})

	app := fiber.New()

	routes.UserRoutes(app)

	fmt.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}
