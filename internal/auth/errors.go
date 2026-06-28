package auth

import "errors"

// Standard auth errors - providers should return these (or wrap them)
// so middleware can handle them consistently.
var (
	// Token errors
	ErrMissingToken   = errors.New("auth: missing token")
	ErrInvalidToken   = errors.New("auth: invalid token")
	ErrExpiredToken   = errors.New("auth: token expired")
	ErrRevokedToken   = errors.New("auth: token revoked")
	ErrMalformedToken = errors.New("auth: malformed token")

	// Session cookie errors
	ErrMissingCookie        = errors.New("auth: missing session cookie")
	ErrInvalidSessionCookie = errors.New("auth: invalid session cookie")
	ErrExpiredSessionCookie = errors.New("auth: session cookie expired")
	ErrRevokedSessionCookie = errors.New("auth: session cookie revoked")

	// User errors
	ErrUserNotFound = errors.New("auth: user not found")
	ErrUserDisabled = errors.New("auth: user disabled")

	// Context errors
	ErrTokenNotInContext = errors.New("auth: token not found in context")
)

// IsExpiredError returns true if the error indicates an expired token or session.
func IsExpiredError(err error) bool {
	return errors.Is(err, ErrExpiredToken) || errors.Is(err, ErrExpiredSessionCookie)
}

// IsRevokedError returns true if the error indicates a revoked token or session.
func IsRevokedError(err error) bool {
	return errors.Is(err, ErrRevokedToken) || errors.Is(err, ErrRevokedSessionCookie)
}

// IsInvalidError returns true if the error indicates an invalid/malformed token.
func IsInvalidError(err error) bool {
	return errors.Is(err, ErrInvalidToken) ||
		errors.Is(err, ErrMalformedToken) ||
		errors.Is(err, ErrInvalidSessionCookie)
}
