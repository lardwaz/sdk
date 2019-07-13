package hook

import "github.com/graphql-go/graphql"

type (
	// GlobalPreSingleHookFn defines the signature of resolver pre hook function
	GlobalPreSingleHookFn func(*graphql.ResolveParams) error

	// GlobalPostSingleHookFn defines the signature of resolver post hook function
	GlobalPostSingleHookFn func(*graphql.ResolveParams, interface{}) (interface{}, error)

	// GlobalPreSingleHooksFn defines a map of entityname to resolver hook function
	GlobalPreSingleHooksFn []GlobalPreSingleHookFn

	// GlobalPostSingleHooksFn defines a map of entityname to resolver hook function
	GlobalPostSingleHooksFn []GlobalPostSingleHookFn

	// GlobalPreListingHookFn defines the signature of resolver pre hook function
	GlobalPreListingHookFn func(*graphql.ResolveParams) error

	// GlobalPostListingHookFn defines the signature of resolver post hook function
	GlobalPostListingHookFn func(*graphql.ResolveParams, []interface{}) ([]interface{}, error)

	// GlobalPreListingHooksFn defines a map of entityname to resolver hook function
	GlobalPreListingHooksFn []GlobalPreListingHookFn

	// GlobalPostListingHooksFn defines a map of entityname to resolver hook function
	GlobalPostListingHooksFn []GlobalPostListingHookFn

	// GlobalPreCreateHookFn defines the signature of resolver pre hook function
	GlobalPreCreateHookFn func(*graphql.ResolveParams) error

	// GlobalPreCreateHooksFn defines a map of entityname to resolver hook function
	GlobalPreCreateHooksFn []GlobalPreCreateHookFn

	// GlobalPostCreateHookFn defines the signature of resolver post hook function
	GlobalPostCreateHookFn func(*graphql.ResolveParams, interface{}) (interface{}, error)

	// GlobalPostCreateHooksFn defines a map of entityname to resolver hook function
	GlobalPostCreateHooksFn []GlobalPostCreateHookFn

	// GlobalPreUpdateHookFn defines the signature of resolver pre hook function
	GlobalPreUpdateHookFn func(*graphql.ResolveParams) error

	// GlobalPostUpdateHookFn defines the signature of resolver post hook function
	GlobalPostUpdateHookFn func(*graphql.ResolveParams, interface{}) (interface{}, error)

	// GlobalPreUpdateHooksFn defines a map of entityname to resolver hook function
	GlobalPreUpdateHooksFn []GlobalPreUpdateHookFn

	// GlobalPostUpdateHooksFn defines a map of entityname to resolver hook function
	GlobalPostUpdateHooksFn []GlobalPostUpdateHookFn

	// GlobalPreDeleteHookFn defines the signature of resolver pre hook function
	GlobalPreDeleteHookFn func(*graphql.ResolveParams) error

	// GlobalPostDeleteHookFn defines the signature of resolver post hook function
	GlobalPostDeleteHookFn func(*graphql.ResolveParams, interface{}) (interface{}, error)

	// GlobalPreDeleteHooksFn defines a map of entityname to resolver hook function
	GlobalPreDeleteHooksFn []GlobalPreDeleteHookFn

	// GlobalPostDeleteHooksFn defines a map of entityname to resolver hook function
	GlobalPostDeleteHooksFn []GlobalPostDeleteHookFn

	// GlobalResolverHooks defines a list of common hooks for the various resolvers
	GlobalResolverHooks interface {
		RegisterNewGlobalPreSingleHook(hook GlobalPreSingleHookFn)
		RegisterNewGlobalPostSingleHook(hook GlobalPostSingleHookFn)
		RegisterNewGlobalPreListingHook(hook GlobalPreListingHookFn)
		RegisterNewGlobalPostListingHook(hook GlobalPostListingHookFn)
		RegisterNewGlobalPreCreateHook(hook GlobalPreCreateHookFn)
		RegisterNewGlobalPostCreateHook(hook GlobalPostCreateHookFn)
		RegisterNewGlobalPreUpdateHook(hook GlobalPreUpdateHookFn)
		RegisterNewGlobalPostUpdateHook(hook GlobalPostUpdateHookFn)
		RegisterNewGlobalPreDeleteHook(hook GlobalPreDeleteHookFn)
		RegisterNewGlobalPostDeleteHook(hook GlobalPostDeleteHookFn)

		GlobalPreSingleHooks() GlobalPreSingleHooksFn
		GlobalPostSingleHooks() GlobalPostSingleHooksFn
		GlobalPreListingHooks() GlobalPreListingHooksFn
		GlobalPostListingHooks() GlobalPostListingHooksFn
		GlobalPreCreateHooks() GlobalPreCreateHooksFn
		GlobalPostCreateHooks() GlobalPostCreateHooksFn
		GlobalPreUpdateHooks() GlobalPreUpdateHooksFn
		GlobalPostUpdateHooks() GlobalPostUpdateHooksFn
		GlobalPreDeleteHooks() GlobalPreDeleteHooksFn
		GlobalPostDeleteHooks() GlobalPostDeleteHooksFn
	}

	globalResolverHooks struct {
		globalPreSingleHooksFn   GlobalPreSingleHooksFn
		globalPostSingleHooksFn  GlobalPostSingleHooksFn
		globalPreListingHooksFn  GlobalPreListingHooksFn
		globalPostListingHooksFn GlobalPostListingHooksFn
		globalPreCreateHooksFn   GlobalPreCreateHooksFn
		globalPostCreateHooksFn  GlobalPostCreateHooksFn
		globalPreUpdateHooksFn   GlobalPreUpdateHooksFn
		globalPostUpdateHooksFn  GlobalPostUpdateHooksFn
		globalPreDeleteHooksFn   GlobalPreDeleteHooksFn
		globalPostDeleteHooksFn  GlobalPostDeleteHooksFn
	}
)

