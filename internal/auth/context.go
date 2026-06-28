package auth

import "context"

// TokenFromContext extracts the auth token from context.
// Returns nil and false if no token is present.
func TokenFromContext(ctx context.Context) (*Token, bool) {
	token, ok := ctx.Value(authTokenKey).(*Token)
	return token, ok
}

// ExtractAuthTokenFromContext extracts the auth token from context.
// It mirrors the longer helper name used in larger apps.
func ExtractAuthTokenFromContext(ctx context.Context) (*Token, bool) {
	return TokenFromContext(ctx)
}

// MustTokenFromContext extracts the auth token from context.
// Panics if no token is present - use only in routes protected by RequireAuth.
func MustTokenFromContext(ctx context.Context) *Token {
	token, ok := TokenFromContext(ctx)
	if !ok || token == nil {
		panic("auth: MustTokenFromContext called without token in context - is RequireAuth middleware applied?")
	}
	return token
}

// ContextWithToken returns a new context with the token added.
func ContextWithToken(ctx context.Context, token *Token) context.Context {
	return context.WithValue(ctx, authTokenKey, token)
}

// ContextWithAuthToken returns a new context with the token added.
// It mirrors the longer helper name used in larger apps.
func ContextWithAuthToken(ctx context.Context, token *Token) context.Context {
	return ContextWithToken(ctx, token)
}

// UIDFromContext is a convenience function to get just the user ID.
// Returns empty string if no token is present.
func UIDFromContext(ctx context.Context) string {
	token, ok := TokenFromContext(ctx)
	if !ok || token == nil {
		return ""
	}
	return token.UID
}
