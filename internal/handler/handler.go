package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucaslimafernandes/desafio-1-1s-vs-3j/internal/models"
)

func PostUsers(c *fiber.Ctx) error {

	users := new([]models.User)

	err := c.BodyParser(users)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error(), "user_count": 0})
	}

	response := fiber.Map{
		"message":    "Arquivo recebido com sucesso",
		"user_count": len(*users),
	}

	return c.Status(200).JSON(response)

}

func GetSuperUsers(c *fiber.Ctx) {

}

func GetTopCountries(c *fiber.Ctx) {

}

func GetTeamInsights(c *fiber.Ctx) {

}

func GetActiveUsersPerDay(c *fiber.Ctx) {

}

func GetEvaluation(c *fiber.Ctx) {

}
