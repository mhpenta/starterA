package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequireAuthWithBearerTokenAddsTokenToContext(t *testing.T) {
	provider := NewMockProvider()
	handler := RequireAuth(provider)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := TokenFromContext(r.Context())
		if !ok {
			t.Fatal("expected token in context")
		}
		if token.UID != "user-1" {
			t.Fatalf("UID = %q, want user-1", token.UID)
		}
		w.WriteHeader(http.StatusNoContent)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(AuthHeader, BearerPrefix+"test-token-1")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusNoContent)
	}
}

func TestRequireAuthWithSessionCookieAddsTokenToContext(t *testing.T) {
	provider := NewMockProvider()
	handler := RequireAuth(provider)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := MustTokenFromContext(r.Context())
		if token.UID != "user-admin" {
			t.Fatalf("UID = %q, want user-admin", token.UID)
		}
		w.WriteHeader(http.StatusNoContent)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{Name: SessionCookieName, Value: "test-token-admin"})
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusNoContent)
	}
}

func TestRequireAuthRejectsMissingToken(t *testing.T) {
	provider := NewMockProvider()
	handler := RequireAuth(provider)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("next handler should not run")
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusUnauthorized)
	}
}

func TestOptionalAuthAllowsInvalidTokenWithoutContext(t *testing.T) {
	provider := NewMockProvider()
	handler := OptionalAuth(provider)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if token, ok := TokenFromContext(r.Context()); ok || token != nil {
			t.Fatalf("expected no token in context, got %#v", token)
		}
		w.WriteHeader(http.StatusNoContent)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(AuthHeader, BearerPrefix+"bad-token")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusNoContent)
	}
}
