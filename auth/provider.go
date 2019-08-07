package auth

import (
	"net/http"
)

// Provider defines an auth interface
type Provider interface {
	// Name returns the name of provider
	Name() string

	// Endpoints returns the provider endpoints
	Endpoints() *Endpoints

	// SetEndpoints sets the provider endpoints
	SetEndpoints(endpoints *Endpoints)

	// Session returns a Session
	Session(req *http.Request) (Session, error)
}

// Providers is a map of provider name to Provider
type Providers map[string]Provider

// ProviderFn returns an auth.Provider
type ProviderFn func() Provider
