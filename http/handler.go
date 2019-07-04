package http

import "net/http"

type (
	// HandlerFn defines the signature of <any> http handler function
	HandlerFn func(http.ResponseWriter, *http.Request)

	// HandlersFn a map of path to handler function
	HandlersFn map[string]HandlerFn

	// Handlers defines a list of HandlersFn
	Handlers interface {
		RegisterNewHandler(path string, handler HandlerFn)
		Handlers() HandlersFn
	}

	handlers struct {
		handlersFn HandlersFn
	}
)

// NewHandlers returns a new handlers
func NewHandlers() Handlers {
	return handlers{
		handlersFn: make(HandlersFn),
	}
}

// AddHandler adds a new handler
func (h handlers) RegisterNewHandler(path string, handler HandlerFn) {
	h.handlersFn[path] = handler
}

// Handlers return a list of HandlersFn
func (h handlers) Handlers() HandlersFn {
	return h.handlersFn
}
