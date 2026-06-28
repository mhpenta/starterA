package auth

import (
	"net/http"
	"strings"
)

// TokenSource indicates where the token was found.
type TokenSource int

const (
	TokenSourceNone TokenSource = iota
	TokenSourceHeader
	TokenSourceCookie
)

// extractToken attempts to get a token from the request.
// It checks the Authorization header first, then falls back to the session cookie.
func extractToken(r *http.Request) (token string, source TokenSource) {
	// Check Authorization header first
	authHeader := r.Header.Get(AuthHeader)
	if strings.HasPrefix(authHeader, BearerPrefix) {
		return strings.TrimPrefix(authHeader, BearerPrefix), TokenSourceHeader
	}

	// Fall back to session cookie
	cookie, err := r.Cookie(SessionCookieName)
	if err == nil && cookie.Value != "" {
		return cookie.Value, TokenSourceCookie
	}

	return "", TokenSourceNone
}

// RequireAuth returns middleware that requires a valid auth token.
// Requests without a valid token receive a 401 Unauthorized response.
func RequireAuth(provider Servicer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr, source := extractToken(r)
			if tokenStr == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			var token *Token
			var err error

			// Verify based on source
			switch source {
			case TokenSourceHeader:
				token, err = provider.VerifyIDToken(r.Context(), tokenStr)
			case TokenSourceCookie:
				token, err = provider.VerifySessionCookie(r.Context(), tokenStr)
			}

			if err != nil {
				// Could log the error here for debugging
				// But don't leak error details to client
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Add token to context and continue
			ctx := ContextWithToken(r.Context(), token)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// OptionalAuth returns middleware that validates a token if present,
// but allows the request to continue even without authentication.
// If a token is present and valid, it's added to the context.
// If no token or invalid token, the request continues without a token in context.
func OptionalAuth(provider Servicer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr, source := extractToken(r)
			if tokenStr == "" {
				// No token, continue without auth
				next.ServeHTTP(w, r)
				return
			}

			var token *Token
			var err error

			switch source {
			case TokenSourceHeader:
				token, err = provider.VerifyIDToken(r.Context(), tokenStr)
			case TokenSourceCookie:
				token, err = provider.VerifySessionCookie(r.Context(), tokenStr)
			}

			if err != nil {
				// Invalid token, but optional auth - continue without it
				next.ServeHTTP(w, r)
				return
			}

			// Valid token - add to context
			ctx := ContextWithToken(r.Context(), token)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// RequireAuthWithRevocationCheck is like RequireAuth but also checks if the
// session has been revoked. Use this for sensitive operations.
func RequireAuthWithRevocationCheck(provider Servicer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr, source := extractToken(r)
			if tokenStr == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			var token *Token
			var err error

			switch source {
			case TokenSourceHeader:
				token, err = provider.VerifyIDToken(r.Context(), tokenStr)
			case TokenSourceCookie:
				// Use the revocation-checking method for cookies
				token, err = provider.VerifySessionCookieAndCheckRevoked(r.Context(), tokenStr)
			}

			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := ContextWithToken(r.Context(), token)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
