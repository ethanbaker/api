package main

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/ethanbaker/api/example/modules/health"
	"github.com/ethanbaker/api/example/modules/key"
	"github.com/ethanbaker/api/example/modules/users"
	"github.com/ethanbaker/api/pkg/api_types"
	"github.com/ethanbaker/api/pkg/config"
	"github.com/ethanbaker/api/pkg/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// Get config values from environment
	port := config.GetEnvWithDefault("PORT", "8080")

	// Add app-level settings/routes
	engine := gin.Default()
	engine.NoRoute(utils.NoRouteHandler)

	// Add custom modules
	health.RegisterRoutes(engine)
	key.RegisterRoutes(engine)
	users.RegisterRoutes(engine)

	// Catch panics and log them
	engine.Use(gin.RecoveryWithWriter(log.Writer(), func(c *gin.Context, recovered interface{}) {
		// Log the panic value
		log.Printf("panic recovered: %v", recovered)
		debug.PrintStack()

		// Return a JSON error response
		res := api_types.NewErrorResponse(http.StatusInternalServerError, "Internal Server Error")
		c.AbortWithStatusJSON(res.AsGinResponse())
	}))

	// start http server
	if err := http.ListenAndServe(":"+port, engine); err != nil {
		log.Fatal(err)
	}
}
