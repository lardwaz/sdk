package hook

import (
	"github.com/spf13/cobra"
)

type (
	// CommandHookFn defines the signature of <any> command hook function
	CommandHookFn func(*cobra.Command) error

	// CommandHook defines a list of HookFn
	CommandHook interface {
		RegisterNewHook(CommandHookFn)
		Hook() CommandHookFn
	}

	commandHook struct {
		hookFn CommandHookFn
	}
)

// NewCommandHook returns a new CommandHook
func NewCommandHook() CommandHook {
	return &commandHook{}
}

// RegisterNewHook registers a new hook function
func (h *commandHook) RegisterNewHook(hook CommandHookFn) {
	h.hookFn = hook
}

// Hook return a list of hook function
func (h commandHook) Hook() CommandHookFn {
	return h.hookFn
}
