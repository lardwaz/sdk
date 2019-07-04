package resolver

import "github.com/graphql-go/graphql"

type (
	// PreSingleHookFn defines the signature of resolver pre hook function
	PreSingleHookFn func(*graphql.ResolveParams) error

	// PostSingleHookFn defines the signature of resolver post hook function
	PostSingleHookFn func(*graphql.ResolveParams, interface{}) (interface{}, error)

	// PreSingleHooksFn defines a map of entityname to resolver hook function
	PreSingleHooksFn map[string][]PreSingleHookFn

	// PostSingleHooksFn defines a map of entityname to resolver hook function
	PostSingleHooksFn map[string][]PostSingleHookFn

	// PreListingHookFn defines the signature of resolver pre hook function
	PreListingHookFn func(*graphql.ResolveParams) error

	// PostListingHookFn defines the signature of resolver post hook function
	PostListingHookFn func(*graphql.ResolveParams, []interface{}) ([]interface{}, error)

	// PreListingHooksFn defines a map of entityname to resolver hook function
	PreListingHooksFn map[string][]PreListingHookFn

	// PostListingHooksFn defines a map of entityname to resolver hook function
	PostListingHooksFn map[string][]PostListingHookFn

	// PreCreateHookFn defines the signature of resolver pre hook function
	PreCreateHookFn func(*graphql.ResolveParams) error

	// PreCreateHooksFn defines a map of entityname to resolver hook function
	PreCreateHooksFn map[string][]PreCreateHookFn

	// PostCreateHookFn defines the signature of resolver post hook function
	PostCreateHookFn func(*graphql.ResolveParams, interface{}) (interface{}, error)

	// PostCreateHooksFn defines a map of entityname to resolver hook function
	PostCreateHooksFn map[string][]PostCreateHookFn

	// PreUpdateHookFn defines the signature of resolver pre hook function
	PreUpdateHookFn func(*graphql.ResolveParams) error

	// PostUpdateHookFn defines the signature of resolver post hook function
	PostUpdateHookFn func(*graphql.ResolveParams, interface{}) (interface{}, error)

	// PreUpdateHooksFn defines a map of entityname to resolver hook function
	PreUpdateHooksFn map[string][]PreUpdateHookFn

	// PostUpdateHooksFn defines a map of entityname to resolver hook function
	PostUpdateHooksFn map[string][]PostUpdateHookFn

	// PreDeleteHookFn defines the signature of resolver pre hook function
	PreDeleteHookFn func(*graphql.ResolveParams) error

	// PostDeleteHookFn defines the signature of resolver post hook function
	PostDeleteHookFn func(*graphql.ResolveParams, interface{}) (interface{}, error)

	// PreDeleteHooksFn defines a map of entityname to resolver hook function
	PreDeleteHooksFn map[string][]PreDeleteHookFn

	// PostDeleteHooksFn defines a map of entityname to resolver hook function
	PostDeleteHooksFn map[string][]PostDeleteHookFn

	// Hooks defines a list of common hooks for the various resolvers
	Hooks interface {
		RegisterNewPreSingleHook(entityname string, hook PreSingleHookFn)
		RegisterNewPostSingleHook(entityname string, hook PostSingleHookFn)
		RegisterNewPreListingHook(entityname string, hook PreListingHookFn)
		RegisterNewPostListingHook(entityname string, hook PostListingHookFn)
		RegisterNewPreCreateHook(entityname string, hook PreCreateHookFn)
		RegisterNewPostCreateHook(entityname string, hook PostCreateHookFn)
		RegisterNewPreUpdateHook(entityname string, hook PreUpdateHookFn)
		RegisterNewPostUpdateHook(entityname string, hook PostUpdateHookFn)
		RegisterNewPreDeleteHook(entityname string, hook PreDeleteHookFn)
		RegisterNewPostDeleteHook(entityname string, hook PostDeleteHookFn)

		PreSingleHooks() PreSingleHooksFn
		PostSingleHooks() PostSingleHooksFn
		PreListingHooks() PreListingHooksFn
		PostListingHooks() PostListingHooksFn
		PreCreateHooks() PreCreateHooksFn
		PostCreateHooks() PostCreateHooksFn
		PreUpdateHooks() PreUpdateHooksFn
		PostUpdateHooks() PostUpdateHooksFn
		PreDeleteHooks() PreDeleteHooksFn
		PostDeleteHooks() PostDeleteHooksFn
	}

	hooks struct {
		preSingleHooksFn   PreSingleHooksFn
		postSingleHooksFn  PostSingleHooksFn
		preListingHooksFn  PreListingHooksFn
		postListingHooksFn PostListingHooksFn
		preCreateHooksFn   PreCreateHooksFn
		postCreateHooksFn  PostCreateHooksFn
		preUpdateHooksFn   PreUpdateHooksFn
		postUpdateHooksFn  PostUpdateHooksFn
		preDeleteHooksFn   PreDeleteHooksFn
		postDeleteHooksFn  PostDeleteHooksFn
	}
)

