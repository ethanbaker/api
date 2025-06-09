package users

import (
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	auth "github.com/ethanbaker/api/pkg/jwt_auth"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers the routes for the key module
func RegisterRoutes(r *gin.Engine) {
	// Create route group
	group := r.Group("/users")

	// Add JWT logic
	params, err := auth.GenerateParams[*User](auth.JwtParams{IdentityKey: "username"})
	if err != nil {
		log.Fatalf("[ERR]: error generating JWT middleware parameters (%v)\n", err)
	}

	params.IdentityHandler = identityHandler
	params.Authenticator = authenticator
	params.Authorizator = authorizator

	middleware, err := jwt.New(params)
	if err != nil {
		log.Fatalf("[ERR]: error creating JWT middleware (%v)\n", err)
	}

	mwh, err := auth.MiddlewareHandler(middleware)
	if err != nil {
		log.Fatalf("[ERR]: error creating JWT middleware handler (%v)\n", err)
	}

	group.Use(mwh)
	auth.RegisterRoute(group, middleware)

	// Add custom routes
	group.GET("/response", middleware.MiddlewareFunc(), getResponse)
	group.GET("/anon-response", getAnonResponse)
}
