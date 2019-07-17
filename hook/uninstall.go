package hook

type (
	// UninstallHookFn defines the signature of the Uninstall hook
	UninstallHookFn func() error

	// UninstallHook defines the Uninstall hook
	UninstallHook interface {
		RegisterNewHook(UninstallHookFn)
		Hook() UninstallHookFn
	}

	uninstallHook struct {
		hookFn UninstallHookFn
	}
)

// NewUninstallHook returns a new UninstallHook
func NewUninstallHook() UninstallHook {
	return &uninstallHook{}
}

// RegisterNewHook sets a new hook
func (h *uninstallHook) RegisterNewHook(hook UninstallHookFn) {
	h.hookFn = hook
}

// Hook returns the current hook
func (h uninstallHook) Hook() UninstallHookFn {
	return h.hookFn
}
