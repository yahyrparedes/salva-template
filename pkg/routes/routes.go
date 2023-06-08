package routes

import "github.com/gofiber/fiber/v2"

func ConfigRoutes(app *fiber.App) {
	SwaggerRoute(app)
	BrandRoute(app)
}
