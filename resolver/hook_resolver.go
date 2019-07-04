package resolver

import "github.com/graphql-go/graphql"

type (
	// PreHookFn defines the signature of <any> resolver pre hook function
	PreHookFn func(*graphql.ResolveParams) error

	// PostHookFn defines the signature of <any> resolver post hook function
	PostHookFn func(*graphql.ResolveParams, interface{}) (interface{}, error)

	// PreHooksFn defines a map of entityname to resolver hook function
	PreHooksFn map[string][]PreHookFn

	// PostHooksFn defines a map of entityname to resolver hook function
	PostHooksFn map[string][]PostHookFn

	// Hooks defines a list of common hooks for the various resolvers
	Hooks interface {
		RegisterNewPreSingleHook(entityname string, hook PreHookFn)
		RegisterNewPostSingleHook(entityname string, hook PostHookFn)
		RegisterNewPreListingHook(entityname string, hook PreHookFn)
		RegisterNewPostListingHook(entityname string, hook PostHookFn)
		RegisterNewPreCreateHook(entityname string, hook PreHookFn)
		RegisterNewPostCreateHook(entityname string, hook PostHookFn)
		RegisterNewPreUpdateHook(entityname string, hook PreHookFn)
		RegisterNewPostUpdateHook(entityname string, hook PostHookFn)
		RegisterNewPreDeleteHook(entityname string, hook PreHookFn)
		RegisterNewPostDeleteHook(entityname string, hook PostHookFn)

		PreSingleHooks() PreHooksFn
		PostSingleHooks() PostHooksFn
		PreListingHooks() PreHooksFn
		PostListingHooks() PostHooksFn
		PreCreateHooks() PreHooksFn
		PostCreateHooks() PostHooksFn
		PreUpdateHooks() PreHooksFn
		PostUpdateHooks() PostHooksFn
		PreDeleteHooks() PreHooksFn
		PostDeleteHooks() PostHooksFn
	}

	hooks struct {
		preSingleHooksFn   PreHooksFn
		postSingleHooksFn  PostHooksFn
		preListingHooksFn  PreHooksFn
		postListingHooksFn PostHooksFn
		preCreateHooksFn   PreHooksFn
		postCreateHooksFn  PostHooksFn
		preUpdateHooksFn   PreHooksFn
		postUpdateHooksFn  PostHooksFn
		preDeleteHooksFn   PreHooksFn
		postDeleteHooksFn  PostHooksFn
	}
)

// NewHooks return a new Hooks
func NewHooks() Hooks {
	return hooks{
		preSingleHooksFn:   make(PreHooksFn),
		postSingleHooksFn:  make(PostHooksFn),
		preListingHooksFn:  make(PreHooksFn),
		postListingHooksFn: make(PostHooksFn),
		preCreateHooksFn:   make(PreHooksFn),
		postCreateHooksFn:  make(PostHooksFn),
		preUpdateHooksFn:   make(PreHooksFn),
		postUpdateHooksFn:  make(PostHooksFn),
		preDeleteHooksFn:   make(PreHooksFn),
		postDeleteHooksFn:  make(PostHooksFn),
	}
}

// RegisterNewPreSingleHookrregister a new hook to the PreSingle hook functions
func (h hooks) RegisterNewPreSingleHook(entityname string, hook PreHookFn) {
	h.preSingleHooksFn[entityname] = append(h.preSingleHooksFn[entityname], hook)
}

// RegisterNewPostSingleHook register a new hook to the PostSingle hook functions
func (h hooks) RegisterNewPostSingleHook(entityname string, hook PostHookFn) {
	h.postSingleHooksFn[entityname] = append(h.postSingleHooksFn[entityname], hook)
}

// RegisterNewPreListingHook register a new hook to the PreListing hook functions
func (h hooks) RegisterNewPreListingHook(entityname string, hook PreHookFn) {
	h.preListingHooksFn[entityname] = append(h.preListingHooksFn[entityname], hook)
}

// RegisterNewPostListingHookgregister a new hook to the PostListing hook functions
func (h hooks) RegisterNewPostListingHook(entityname string, hook PostHookFn) {
	h.postListingHooksFn[entityname] = append(h.postListingHooksFn[entityname], hook)
}

// RegisterNewPreCreateHookrregister a new hook to the PreCreate hook functions
func (h hooks) RegisterNewPreCreateHook(entityname string, hook PreHookFn) {
	h.preCreateHooksFn[entityname] = append(h.preCreateHooksFn[entityname], hook)
}

// RegisterNewPostCreateHook register a new hook to the PostCreate hook functions
func (h hooks) RegisterNewPostCreateHook(entityname string, hook PostHookFn) {
	h.postCreateHooksFn[entityname] = append(h.postCreateHooksFn[entityname], hook)
}

// RegisterNewPreUpdateHookrregister a new hook to the PreUpdate hook functions
func (h hooks) RegisterNewPreUpdateHook(entityname string, hook PreHookFn) {
	h.preUpdateHooksFn[entityname] = append(h.preUpdateHooksFn[entityname], hook)
}

// RegisterNewPostUpdateHook register a new hook to the PostUpdate hook functions
func (h hooks) RegisterNewPostUpdateHook(entityname string, hook PostHookFn) {
	h.postUpdateHooksFn[entityname] = append(h.postUpdateHooksFn[entityname], hook)
}

// RegisterNewPreDeleteHookrregister a new hook to the PreDelete hook functions
func (h hooks) RegisterNewPreDeleteHook(entityname string, hook PreHookFn) {
	h.preDeleteHooksFn[entityname] = append(h.preDeleteHooksFn[entityname], hook)
}

// RegisterNewPostDeleteHook register a new hook to the PostDelete hook functions
func (h hooks) RegisterNewPostDeleteHook(entityname string, hook PostHookFn) {
	h.postDeleteHooksFn[entityname] = append(h.postDeleteHooksFn[entityname], hook)
}

// PreSingleHooksreturns a list of PreSingle hook functions
func (h hooks) PreSingleHooks() PreHooksFn {
	return h.preSingleHooksFn
}

// PostSingleHooks returns a list of PostSingle hook functions
func (h hooks) PostSingleHooks() PostHooksFn {
	return h.postSingleHooksFn
}

// PreListingHooks returns a list of PreListing hook functions
func (h hooks) PreListingHooks() PreHooksFn {
	return h.preListingHooksFn
}

// PostListingHooksg returns a list of PostListing hook functions
func (h hooks) PostListingHooks() PostHooksFn {
	return h.postListingHooksFn
}

// PreCreateHooksreturns a list of PreCreate hook functions
func (h hooks) PreCreateHooks() PreHooksFn {
	return h.preCreateHooksFn
}

// PostCreateHooks returns a list of PostCreate hook functions
func (h hooks) PostCreateHooks() PostHooksFn {
	return h.postCreateHooksFn
}

// PreUpdateHooksreturns a list of PreUpdate hook functions
func (h hooks) PreUpdateHooks() PreHooksFn {
	return h.preUpdateHooksFn
}

// PostUpdateHooks returns a list of PostUpdate hook functions
func (h hooks) PostUpdateHooks() PostHooksFn {
	return h.postUpdateHooksFn
}

// PreDeleteHooksreturns a list of PreDelete hook functions
func (h hooks) PreDeleteHooks() PreHooksFn {
	return h.preDeleteHooksFn
}

// PostDeleteHooks returns a list of PostDelete hook functions
func (h hooks) PostDeleteHooks() PostHooksFn {
	return h.postDeleteHooksFn
}
