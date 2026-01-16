package routes

import "github.com/gofiber/fiber/v2"

func SetupRouter(app *fiber.App, reviewRouter *ReviewRouter) {
	api := app.Group("/api")
	reviewRouter.Register(api.Group("/reviews"))
}
