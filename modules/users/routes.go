package users

import (
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	auth "github.com/ethanbaker/api/internal/jwt_auth"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers the routes for the key module
func RegisterRoutes(r *gin.Engine) {
	// Create route group
	group := r.Group("/users")

	// Add JWT logic
	params := auth.Params[*User]()
	params.IdentityHandler = identityHandler
	params.Authenticator = authenticator
	params.Authorizator = authorizator

	identityKey = params.IdentityKey

	middleware, err := jwt.New(params)
	if err != nil {
		log.Fatalf("[ERR]: error creating JWT middleware (%v)\n", err)
	}

	group.Use(auth.MiddlewareHandler(middleware))
	auth.RegisterRoute(group, middleware)

	// Add custom routes
	group.GET("/response", middleware.MiddlewareFunc(), getResponse)
	group.GET("/anon-response", getAnonResponse)
}
