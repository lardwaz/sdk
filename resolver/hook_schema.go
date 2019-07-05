package resolver

import "github.com/graphql-go/graphql"

type (
	// SchemaHookFn defines the signature for the schema hook function
	SchemaHookFn func(graphql.Fields) error

	// SchemaHooks operates on SchemaHooksFn
	SchemaHooks interface {
		RegisterNewQueryHook(SchemaHookFn)
		RegisterNewMutationHook(SchemaHookFn)

		QueryHooks() []SchemaHookFn
		MutationHooks() []SchemaHookFn
	}

	schemaHooks struct {
		queryhooksFn    []SchemaHookFn
		mutationhooksFn []SchemaHookFn
	}
)

// NewSchemaHooks returns a new SchemaHooks
func NewSchemaHooks() SchemaHooks {
	return &schemaHooks{
		queryhooksFn:    make([]SchemaHookFn, 0),
		mutationhooksFn: make([]SchemaHookFn, 0),
	}
}

// RegisterNewQueryHook
func (s *schemaHooks) RegisterNewQueryHook(hook SchemaHookFn) {
	s.queryhooksFn = append(s.queryhooksFn, hook)
}

// RegisterNewMutationHook
func (s *schemaHooks) RegisterNewMutationHook(hook SchemaHookFn) {
	s.mutationhooksFn = append(s.mutationhooksFn, hook)
}

// QueryHooks() returns a list of SchemaHookFn
func (s schemaHooks) QueryHooks() []SchemaHookFn {
	return s.queryhooksFn
}

// MutationHooks() returns a list of SchemaHookFn
func (s schemaHooks) MutationHooks() []SchemaHookFn {
	return s.mutationhooksFn
}
