package hook

type (
	// InstallHookFn defines the signature of the install hook
	InstallHookFn func() error

	// InstallHook defines the install hook
	InstallHook interface {
		RegisterNewHook(InstallHookFn)
		Hook() InstallHookFn
	}

	installHook struct {
		hookFn InstallHookFn
	}
)

// NewInstallHook returns a new InstallHook
func NewInstallHook() InstallHook {
	return &installHook{}
}

// RegisterNewHook sets a new hook
func (h *installHook) RegisterNewHook(hook InstallHookFn) {
	h.hookFn = hook
}

// Hook returns the current hook
func (h installHook) Hook() InstallHookFn {
	return h.hookFn
}
