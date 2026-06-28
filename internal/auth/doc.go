// Package auth provides optional, provider-agnostic authentication scaffolding.
//
// The package defines a small provider interface, shared token/user types,
// context helpers, HTTP middleware, and a mock provider for development and
// tests. It intentionally avoids app-specific user loading, role checks, and
// global provider setup so templates can wire their chosen auth provider at the
// route boundary.
package auth
