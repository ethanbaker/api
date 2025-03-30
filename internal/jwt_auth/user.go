package jwt_auth

// user interface contains the most basic functions needed to perform JWT verification
type user interface {
	GetUsername() string
}
