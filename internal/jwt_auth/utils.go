package jwt_auth

import (
	"log"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// Handle requests with no route
func HandleNoRoute(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	log.Printf("[INFO]: NoRoute claims: %#v\n", claims)

	c.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
}
