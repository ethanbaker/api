package key

// NOTE: This module demonstrates API key verification

import (
	"github.com/ethanbaker/api/internal/api_key"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers the routes for the key module
func RegisterRoutes(r *gin.Engine) {
	group := r.Group("/key")

	group.GET("/response", api_key.APIKeyMiddlewareHandler(validateAPIKey), getResponse)
}
