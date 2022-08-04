package controllers

import (
	"github.com/RyotaKITA-12/locatask.git/database"
	"github.com/RyotaKITA-12/locatask.git/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match.",
		})
	}

	user := models.User{
		FirstName: data["first"],
		LastName:  data["last"],
		Email:     data["email"],
		Password:  data["password"],
	}

	database.DB.Create(&user)

	return c.JSON(user)
}
