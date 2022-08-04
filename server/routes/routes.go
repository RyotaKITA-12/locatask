package routes

import (
	"github.com/RyotaKITA-12/locatask.git/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/api/user", controllers.User)
	app.Get("/api/logout", controllers.Logout)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
}
