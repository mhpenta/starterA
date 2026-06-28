package auth

import (
	"context"
	"time"
)

// Servicer handles authentication operations
type Servicer interface {
	VerifyIDToken(ctx context.Context, idToken string) (*Token, error)
	VerifySessionCookie(ctx context.Context, sessionCookie string) (*Token, error)
	VerifySessionCookieRevoked(ctx context.Context, sessionCookie string) (*Token, error)
	VerifySessionCookieAndCheckRevoked(ctx context.Context, sessionCookie string) (*Token, error)
	CreateSessionCookie(ctx context.Context, idToken string, expiresIn time.Duration) (string, error)
	GetUserInfo(ctx context.Context, uid string) (*UserInfo, error)
}
