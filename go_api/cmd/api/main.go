package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/lucaslimafernandes/desafio-1-1s-vs-3j/internal/routes"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func main() {

	app := fiber.New()
	app.Use(cors.New())

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("PORT")))

}
