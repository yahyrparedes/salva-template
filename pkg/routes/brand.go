package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yahyrparedes/salva-template/pkg/controllers"
)

func BrandRoute(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Get("/brands", controllers.GetBrands)   // Lista
	route.Get("/brand/:id", controllers.GetBrand) // Detalle
}
