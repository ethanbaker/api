// Handle JWT authentication using the gin-jwt middleware
package jwt_auth

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/ethanbaker/api/config"
	"github.com/gin-gonic/gin"
)

// Middleware handler for JWT authentication
func MiddlewareHandler(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := authMiddleware.MiddlewareInit(); err != nil {
			log.Fatalf("[MIDDLEWARE_ERR]: Failed to initialize JWT middleware (%s)\n", err)
		}
	}
}

// Register JWT routes with the given middleware
func RegisterRoute(g *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	auth := g.Group("/auth")

	auth.POST("/login", handle.LoginHandler)
	auth.GET("/refresh", handle.MiddlewareFunc(), handle.RefreshHandler)
	auth.DELETE("/logout", handle.MiddlewareFunc(), handle.LogoutHandler)
}

// Generate default JWT parameters for the middleware
func Params[T user]() *jwt.GinJWTMiddleware {
	// Read signing algorithm from environment variables
	signingAlgorithm := config.GetEnvWithDefault("JWT_SIGNING_ALGORITHM", "HS256")

	// Read timeout value from environment variables
	timeoutStr := config.GetEnvWithDefault("JWT_TIMEOUT", "1h")
	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		log.Fatalf("[MIDDLEWARE_ERR]: Failed to parse 'JWT_EXPIRATION' environment variable (%s)\n", err)
	}

	// Read max refresh value from environment variables
	maxRefreshStr := config.GetEnvWithDefault("JWT_MAX_REFRESH", "1h")
	maxRefresh, err := time.ParseDuration(maxRefreshStr)
	if err != nil {
		log.Fatalf("[MIDDLEWARE_ERR]: Failed to parse 'JWT_MAX_REFRESH' environment variable (%s)\n", err)
	}

	// Read token lookup from environment variables
	tokenLookup := config.GetEnvWithDefault("JWT_TOKEN_LOOKUP", "header: Authorization, query: token, cookie: jwt")

	// Read token head name from environment variables
	tokenHeadName := config.GetEnvWithDefault("JWT_TOKEN_HEAD_NAME", "Bearer")

	return &jwt.GinJWTMiddleware{
		// Statically defined here
		Realm:    "api-jwt-middleware",
		TimeFunc: time.Now,

		// Can be defined in environment variables
		SigningAlgorithm: signingAlgorithm,
		Key:              secret,
		Timeout:          timeout,
		MaxRefresh:       maxRefresh,
		IdentityKey:      identityKey,
		TokenLookup:      tokenLookup,
		TokenHeadName:    tokenHeadName,

		// Functions
		PayloadFunc:  payloadFunc[T],
		Unauthorized: unauthorized,

		/**
		These functions need to be defined elsewhere and initialized in the main package:
			- IdentityHandler
			- Authenticator
			- Authorizator

		This is for any custom implementation of a 'User' object, which optimally shouldn't
		be defined in this package. In addition, this allows for multiple JWT auth implementations
		on different modules, if your implementation requires it.
		*/
	}
}

// Handle an incoming payload from a JWT request
func payloadFunc[T user](data any) jwt.MapClaims {
	// Cast incoming data into a user
	v, ok := data.(T)
	if !ok {
		log.Printf("[MIDDLEWARE_WARN]: Failed to cast incoming data\n")
		return jwt.MapClaims{}
	}

	// Return the user's username as the identity key
	return jwt.MapClaims{
		identityKey: v.GetUsername(),
	}
}

// Handle unauthorized requests
func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"message": message})
}
