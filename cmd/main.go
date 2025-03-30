package main

import (
	"log"
	"net/http"

	"github.com/ethanbaker/api/config"
	"github.com/ethanbaker/api/internal/jwt_auth"
	"github.com/ethanbaker/api/modules/health"
	"github.com/ethanbaker/api/modules/key"
	"github.com/ethanbaker/api/modules/users"
	"github.com/gin-gonic/gin"
)

func main() {
	// Get config values from environment
	port := config.GetEnvWithDefault("PORT", "8080")

	// Add app-level settings/routes
	engine := gin.Default()
	engine.NoRoute(jwt_auth.HandleNoRoute)

	// Add custom modules
	health.RegisterRoutes(engine)
	key.RegisterRoutes(engine)
	users.RegisterRoutes(engine)

	// start http server
	if err := http.ListenAndServe(":"+port, engine); err != nil {
		log.Fatal(err)
	}
}
