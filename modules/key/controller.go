package key

import "github.com/gin-gonic/gin"

// NOTE: Other API endpoints can be added here to fit implementation requirements
// The method below is a placeholder to demonstrate API key validation

// Return a success response on a call
func getResponse(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}
