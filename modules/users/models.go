package users

// Simple User type
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`

	First string `json:"first"`
	Last  string `json:"last"`
}

// GetUsername returns the user's username
// This method MUST be implemented for JWT authentication to work
func (u User) GetUsername() string {
	return u.Username
}
