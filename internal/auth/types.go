package auth

import "time"

type contextKey string

const (
	// authTokenKey is the context key for storing the authenticated token.
	authTokenKey contextKey = "auth-token"
)

// Token represents a validated token.
type Token struct {
	UID           string
	Email         string
	EmailVerified bool
	Claims        map[string]interface{}
	Expiry        time.Time
	IssuedAt      time.Time
}

// UserInfo contains basic user information for initial user creation.
type UserInfo struct {
	UID           string
	Email         string
	EmailVerified bool
	DisplayName   string
	PhotoURL      string
}
