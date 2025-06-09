package users

import (
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/ethanbaker/api/pkg/config"
	auth "github.com/ethanbaker/api/pkg/jwt_auth"
	"github.com/gin-gonic/gin"
)

const identityKey = "username"

// Initialize jwt auth with the secret key
func init() {
	key, ok := config.GetEnv("JWT_SECRET")
	if !ok {
		log.Fatal("[ERR]: JWT_SECRET environment variable not set")
	}

	if err := auth.SetSecretKey(key); err != nil {
		log.Fatalf("[ERR]: error setting JWT secret key (%v)\n", err)
	}
}

// Return the claimed identity of the user
func identityHandler(c *gin.Context) any {
	// Extract the claims from the context
	claims := jwt.ExtractClaims(c)

	return &User{
		Username: claims[identityKey].(string),
	}
}

// Handle an incoming authentication request by verifying the user's credentials
func authenticator(c *gin.Context) (any, error) {
	// Get the JWT request
	var req auth.JwtRequest
	if err := c.ShouldBind(&req); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	// Extract the user's credentials
	username := req.Username
	password := req.Password

	// Validate the user's credentials

	// NOTE: This is a placeholder to demonstrate jwt authentication. An actual
	// service would have more complex logic to validate the user's credentials, such as
	// checking a database or external service

	// NOTE: You should NOT use this code in any live server. Storing/checking passwords
	// in plaintext is NOT secure. This is only for demonstration purposes

	if username == "admin" && password == "admin" {
		return &User{
			Username: username,
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

// Handle an incoming authorization request by verifying the user's identity
func authorizator(data any, c *gin.Context) bool {
	// NOTE: This is a placeholder to demonstrate jwt authorization. An actual
	// service would have more complex logic to validate the user's identity, such as
	// checking a database or external service

	// Extract the user's identity
	if v, ok := data.(*User); ok && v.Username == "admin" {
		return true
	}

	return false
}
