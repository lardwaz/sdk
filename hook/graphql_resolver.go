package hook

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

	// ResolverHooks defines a list of common hooks for the various resolvers
	ResolverHooks interface {
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

	resolverHooks struct {
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

// NewResolverHooks return a new ResolverHooks
func NewResolverHooks() ResolverHooks {
	return resolverHooks{
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
func (r resolverHooks) RegisterNewPreSingleHook(entityname string, hook PreSingleHookFn) {
	r.preSingleHooksFn[entityname] = append(r.preSingleHooksFn[entityname], hook)
}

// RegisterNewPostSingleHook register a new hook to the PostSingle hook functions
func (r resolverHooks) RegisterNewPostSingleHook(entityname string, hook PostSingleHookFn) {
	r.postSingleHooksFn[entityname] = append(r.postSingleHooksFn[entityname], hook)
}

// RegisterNewPreListingHook register a new hook to the PreListing hook functions
func (r resolverHooks) RegisterNewPreListingHook(entityname string, hook PreListingHookFn) {
	r.preListingHooksFn[entityname] = append(r.preListingHooksFn[entityname], hook)
}

// RegisterNewPostListingHookgregister a new hook to the PostListing hook functions
func (r resolverHooks) RegisterNewPostListingHook(entityname string, hook PostListingHookFn) {
	r.postListingHooksFn[entityname] = append(r.postListingHooksFn[entityname], hook)
}

// RegisterNewPreCreateHookrregister a new hook to the PreCreate hook functions
func (r resolverHooks) RegisterNewPreCreateHook(entityname string, hook PreCreateHookFn) {
	r.preCreateHooksFn[entityname] = append(r.preCreateHooksFn[entityname], hook)
}

// RegisterNewPostCreateHook register a new hook to the PostCreate hook functions
func (r resolverHooks) RegisterNewPostCreateHook(entityname string, hook PostCreateHookFn) {
	r.postCreateHooksFn[entityname] = append(r.postCreateHooksFn[entityname], hook)
}

// RegisterNewPreUpdateHookrregister a new hook to the PreUpdate hook functions
func (r resolverHooks) RegisterNewPreUpdateHook(entityname string, hook PreUpdateHookFn) {
	r.preUpdateHooksFn[entityname] = append(r.preUpdateHooksFn[entityname], hook)
}

// RegisterNewPostUpdateHook register a new hook to the PostUpdate hook functions
func (r resolverHooks) RegisterNewPostUpdateHook(entityname string, hook PostUpdateHookFn) {
	r.postUpdateHooksFn[entityname] = append(r.postUpdateHooksFn[entityname], hook)
}

// RegisterNewPreDeleteHookrregister a new hook to the PreDelete hook functions
func (r resolverHooks) RegisterNewPreDeleteHook(entityname string, hook PreDeleteHookFn) {
	r.preDeleteHooksFn[entityname] = append(r.preDeleteHooksFn[entityname], hook)
}

// RegisterNewPostDeleteHook register a new hook to the PostDelete hook functions
func (r resolverHooks) RegisterNewPostDeleteHook(entityname string, hook PostDeleteHookFn) {
	r.postDeleteHooksFn[entityname] = append(r.postDeleteHooksFn[entityname], hook)
}

// PreSingleHooksreturns a list of PreSingle hook functions
func (r resolverHooks) PreSingleHooks() PreSingleHooksFn {
	return r.preSingleHooksFn
}

// PostSingleHooks returns a list of PostSingle hook functions
func (r resolverHooks) PostSingleHooks() PostSingleHooksFn {
	return r.postSingleHooksFn
}

// PreListingHooks returns a list of PreListing hook functions
func (r resolverHooks) PreListingHooks() PreListingHooksFn {
	return r.preListingHooksFn
}

// PostListingHooksg returns a list of PostListing hook functions
func (r resolverHooks) PostListingHooks() PostListingHooksFn {
	return r.postListingHooksFn
}

// PreCreateHooksreturns a list of PreCreate hook functions
func (r resolverHooks) PreCreateHooks() PreCreateHooksFn {
	return r.preCreateHooksFn
}

// PostCreateHooks returns a list of PostCreate hook functions
func (r resolverHooks) PostCreateHooks() PostCreateHooksFn {
	return r.postCreateHooksFn
}

// PreUpdateHooksreturns a list of PreUpdate hook functions
func (r resolverHooks) PreUpdateHooks() PreUpdateHooksFn {
	return r.preUpdateHooksFn
}

// PostUpdateHooks returns a list of PostUpdate hook functions
func (r resolverHooks) PostUpdateHooks() PostUpdateHooksFn {
	return r.postUpdateHooksFn
}

// PreDeleteHooksreturns a list of PreDelete hook functions
func (r resolverHooks) PreDeleteHooks() PreDeleteHooksFn {
	return r.preDeleteHooksFn
}

// PostDeleteHooks returns a list of PostDelete hook functions
func (r resolverHooks) PostDeleteHooks() PostDeleteHooksFn {
	return r.postDeleteHooksFn
}