// NewGlobalResolverHooks return a new GlobalResolverHooks
func NewGlobalResolverHooks() GlobalResolverHooks {
	return globalResolverHooks{
		globalPreSingleHooksFn:   make(GlobalPreSingleHooksFn, 0),
		globalPostSingleHooksFn:  make(GlobalPostSingleHooksFn, 0),
		globalPreListingHooksFn:  make(GlobalPreListingHooksFn, 0),
		globalPostListingHooksFn: make(GlobalPostListingHooksFn, 0),
		globalPreCreateHooksFn:   make(GlobalPreCreateHooksFn, 0),
		globalPostCreateHooksFn:  make(GlobalPostCreateHooksFn, 0),
		globalPreUpdateHooksFn:   make(GlobalPreUpdateHooksFn, 0),
		globalPostUpdateHooksFn:  make(GlobalPostUpdateHooksFn, 0),
		globalPreDeleteHooksFn:   make(GlobalPreDeleteHooksFn, 0),
		globalPostDeleteHooksFn:  make(GlobalPostDeleteHooksFn, 0),
	}
}

// RegisterNewGlobalPreSingleHookrregister a new hook to the PreSingle hook functions
func (r globalResolverHooks) RegisterNewGlobalPreSingleHook(hook GlobalPreSingleHookFn) {
	r.globalPreSingleHooksFn = append(r.globalPreSingleHooksFn, hook)
}

// RegisterNewGlobalPostSingleHook register a new hook to the PostSingle hook functions
func (r globalResolverHooks) RegisterNewGlobalPostSingleHook(hook GlobalPostSingleHookFn) {
	r.globalPostSingleHooksFn = append(r.globalPostSingleHooksFn, hook)
}

// RegisterNewGlobalPreListingHook register a new hook to the PreListing hook functions
func (r globalResolverHooks) RegisterNewGlobalPreListingHook(hook GlobalPreListingHookFn) {
	r.globalPreListingHooksFn = append(r.globalPreListingHooksFn, hook)
}

// RegisterNewGlobalPostListingHookgregister a new hook to the PostListing hook functions
func (r globalResolverHooks) RegisterNewGlobalPostListingHook(hook GlobalPostListingHookFn) {
	r.globalPostListingHooksFn = append(r.globalPostListingHooksFn, hook)
}

