package users

import (
	"fmt"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// NOTE: Other API endpoints can be added here to fit implementation requirements
// The method below is a placeholder to demonstrate JWT validation

// Return a response to the user who called this method
func getResponse(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello %s!", user.(*User).First), "username": claims[identityKey]})
}

// Return an anonymous response (no JWT security)
func getAnonResponse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello anonymous!"})
}
