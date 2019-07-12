package hook

import "github.com/gorilla/mux"

type (
	// HTTPHookFn defines the signature of <any> http hook function
	HTTPHookFn func(*mux.Router) error

	// HTTPHooks defines a list of HookFn
	HTTPHooks interface {
		RegisterNewHook(HTTPHookFn)
		Hooks() []HTTPHookFn
	}

	httpHooks struct {
		hooksFn []HTTPHookFn
	}
)

// NewHTTPHooks returns a new HTTPHooks
func NewHTTPHooks() HTTPHooks {
	return &httpHooks{
		hooksFn: make([]HTTPHookFn, 0),
	}
}

// RegisterNewHook registers a new hook function
func (h *httpHooks) RegisterNewHook(hook HTTPHookFn) {
	h.hooksFn = append(h.hooksFn, hook)
}

// Hooks return a list of hook function
func (h httpHooks) Hooks() []HTTPHookFn {
	return h.hooksFn
}