// RegisterNewGlobalPreCreateHookrregister a new hook to the PreCreate hook functions
func (r globalResolverHooks) RegisterNewGlobalPreCreateHook(hook GlobalPreCreateHookFn) {
	r.globalPreCreateHooksFn = append(r.globalPreCreateHooksFn, hook)
}

// RegisterNewGlobalPostCreateHook register a new hook to the PostCreate hook functions
func (r globalResolverHooks) RegisterNewGlobalPostCreateHook(hook GlobalPostCreateHookFn) {
	r.globalPostCreateHooksFn = append(r.globalPostCreateHooksFn, hook)
}

// RegisterNewGlobalPreUpdateHookrregister a new hook to the PreUpdate hook functions
func (r globalResolverHooks) RegisterNewGlobalPreUpdateHook(hook GlobalPreUpdateHookFn) {
	r.globalPreUpdateHooksFn = append(r.globalPreUpdateHooksFn, hook)
}

// RegisterNewGlobalPostUpdateHook register a new hook to the PostUpdate hook functions
func (r globalResolverHooks) RegisterNewGlobalPostUpdateHook(hook GlobalPostUpdateHookFn) {
	r.globalPostUpdateHooksFn = append(r.globalPostUpdateHooksFn, hook)
}

// RegisterNewGlobalPreDeleteHookrregister a new hook to the PreDelete hook functions
func (r globalResolverHooks) RegisterNewGlobalPreDeleteHook(hook GlobalPreDeleteHookFn) {
	r.globalPreDeleteHooksFn = append(r.globalPreDeleteHooksFn, hook)
}

// RegisterNewGlobalPostDeleteHook register a new hook to the PostDelete hook functions
func (r globalResolverHooks) RegisterNewGlobalPostDeleteHook(hook GlobalPostDeleteHookFn) {
	r.globalPostDeleteHooksFn = append(r.globalPostDeleteHooksFn, hook)
}

// GlobalPreSingleHooksreturns a list of PreSingle hook functions
func (r globalResolverHooks) GlobalPreSingleHooks() GlobalPreSingleHooksFn {
	return r.globalPreSingleHooksFn
}

// GlobalPostSingleHooks returns a list of PostSingle hook functions
func (r globalResolverHooks) GlobalPostSingleHooks() GlobalPostSingleHooksFn {
	return r.globalPostSingleHooksFn
}

// GlobalPreListingHooks returns a list of PreListing hook functions
func (r globalResolverHooks) GlobalPreListingHooks() GlobalPreListingHooksFn {
	return r.globalPreListingHooksFn
}

// GlobalPostListingHooksg returns a list of PostListing hook functions
func (r globalResolverHooks) GlobalPostListingHooks() GlobalPostListingHooksFn {
	return r.globalPostListingHooksFn
}

// GlobalPreCreateHooksreturns a list of PreCreate hook functions
func (r globalResolverHooks) GlobalPreCreateHooks() GlobalPreCreateHooksFn {
	return r.globalPreCreateHooksFn
}

// GlobalPostCreateHooks returns a list of PostCreate hook functions
func (r globalResolverHooks) GlobalPostCreateHooks() GlobalPostCreateHooksFn {
	return r.globalPostCreateHooksFn
}

// GlobalPreUpdateHooksreturns a list of PreUpdate hook functions
func (r globalResolverHooks) GlobalPreUpdateHooks() GlobalPreUpdateHooksFn {
	return r.globalPreUpdateHooksFn
}

// GlobalPostUpdateHooks returns a list of PostUpdate hook functions
func (r globalResolverHooks) GlobalPostUpdateHooks() GlobalPostUpdateHooksFn {
	return r.globalPostUpdateHooksFn
}

// GlobalPreDeleteHooksreturns a list of PreDelete hook functions
func (r globalResolverHooks) GlobalPreDeleteHooks() GlobalPreDeleteHooksFn {
	return r.globalPreDeleteHooksFn
}

// GlobalPostDeleteHooks returns a list of PostDelete hook functions
func (r globalResolverHooks) GlobalPostDeleteHooks() GlobalPostDeleteHooksFn {
	return r.globalPostDeleteHooksFn
}
