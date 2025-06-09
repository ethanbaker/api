// Handle authentication with api keys
package api_key

import (
	"net/http"

	"github.com/ethanbaker/api/pkg/api_types"
	"github.com/gin-gonic/gin"
)

// APIKeyHeaderHandler is a middleware handler for API key authentication. This method
// is provided with a function that validates the API key based on a provided 'X-API-KEY' header
func APIKeyHeaderHandler(validate func(string) bool) gin.HandlerFunc {
	// Make sure validate is provided
	if validate == nil {
		panic("validate function cannot be nil")
	}

	return func(c *gin.Context) {
		// Get the API key from the request header
		apiKey := c.GetHeader("X-API-KEY")

		if apiKey == "" || !validate(apiKey) {
			// If api key is invalid or missing, return forbidden
			res := api_types.NewFailResponse(http.StatusForbidden, "Invalid API Key")
			c.JSON(res.AsGinResponse())
			c.Abort()
			return
		}

		c.Next()
	}
}

// APIKeyQueryHandler is a middleware handler for API key authentication. This method
// is provided with a function that validates the API key based on a provided
// named query parameter
func APIKeyQueryHandler(param string, validate func(string) bool) gin.HandlerFunc {
	// Make sure validate is provided
	if validate == nil {
		panic("validate function cannot be nil")
	}

	// Make sure param is provided
	if param == "" {
		panic("query parameter name cannot be empty")
	}

	return func(c *gin.Context) {
		// Get the API key from the query parameter
		apiKey := c.Query(param)

		if apiKey == "" || !validate(apiKey) {
			// If api key is invalid or missing, return forbidden
			res := api_types.NewFailResponse(http.StatusForbidden, "Invalid API Key")
			c.JSON(res.AsGinResponse())
			c.Abort()
			return
		}

		c.Next()
	}
}
