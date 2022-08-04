package main

import (
	"github.com/RyotaKITA-12/locatask.git/database"
	"github.com/RyotaKITA-12/locatask.git/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8888")
}
