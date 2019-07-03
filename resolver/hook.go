package resolver

import "github.com/graphql-go/graphql"

type (
	// HookFn defines the signature of <any> resolver hook function
	HookFn func(p graphql.ResolveParams, vals interface{}) (bool, error)

	// HooksFn defines a map of entityname to resolver hook function
	HooksFn map[string]HookFn

	// Hooks defines a list of common hooks for the various resolvers
	Hooks interface {
		RegisterNewPreSingleHook(entityname string, hook HookFn)
		RegisterNewPostSingleHook(entityname string, hook HookFn)
		RegisterNewPreListingHook(entityname string, hook HookFn)
		RegisterNewPostListingHook(entityname string, hook HookFn)
		RegisterNewPreCreateHook(entityname string, hook HookFn)
		RegisterNewPostCreateHook(entityname string, hook HookFn)
		RegisterNewPreUpdateHook(entityname string, hook HookFn)
		RegisterNewPostUpdateHook(entityname string, hook HookFn)
		RegisterNewPreDeleteHook(entityname string, hook HookFn)
		RegisterNewPostDeleteHook(entityname string, hook HookFn)

		PreSingleHooks() HooksFn
		PostSingleHooks() HooksFn
		PreListingHooks() HooksFn
		PostListingHooks() HooksFn
		PreCreateHooks() HooksFn
		PostCreateHooks() HooksFn
		PreUpdateHooks() HooksFn
		PostUpdateHooks() HooksFn
		PreDeleteHooks() HooksFn
		PostDeleteHooks() HooksFn
	}

	hook struct {
		preSingleHooksFn   HooksFn
		postSingleHooksFn  HooksFn
		preListingHooksFn  HooksFn
		postListingHooksFn HooksFn
		preCreateHooksFn   HooksFn
		postCreateHooksFn  HooksFn
		preUpdateHooksFn   HooksFn
		postUpdateHooksFn  HooksFn
		preDeleteHooksFn   HooksFn
		postDeleteHooksFn  HooksFn
	}
)

// NewHooks return a new Hooks
func NewHooks() Hooks {
	return hook{
		preSingleHooksFn:   make(HooksFn),
		postSingleHooksFn:  make(HooksFn),
		preListingHooksFn:  make(HooksFn),
		postListingHooksFn: make(HooksFn),
		preCreateHooksFn:   make(HooksFn),
		postCreateHooksFn:  make(HooksFn),
		preUpdateHooksFn:   make(HooksFn),
		postUpdateHooksFn:  make(HooksFn),
		preDeleteHooksFn:   make(HooksFn),
		postDeleteHooksFn:  make(HooksFn),
	}
}

// RegisterNewPreSingleHookrregister a new hook to the PreSingle hook functions
func (h hook) RegisterNewPreSingleHook(entityname string, hook HookFn) {
	h.preSingleHooksFn[entityname] = hook
}

// RegisterNewPostSingleHook register a new hook to the PostSingle hook functions
func (h hook) RegisterNewPostSingleHook(entityname string, hook HookFn) {
	h.postSingleHooksFn[entityname] = hook
}

// RegisterNewPreListingHook register a new hook to the PreListing hook functions
func (h hook) RegisterNewPreListingHook(entityname string, hook HookFn) {
	h.preListingHooksFn[entityname] = hook
}

// RegisterNewPostListingHookgregister a new hook to the PostListing hook functions
func (h hook) RegisterNewPostListingHook(entityname string, hook HookFn) {
	h.postListingHooksFn[entityname] = hook
}

// RegisterNewPreCreateHookrregister a new hook to the PreCreate hook functions
func (h hook) RegisterNewPreCreateHook(entityname string, hook HookFn) {
	h.preCreateHooksFn[entityname] = hook
}

// RegisterNewPostCreateHook register a new hook to the PostCreate hook functions
func (h hook) RegisterNewPostCreateHook(entityname string, hook HookFn) {
	h.postCreateHooksFn[entityname] = hook
}

// RegisterNewPreUpdateHookrregister a new hook to the PreUpdate hook functions
func (h hook) RegisterNewPreUpdateHook(entityname string, hook HookFn) {
	h.preUpdateHooksFn[entityname] = hook
}

// RegisterNewPostUpdateHook register a new hook to the PostUpdate hook functions
func (h hook) RegisterNewPostUpdateHook(entityname string, hook HookFn) {
	h.postUpdateHooksFn[entityname] = hook
}

// RegisterNewPreDeleteHookrregister a new hook to the PreDelete hook functions
func (h hook) RegisterNewPreDeleteHook(entityname string, hook HookFn) {
	h.preDeleteHooksFn[entityname] = hook
}

// RegisterNewPostDeleteHook register a new hook to the PostDelete hook functions
func (h hook) RegisterNewPostDeleteHook(entityname string, hook HookFn) {
	h.postDeleteHooksFn[entityname] = hook
}

// PreSingleHooksreturns a list of PreSingle hook functions
func (h hook) PreSingleHooks() HooksFn {
	return h.preSingleHooksFn
}

// PostSingleHooks returns a list of PostSingle hook functions
func (h hook) PostSingleHooks() HooksFn {
	return h.postSingleHooksFn
}

// PreListingHooks returns a list of PreListing hook functions
func (h hook) PreListingHooks() HooksFn {
	return h.preListingHooksFn
}

// PostListingHooksg returns a list of PostListing hook functions
func (h hook) PostListingHooks() HooksFn {
	return h.postListingHooksFn
}

// PreCreateHooksreturns a list of PreCreate hook functions
func (h hook) PreCreateHooks() HooksFn {
	return h.preCreateHooksFn
}

// PostCreateHooks returns a list of PostCreate hook functions
func (h hook) PostCreateHooks() HooksFn {
	return h.postCreateHooksFn
}

// PreUpdateHooksreturns a list of PreUpdate hook functions
func (h hook) PreUpdateHooks() HooksFn {
	return h.preUpdateHooksFn
}

// PostUpdateHooks returns a list of PostUpdate hook functions
func (h hook) PostUpdateHooks() HooksFn {
	return h.postUpdateHooksFn
}

// PreDeleteHooksreturns a list of PreDelete hook functions
func (h hook) PreDeleteHooks() HooksFn {
	return h.preDeleteHooksFn
}

// PostDeleteHooks returns a list of PostDelete hook functions
func (h hook) PostDeleteHooks() HooksFn {
	return h.postDeleteHooksFn
}
