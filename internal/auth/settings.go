package auth

const (
	// SessionCookieName is the default name for the session cookie.
	SessionCookieName = "session"

	// AuthHeader is the header used for bearer token authentication.
	AuthHeader = "Authorization"

	// BearerPrefix is the prefix for bearer tokens in the Authorization header.
	BearerPrefix = "Bearer "
)