// NewHooks return a new Hooks
func NewHooks() Hooks {
	return hooks{
		preSingleHooksFn:   make(PreSingleHooksFn),
		postSingleHooksFn:  make(PostSingleHooksFn),
		preListingHooksFn:  make(PreListingHooksFn),
		postListingHooksFn: make(PostListingHooksFn),
		preCreateHooksFn:   make(PreCreateHooksFn),
		postCreateHooksFn:  make(PostCreateHooksFn),
		preUpdateHooksFn:   make(PreUpdateHooksFn),
		postUpdateHooksFn:  make(PostUpdateHooksFn),
		preDeleteHooksFn:   make(PreDeleteHooksFn),
		postDeleteHooksFn:  make(PostDeleteHooksFn),
	}
}

// RegisterNewPreSingleHookrregister a new hook to the PreSingle hook functions
func (h hooks) RegisterNewPreSingleHook(entityname string, hook PreSingleHookFn) {
	h.preSingleHooksFn[entityname] = append(h.preSingleHooksFn[entityname], hook)
}

// RegisterNewPostSingleHook register a new hook to the PostSingle hook functions
func (h hooks) RegisterNewPostSingleHook(entityname string, hook PostSingleHookFn) {
	h.postSingleHooksFn[entityname] = append(h.postSingleHooksFn[entityname], hook)
}

// RegisterNewPreListingHook register a new hook to the PreListing hook functions
func (h hooks) RegisterNewPreListingHook(entityname string, hook PreListingHookFn) {
	h.preListingHooksFn[entityname] = append(h.preListingHooksFn[entityname], hook)
}

// RegisterNewPostListingHookgregister a new hook to the PostListing hook functions
func (h hooks) RegisterNewPostListingHook(entityname string, hook PostListingHookFn) {
	h.postListingHooksFn[entityname] = append(h.postListingHooksFn[entityname], hook)
}

// RegisterNewPreCreateHookrregister a new hook to the PreCreate hook functions
func (h hooks) RegisterNewPreCreateHook(entityname string, hook PreCreateHookFn) {
	h.preCreateHooksFn[entityname] = append(h.preCreateHooksFn[entityname], hook)
}

// RegisterNewPostCreateHook register a new hook to the PostCreate hook functions
func (h hooks) RegisterNewPostCreateHook(entityname string, hook PostCreateHookFn) {
	h.postCreateHooksFn[entityname] = append(h.postCreateHooksFn[entityname], hook)
}

// RegisterNewPreUpdateHookrregister a new hook to the PreUpdate hook functions
func (h hooks) RegisterNewPreUpdateHook(entityname string, hook PreUpdateHookFn) {
	h.preUpdateHooksFn[entityname] = append(h.preUpdateHooksFn[entityname], hook)
}

// RegisterNewPostUpdateHook register a new hook to the PostUpdate hook functions
func (h hooks) RegisterNewPostUpdateHook(entityname string, hook PostUpdateHookFn) {
	h.postUpdateHooksFn[entityname] = append(h.postUpdateHooksFn[entityname], hook)
}

// RegisterNewPreDeleteHookrregister a new hook to the PreDelete hook functions
func (h hooks) RegisterNewPreDeleteHook(entityname string, hook PreDeleteHookFn) {
	h.preDeleteHooksFn[entityname] = append(h.preDeleteHooksFn[entityname], hook)
}

// RegisterNewPostDeleteHook register a new hook to the PostDelete hook functions
func (h hooks) RegisterNewPostDeleteHook(entityname string, hook PostDeleteHookFn) {
	h.postDeleteHooksFn[entityname] = append(h.postDeleteHooksFn[entityname], hook)
}

// PreSingleHooksreturns a list of PreSingle hook functions
func (h hooks) PreSingleHooks() PreSingleHooksFn {
	return h.preSingleHooksFn
}

// PostSingleHooks returns a list of PostSingle hook functions
func (h hooks) PostSingleHooks() PostSingleHooksFn {
	return h.postSingleHooksFn
}

// PreListingHooks returns a list of PreListing hook functions
func (h hooks) PreListingHooks() PreListingHooksFn {
	return h.preListingHooksFn
}

// PostListingHooksg returns a list of PostListing hook functions
func (h hooks) PostListingHooks() PostListingHooksFn {
	return h.postListingHooksFn
}

// PreCreateHooksreturns a list of PreCreate hook functions
func (h hooks) PreCreateHooks() PreCreateHooksFn {
	return h.preCreateHooksFn
}

// PostCreateHooks returns a list of PostCreate hook functions
func (h hooks) PostCreateHooks() PostCreateHooksFn {
	return h.postCreateHooksFn
}

// PreUpdateHooksreturns a list of PreUpdate hook functions
func (h hooks) PreUpdateHooks() PreUpdateHooksFn {
	return h.preUpdateHooksFn
}

// PostUpdateHooks returns a list of PostUpdate hook functions
func (h hooks) PostUpdateHooks() PostUpdateHooksFn {
	return h.postUpdateHooksFn
}

// PreDeleteHooksreturns a list of PreDelete hook functions
func (h hooks) PreDeleteHooks() PreDeleteHooksFn {
	return h.preDeleteHooksFn
}

// PostDeleteHooks returns a list of PostDelete hook functions
func (h hooks) PostDeleteHooks() PostDeleteHooksFn {
	return h.postDeleteHooksFn
}
