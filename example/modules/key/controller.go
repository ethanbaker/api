package key

// NOTE: Other API endpoints can be added here to fit implementation requirements
// The method below is a placeholder to demonstrate API key validation

import (
	"github.com/ethanbaker/api/pkg/api_types"
	"github.com/gin-gonic/gin"
)

// Return a success response on a call
func getResponse(c *gin.Context) {
	c.JSON(api_types.NewSuccessResponse("API Key is valid", nil).AsGinResponse())
}
