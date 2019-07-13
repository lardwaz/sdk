package hook

type (
	// GlobalPreSingleHooksFn defines a map of entityname to resolver hook function
	GlobalPreSingleHooksFn []PreSingleHookFn

	// GlobalPostSingleHooksFn defines a map of entityname to resolver hook function
	GlobalPostSingleHooksFn []PostSingleHookFn

	// GlobalPreListingHooksFn defines a map of entityname to resolver hook function
	GlobalPreListingHooksFn []PreListingHookFn

	// GlobalPostListingHooksFn defines a map of entityname to resolver hook function
	GlobalPostListingHooksFn []PostListingHookFn

	// GlobalPreCreateHooksFn defines a map of entityname to resolver hook function
	GlobalPreCreateHooksFn []PreCreateHookFn

	// GlobalPostCreateHooksFn defines a map of entityname to resolver hook function
	GlobalPostCreateHooksFn []PostCreateHookFn

	// GlobalPreUpdateHooksFn defines a map of entityname to resolver hook function
	GlobalPreUpdateHooksFn []PreUpdateHookFn

	// GlobalPostUpdateHooksFn defines a map of entityname to resolver hook function
	GlobalPostUpdateHooksFn []PostUpdateHookFn

	// GlobalPreDeleteHooksFn defines a map of entityname to resolver hook function
	GlobalPreDeleteHooksFn []PreDeleteHookFn

	// GlobalPostDeleteHooksFn defines a map of entityname to resolver hook function
	GlobalPostDeleteHooksFn []PostDeleteHookFn

	// GlobalResolverHooks defines a list of common hooks for the various resolvers
	GlobalResolverHooks interface {
		RegisterNewGlobalPreSingleHook(hook PreSingleHookFn)
		RegisterNewGlobalPostSingleHook(hook PostSingleHookFn)
		RegisterNewGlobalPreListingHook(hook PreListingHookFn)
		RegisterNewGlobalPostListingHook(hook PostListingHookFn)
		RegisterNewGlobalPreCreateHook(hook PreCreateHookFn)
		RegisterNewGlobalPostCreateHook(hook PostCreateHookFn)
		RegisterNewGlobalPreUpdateHook(hook PreUpdateHookFn)
		RegisterNewGlobalPostUpdateHook(hook PostUpdateHookFn)
		RegisterNewGlobalPreDeleteHook(hook PreDeleteHookFn)
		RegisterNewGlobalPostDeleteHook(hook PostDeleteHookFn)

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
func (r globalResolverHooks) RegisterNewGlobalPreSingleHook(hook PreSingleHookFn) {
	r.globalPreSingleHooksFn = append(r.globalPreSingleHooksFn, hook)
}

// RegisterNewGlobalPostSingleHook register a new hook to the PostSingle hook functions
func (r globalResolverHooks) RegisterNewGlobalPostSingleHook(hook PostSingleHookFn) {
	r.globalPostSingleHooksFn = append(r.globalPostSingleHooksFn, hook)
}

// RegisterNewGlobalPreListingHook register a new hook to the PreListing hook functions
func (r globalResolverHooks) RegisterNewGlobalPreListingHook(hook PreListingHookFn) {
	r.globalPreListingHooksFn = append(r.globalPreListingHooksFn, hook)
}

// RegisterNewGlobalPostListingHookgregister a new hook to the PostListing hook functions
func (r globalResolverHooks) RegisterNewGlobalPostListingHook(hook PostListingHookFn) {
	r.globalPostListingHooksFn = append(r.globalPostListingHooksFn, hook)
}

// RegisterNewGlobalPreCreateHookrregister a new hook to the PreCreate hook functions
func (r globalResolverHooks) RegisterNewGlobalPreCreateHook(hook PreCreateHookFn) {
	r.globalPreCreateHooksFn = append(r.globalPreCreateHooksFn, hook)
}

// RegisterNewGlobalPostCreateHook register a new hook to the PostCreate hook functions
func (r globalResolverHooks) RegisterNewGlobalPostCreateHook(hook PostCreateHookFn) {
	r.globalPostCreateHooksFn = append(r.globalPostCreateHooksFn, hook)
}

// RegisterNewGlobalPreUpdateHookrregister a new hook to the PreUpdate hook functions
func (r globalResolverHooks) RegisterNewGlobalPreUpdateHook(hook PreUpdateHookFn) {
	r.globalPreUpdateHooksFn = append(r.globalPreUpdateHooksFn, hook)
}

// RegisterNewGlobalPostUpdateHook register a new hook to the PostUpdate hook functions
func (r globalResolverHooks) RegisterNewGlobalPostUpdateHook(hook PostUpdateHookFn) {
	r.globalPostUpdateHooksFn = append(r.globalPostUpdateHooksFn, hook)
}

// RegisterNewGlobalPreDeleteHookrregister a new hook to the PreDelete hook functions
func (r globalResolverHooks) RegisterNewGlobalPreDeleteHook(hook PreDeleteHookFn) {
	r.globalPreDeleteHooksFn = append(r.globalPreDeleteHooksFn, hook)
}

// RegisterNewGlobalPostDeleteHook register a new hook to the PostDelete hook functions
func (r globalResolverHooks) RegisterNewGlobalPostDeleteHook(hook PostDeleteHookFn) {
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
