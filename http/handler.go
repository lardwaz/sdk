package http

import "github.com/gorilla/mux"

type (
	// HookFn defines the signature of <any> http hook function
	HookFn func(*mux.Router) error

	// Hooks defines a list of HookFn
	Hooks interface {
		RegisterNewHook(HookFn)
		Hooks() []HookFn
	}

	hooks struct {
		hooksFn []HookFn
	}
)

// NewHooks returns a new Hooks
func NewHooks() Hooks {
	return &hooks{
		hooksFn: make([]HookFn, 0),
	}
}

// RegisterNewHook registers a new hook function
func (h *hooks) RegisterNewHook(hook HookFn) {
	h.hooksFn = append(h.hooksFn, hook)
}

// Hooks return a list of hookFn
func (h hooks) Hooks() []HookFn {
	return h.hooksFn
}
