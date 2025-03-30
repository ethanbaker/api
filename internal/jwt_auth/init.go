package jwt_auth

import (
	"log"

	"github.com/ethanbaker/api/config"
)

var secret []byte
var identityKey string

// Initialize important JWT variables
func init() {
	// Load secret
	s, ok := config.GetEnv("JWT_SECRET")
	if !ok {
		log.Fatal("[MIDDLEWARE_ERR]: 'JWT_SECRET' environment variable not set")
	}
	secret = []byte(s)

	// Load identity key
	i, ok := config.GetEnv("JWT_IDENTITY_KEY")
	if !ok {
		log.Fatal("[MIDDLEWARE_ERR]: 'JWT_IDENTITY_KEY' environment variable not set")
	}
	identityKey = i
}
