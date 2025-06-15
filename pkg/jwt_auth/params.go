package jwt_auth

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/ethanbaker/api/pkg/api_types"
	"github.com/gin-gonic/gin"
)

// JwtParams represents the parameters provided for JWT authentication. These are passed by
// the user to the 'Param' function. These are specified by Gin JWT middleware
type JwtParams struct {
	SigningAlgorithm string `json:"signing_algorithm"`
	Timeout          string `json:"timeout"`
	MaxRefresh       string `json:"max_refresh"`
	TokenLookup      string `json:"token_lookup"`
	TokenHeadName    string `json:"token_head_name"`
	IdentityKey      string `json:"identity_key"`
}

// Generate default JWT parameters for the middleware
func GenerateParams[T User](params JwtParams) (*jwt.GinJWTMiddleware, error) {
	log.Printf("[JWT_AUTH]: generating JWT middleware parameters with values: %#v", params)

	// Read signing algorithm
	signingAlgorithm := orDefault(params.SigningAlgorithm, "HS256")

	// Read timeout value
	timeoutStr := orDefault(params.Timeout, "1h")
	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT expiration duration: %w", err)
	}

	// Read max refresh value
	maxRefreshStr := orDefault(params.MaxRefresh, "1h")
	maxRefresh, err := time.ParseDuration(maxRefreshStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT max refresh duration: %w", err)
	}

	// Read token lookup
	tokenLookup := orDefault(params.TokenLookup, "header: Authorization, query: token, cookie: jwt")

	// Read token head name
	tokenHeadName := orDefault(params.TokenHeadName, "Bearer")

	// Read identity key
	identityKey := params.IdentityKey
	if identityKey == "" {
		return nil, fmt.Errorf("identity key cannot be empty")
	}

	log.Printf("[JWT_AUTH]: successfully generated JWT middleware")

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
		PayloadFunc:  generatePayloadFunc[T](identityKey),
		Unauthorized: unauthorized,

		/**
		These functions need to be defined elsewhere and initialized in the main package:
			- IdentityHandler
			- Authenticator
			- Authorizator

		This is for any custom implementation of a 'User' object, which optimally shouldn't
		be defined in this package. In addition, this allows for multiple JWT auth implementations
		on different modules, if your implementation requires it
		*/
	}, nil
}

// Handle an incoming payload from a JWT request
func generatePayloadFunc[T User](identityKey string) func(data any) jwt.MapClaims {
	return func(data any) jwt.MapClaims {
		// Cast incoming data into a user
		v, ok := data.(T)
		if !ok {
			log.Printf("[JWT_AUTH]: warning, failed to cast incoming data %#v\n", data)
			return jwt.MapClaims{}
		}

		// Return the user's username as the identity key
		return jwt.MapClaims{
			identityKey: v.GetUsername(),
		}
	}
}

// Handle unauthorized requests
func unauthorized(c *gin.Context, code int, message string) {
	res := api_types.NewFailResponse(code, message)
	c.JSON(res.AsGinResponse())
}

// Helper function to return either given value or default
func orDefault(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
