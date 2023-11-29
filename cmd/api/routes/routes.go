package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/general-market-GO/cmd/api/controllers"
)

func InitRoutes(app *fiber.App, controllers []controllers.Controller) {
	for _, controller := range controllers {
		controller.RegisterRoutes(app)
	}
}
