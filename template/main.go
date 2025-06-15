package main

import (
	"log"
	"time"

	"github.com/ethanbaker/api/pkg/config"
	"github.com/ethanbaker/api/pkg/utils"
	"github.com/ethanbaker/api/template/modules/custom"
	"github.com/ethanbaker/api/template/modules/health"
	"github.com/gin-contrib/cors"
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

	// Add trusted proxies
	engine.SetTrustedProxies(nil)

	// Add CORS using gin-contrib/cors (https://github.com/gin-contrib/cors for documentation)
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // Your custom origins here
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"}, // Your custom methods here
		AllowHeaders:     []string{"Origin", "Content-Type"},                  // Your custom headers here
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	// Base group '/api' for all API routes
	baseGroup := engine.Group("/api")

	// - Adding custom modules
	health.RegisterRoutes(baseGroup)
	custom.RegisterRoutes(baseGroup)

	// Then after performing initial setup, start the server
	if err := engine.Run(":" + port); err != nil {
		log.Fatal("[MAIN]: Failed to start server: ", err)
	}
}
