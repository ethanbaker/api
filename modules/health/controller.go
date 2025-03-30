package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Return status of the API
func getStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
