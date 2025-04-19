package routes

import (
	"kasir/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Post("/users", controllers.CreateUser)
	app.Get("/users", controllers.ReadUsersPaginated)
	app.Get("/users/:id", controllers.ReadUserByID)
	app.Put("/users/:id", controllers.UpdateUser)
	app.Delete("/users/:id", controllers.DeleteUser)
	app.Delete("/users", controllers.DeleteAllUsers)
}
