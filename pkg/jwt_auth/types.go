package jwt_auth

import "time"

// User interface contains the most basic functions needed to perform JWT verification
type User interface {
	GetUsername() string
}

// JwtRequest represents the incoming request for a JWT token
type JwtRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JwtResponse represents the response containing the JWT token
type JwtResponse struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
	Code   int       `json:"code"`
}
