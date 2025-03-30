package jwt_auth

// JwtRequest represents the incoming request for a JWT
type JwtRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
