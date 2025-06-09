package health

import "github.com/gin-gonic/gin"

// RegisterRoutes registers the routes for the health module
func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", getStatus)
}
