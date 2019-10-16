package hook

import (
	"go.lsl.digital/lardwaz/sdk/config"
	"go.lsl.digital/lardwaz/sdk/schematic"
)

type (
	// LoadedHookFn defines the signature of the Loaded hook
	// Parameters are configuration, plugin -> schematics, plugin_setting_name -> val
	LoadedHookFn func(config.Config, map[string]schematic.Schematics, map[string]interface{}) error

	// LoadedHook defines the Loaded hook
	LoadedHook interface {
		RegisterNewHook(LoadedHookFn)
		Hook() LoadedHookFn
	}

	loadedHook struct {
		hookFn LoadedHookFn
	}
)

// NewLoadedHook returns a new LoadedHook
func NewLoadedHook() LoadedHook {
	return &loadedHook{}
}

// RegisterNewHook sets a new hook
func (h *loadedHook) RegisterNewHook(hook LoadedHookFn) {
	h.hookFn = hook
}

// Hook returns the current hook
func (h loadedHook) Hook() LoadedHookFn {
	return h.hookFn
}
