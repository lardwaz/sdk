package config

import (
	"database/sql"

	log "github.com/sirupsen/logrus"
)

type (
	// Config from lardwaz app
	// TODO: might need to augment
	Config interface {
		Env() string
		DB() *sql.DB // FIXME: Don't like the ref!
		BaseURL() string
		PortAPI() int
		PluginPath() string
		Logger() *log.Entry
	}

	config struct {
		env        string
		db         *sql.DB
		baseURL    string
		portAPI    int
		pluginPath string
		logger     *log.Entry
	}
)

// NewConfig returns a new Config
func NewConfig(env string, db *sql.DB, baseURL string, portAPI int, pluginPath string, logger *log.Entry) Config {
	return &config{
		env:        env,
		db:         db,
		baseURL:    baseURL,
		portAPI:    portAPI,
		pluginPath: pluginPath,
		logger:     logger,
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

// Logger returns the Logger
func (c *config) Logger() *log.Entry {
	return c.logger
}
