package config

import "database/sql"

type (
	// Config from lardwaz app
	// TODO: might need to augment
	Config interface {
		Env() string
		DB() *sql.DB // Note: Don't like the ref!
		BaseURL() string
		PortAPI() int
		PluginPath() string
	}

	config struct {
		env        string
		db         *sql.DB
		baseURL    string
		portAPI    int
		pluginPath string
	}
)

// NewConfig returns a new Config
func NewConfig(env string, db *sql.DB, baseURL string, portAPI int, pluginPath string) Config {
	return &config{
		env:        env,
		db:         db,
		baseURL:    baseURL,
		portAPI:    portAPI,
		pluginPath: pluginPath,
	}
}

// Env returns the Env
func (c *config) Env() string {
	return c.env
}

// DB returns the DB
func (c *config) DB() *sql.DB {
	return c.db
}

// BaseURL returns the BaseURL
func (c *config) BaseURL() string {
	return c.baseURL
}

// PortAPI returns the PortAPI
func (c *config) PortAPI() int {
	return c.portAPI
}

// PluginPath returns the PluginPath
func (c *config) PluginPath() string {
	return c.pluginPath
}
