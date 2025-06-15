package jwt_auth

import (
	"log"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/ethanbaker/api/pkg/api_types"
	"github.com/gin-gonic/gin"
)

// Handle requests with no route and log the claims
func HandleNoRoute(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	log.Printf("[JWT_AUTH]: NoRoute claims: %#v\n", claims)

	res := api_types.NewFailResponse(http.StatusNotFound, "Route not found")
	c.JSON(res.AsGinResponse())
}
