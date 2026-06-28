package auth

import (
	"context"
	"sync"
	"time"
)

// MockProvider is a simple auth provider for development and testing.
// It uses a predefined map of tokens to users.
type MockProvider struct {
	mu     sync.RWMutex
	tokens map[string]*Token    // token string -> Token
	users  map[string]*UserInfo // UID -> UserInfo
}

// NewMockProvider creates a new mock provider with some default test users.
func NewMockProvider() *MockProvider {
	p := &MockProvider{
		tokens: make(map[string]*Token),
		users:  make(map[string]*UserInfo),
	}
	p.addDefaultTestUsers()
	return p
}

// addDefaultTestUsers adds a set of predefined test users.
func (p *MockProvider) addDefaultTestUsers() {
	// Test user 1 - regular user
	p.AddUser("test-token-1", &Token{
		UID:           "user-1",
		Email:         "user1@example.com",
		EmailVerified: true,
		Claims:        map[string]interface{}{"role": "user"},
		Expiry:        time.Now().Add(24 * time.Hour),
		IssuedAt:      time.Now(),
	}, &UserInfo{
		UID:           "user-1",
		Email:         "user1@example.com",
		EmailVerified: true,
		DisplayName:   "Test User One",
		PhotoURL:      "",
	})

	// Test user 2 - admin user
	p.AddUser("test-token-admin", &Token{
		UID:           "user-admin",
		Email:         "admin@example.com",
		EmailVerified: true,
		Claims:        map[string]interface{}{"role": "admin"},
		Expiry:        time.Now().Add(24 * time.Hour),
		IssuedAt:      time.Now(),
	}, &UserInfo{
		UID:           "user-admin",
		Email:         "admin@example.com",
		EmailVerified: true,
		DisplayName:   "Admin User",
		PhotoURL:      "",
	})

	// Test user 3 - unverified user
	p.AddUser("test-token-unverified", &Token{
		UID:           "user-unverified",
		Email:         "unverified@example.com",
		EmailVerified: false,
		Claims:        map[string]interface{}{"role": "user"},
		Expiry:        time.Now().Add(24 * time.Hour),
		IssuedAt:      time.Now(),
	}, &UserInfo{
		UID:           "user-unverified",
		Email:         "unverified@example.com",
		EmailVerified: false,
		DisplayName:   "Unverified User",
		PhotoURL:      "",
	})
}

// AddUser adds a token and user to the mock provider.
func (p *MockProvider) AddUser(token string, t *Token, info *UserInfo) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.tokens[token] = t
	if info != nil {
		p.users[t.UID] = info
	}
}

// RemoveToken removes a token from the mock provider.
func (p *MockProvider) RemoveToken(token string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.tokens, token)
}

// VerifyIDToken validates an ID token and returns the associated Token.
func (p *MockProvider) VerifyIDToken(ctx context.Context, idToken string) (*Token, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	token, ok := p.tokens[idToken]
	if !ok {
		return nil, ErrInvalidToken
	}

	if time.Now().After(token.Expiry) {
		return nil, ErrExpiredToken
	}

	return token, nil
}

// VerifySessionCookie validates a session cookie.
// In the mock, we treat session cookies the same as ID tokens.
func (p *MockProvider) VerifySessionCookie(ctx context.Context, sessionCookie string) (*Token, error) {
	return p.VerifyIDToken(ctx, sessionCookie)
}

// VerifySessionCookieRevoked checks if a session cookie has been revoked.
// In the mock, we just verify the token exists.
func (p *MockProvider) VerifySessionCookieRevoked(ctx context.Context, sessionCookie string) (*Token, error) {
	return p.VerifyIDToken(ctx, sessionCookie)
}

// VerifySessionCookieAndCheckRevoked combines verification and revocation check.
func (p *MockProvider) VerifySessionCookieAndCheckRevoked(ctx context.Context, sessionCookie string) (*Token, error) {
	return p.VerifyIDToken(ctx, sessionCookie)
}

// CreateSessionCookie creates a session cookie from an ID token.
// In the mock, we just return the same token string.
func (p *MockProvider) CreateSessionCookie(ctx context.Context, idToken string, expiresIn time.Duration) (string, error) {
	// Verify the token exists first
	_, err := p.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", err
	}
	// In mock mode, just return the same token
	return idToken, nil
}

// GetUserInfo returns user information for a given UID.
func (p *MockProvider) GetUserInfo(ctx context.Context, uid string) (*UserInfo, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	info, ok := p.users[uid]
	if !ok {
		return nil, ErrUserNotFound
	}
	return info, nil
}

// Ensure MockProvider implements Servicer
var _ Servicer = (*MockProvider)(nil)
