// Handle JWT authentication using the gin-jwt middleware
package jwt_auth

import (
	"fmt"
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var secret []byte

// SetSecretKey sets the JWT secret key used for signing tokens
func SetSecretKey(jwtSecret string) error {
	if jwtSecret == "" {
		return fmt.Errorf("jwt secret cannot be empty")
	}

	// Load secret
	secret = []byte(jwtSecret)

	log.Printf("[JWT_AUTH|INFO]: jwt middleware initialized with secret")

	return nil
}

// Middleware handler for JWT authentication
func MiddlewareHandler(mw *jwt.GinJWTMiddleware) (gin.HandlerFunc, error) {
	// Check if the middleware is nil
	if mw == nil {
		return nil, fmt.Errorf("jwt middleware cannot be nil")
	}

	return func(context *gin.Context) {
		if err := mw.MiddlewareInit(); err != nil {
			log.Fatalf("[JWT_AUTH|ERR]: Failed to initialize JWT middleware (%s)\n", err)
		}
	}, nil
}

// Register JWT routes with the given middleware. An optional prefix can be provided
func RegisterRoute(g *gin.RouterGroup, handle *jwt.GinJWTMiddleware, prefix ...string) error {
	if len(prefix) > 1 {
		// Make sure prefix is not greater than 1
		return fmt.Errorf("prefix must be a single string, got %d elements", len(prefix))
	} else if len(prefix) == 1 && prefix[0] == "" {
		// Make sure provided prefix is not empty
		return fmt.Errorf("prefix cannot be an empty string")
	} else if len(prefix) == 0 {
		prefix = []string{"auth"}
	}

	// Register the routes under the given prefix
	auth := g.Group(prefix[0])

	auth.POST("/login", handle.LoginHandler)
	auth.GET("/refresh", handle.MiddlewareFunc(), handle.RefreshHandler)
	auth.DELETE("/logout", handle.MiddlewareFunc(), handle.LogoutHandler)

	// Log and return
	log.Printf("[JWT_AUTH|INFO]: JWT routes registered with prefix '%s'", prefix[0])
	return nil
}
