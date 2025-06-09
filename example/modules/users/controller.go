package users

// NOTE: Other API endpoints can be added here to fit implementation requirements
// The method below is a placeholder to demonstrate JWT validation

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/ethanbaker/api/pkg/api_types"
	"github.com/gin-gonic/gin"
)

// Return a response to the user who called this method
func getResponse(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, ok := c.Get(identityKey)
	if !ok {
		res := api_types.NewErrorResponse(http.StatusNotFound, "User not found")
		c.JSON(res.AsGinResponse())
		return
	}

	log.Printf("[INFO]: user called getResponse: %#v\n", user)

	res := api_types.NewSuccessResponse(
		fmt.Sprintf("Hello %s!", user.(*User).Username),
		gin.H{"username": claims[identityKey]},
	)
	c.JSON(res.AsGinResponse())
}

// Return an anonymous response (no JWT security)
func getAnonResponse(c *gin.Context) {
	c.JSON(api_types.NewSuccessResponse("Hello, anonymous user!", nil).AsGinResponse())
}
