package key

// NOTE: This module demonstrates API key verification. This file contains the initial
// route setup with Gin

import (
	"github.com/ethanbaker/api/pkg/api_key"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers the routes for the key module
func RegisterRoutes(r *gin.Engine) {
	group := r.Group("/key")

	group.GET("/response", api_key.APIKeyHeaderHandler(validateAPIKey), getResponse)
}
