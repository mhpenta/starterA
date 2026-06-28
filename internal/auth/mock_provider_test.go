package auth

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestMockProviderVerifiesKnownToken(t *testing.T) {
	provider := NewMockProvider()

	token, err := provider.VerifyIDToken(context.Background(), "test-token-1")
	if err != nil {
		t.Fatalf("verify token: %v", err)
	}
	if token.UID != "user-1" {
		t.Fatalf("UID = %q, want user-1", token.UID)
	}
}

func TestMockProviderRejectsInvalidToken(t *testing.T) {
	provider := NewMockProvider()

	_, err := provider.VerifyIDToken(context.Background(), "missing-token")
	if !errors.Is(err, ErrInvalidToken) {
		t.Fatalf("err = %v, want %v", err, ErrInvalidToken)
	}
}

func TestMockProviderRejectsExpiredToken(t *testing.T) {
	provider := NewMockProvider()
	provider.AddUser("expired-token", &Token{
		UID:      "expired-user",
		Email:    "expired@example.com",
		Expiry:   time.Now().Add(-time.Minute),
		IssuedAt: time.Now().Add(-time.Hour),
	}, nil)

	_, err := provider.VerifyIDToken(context.Background(), "expired-token")
	if !errors.Is(err, ErrExpiredToken) {
		t.Fatalf("err = %v, want %v", err, ErrExpiredToken)
	}
}

func TestMockProviderGetsUserInfo(t *testing.T) {
	provider := NewMockProvider()

	info, err := provider.GetUserInfo(context.Background(), "user-admin")
	if err != nil {
		t.Fatalf("get user info: %v", err)
	}
	if info.Email != "admin@example.com" {
		t.Fatalf("Email = %q, want admin@example.com", info.Email)
	}
}
