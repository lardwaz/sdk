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
		queryhooks    []SchemaHookFn
		mutationhooks []SchemaHookFn
	}
)

// NewSchemaHooks returns a new SchemaHooks
func NewSchemaHooks() SchemaHooks {
	return &schemaHooks{
		queryhooks:    make([]SchemaHookFn, 0),
		mutationhooks: make([]SchemaHookFn, 0),
	}
}

// RegisterNewQueryHook
func (s *schemaHooks) RegisterNewQueryHook(hook SchemaHookFn) {
	s.queryhooks = append(s.queryhooks, hook)
}

// RegisterNewMutationHook
func (s *schemaHooks) RegisterNewMutationHook(hook SchemaHookFn) {
	s.mutationhooks = append(s.mutationhooks, hook)
}

// QueryHooks() returns a list of SchemaHookFn
func (s schemaHooks) QueryHooks() []SchemaHookFn {
	return s.queryhooks
}

// MutationHooks() returns a list of SchemaHookFn
func (s schemaHooks) MutationHooks() []SchemaHookFn {
	return s.mutationhooks
}
