package info

// Info holds information about the plugin
type Info interface {
	Name() string
	MachineName() string
	Description() string
	Logo() string
	Author() string
	Maintainers() []string
	Version() string
	Depends() []string
}

type info struct {
	name        string
	machinename string
	description string
	logo        string
	author      string
	maintainers []string
	version     string
	depends     []string
}

// New returns a new Info
func New(name, machinename, description, logo, author string, maintainers []string, version string, depends ...string) Info {
	return info{
		name:        name,
		machinename: machinename,
		description: description,
		logo:        logo,
		author:      author,
		maintainers: maintainers,
		version:     version,
		depends:     depends,
	}
}

// Name returns the user friendly name of the plugin
func (i info) Name() string {
	return i.name
}

// MachineName returns the machine friendly name of the plugin
func (i info) MachineName() string {
	return i.machinename
}

// Description returns the description of the plugin
func (i info) Description() string {
	return i.description
}

// Logo returns the logo of the plugin
func (i info) Logo() string {
	return i.logo
}

// Author returns the author of the plugin
func (i info) Author() string {
	return i.author
}

// Maintainers returns the maintainers of the plugin
// applicable mostly to open-source plugins
func (i info) Maintainers() []string {
	return i.maintainers
}

// Version returns the version of the plugin
func (i info) Version() string {
	return i.version
}

// Depends returns the machine friendly name of other plugins
// the plugin is dependent on
func (i info) Depends() []string {
	return i.depends
}
