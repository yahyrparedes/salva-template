package main

import (
	"encoding/json"
	"fmt"
	"github.com/yahyrparedes/salva-template/cmd/config"
	"github.com/yahyrparedes/salva-template/cmd/database"
	"github.com/yahyrparedes/salva-template/cmd/server"
	_ "github.com/yahyrparedes/salva-template/docs"
	"github.com/yahyrparedes/salva-template/pkg/middleware"
	"github.com/yahyrparedes/salva-template/pkg/routes"
)

func main() {
	config.InitializeBasicConfig()
	database.InitializeConnection()
	app := server.InitializeServer()

	// Middlewares
	middleware.FiberMiddleware(app) // cors  --

	// Routes
	routes.ConfigRoutes(app)

	if config.IsLocal() {
		// Print the router stack in JSON format
		data, _ := json.MarshalIndent(app.GetRoutes(true), "", "  ")
		fmt.Println(string(data))
	}

	server.RunServer(app)
}
