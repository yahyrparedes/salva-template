package test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yahyrparedes/salva-template/cmd/config"
	"github.com/yahyrparedes/salva-template/cmd/database"
	"github.com/yahyrparedes/salva-template/cmd/server"
	"github.com/yahyrparedes/salva-template/pkg/middleware"
	"github.com/yahyrparedes/salva-template/pkg/routes"
	"gorm.io/gorm"
)

func InitializeConfigTest() (*fiber.App, *gorm.DB) {

	config.ReadConfigFile(config.RootPathTest)
	app := server.InitializeServer()
	db := database.InitializeTestConnection()
	// Middlewares.
	middleware.FiberMiddleware(app) // cors  --

	// Routes.
	routes.ConfigRoutes(app) // Register a public routes for app.

	return app, db
}
