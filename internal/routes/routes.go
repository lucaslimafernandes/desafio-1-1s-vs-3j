package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucaslimafernandes/desafio-1-1s-vs-3j/internal/handler"
)

func SetupRoutes(app *fiber.App) {

	// POST
	user := app.Group("/users")
	user.Post("/", handler.PostUsers)

}
