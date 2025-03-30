// Handle authentication with api keys
package api_key

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIKeyMiddlewareHandler is a middleware handler for API key authentication. This method
// is provided with a function that validates the API key
func APIKeyMiddlewareHandler(validate func(string) bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the API key from the request header
		apiKey := c.GetHeader("X-API-KEY")

		if apiKey == "" || !validate(apiKey) {
			// If api key is invalid or missing, return forbidden
			c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
			c.Abort()
			return
		}
		c.Next()
	}
}
