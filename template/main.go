package main

import (
	"log"
	"net/http"

	"github.com/ethanbaker/api/pkg/config"
	"github.com/ethanbaker/api/pkg/utils"
	"github.com/ethanbaker/api/template/modules/health"
	"github.com/gin-gonic/gin"
)

func main() {
	// This is the main entry point for the application
	// The actual implementation will depend on the specific requirements of your application
	// You can initialize your server, set up routes, and start listening for requests here
	// For example, you might want to call a function to set up your HTTP server and routes

	// Common tasks might include:

	// - Initializing configuration settings
	port := config.GetEnvWithDefault("PORT", "8080")

	// - Add app level settings/routes
	engine := gin.Default()
	engine.NoRoute(utils.NoRouteHandler)

	// - Adding custom modules
	health.RegisterRoutes(engine)

	// Then after performing initial setup, start the server
	if err := http.ListenAndServe(":"+port, engine); err != nil {
		log.Fatal(err)
	}
}
