package config

import "database/sql"

type (
	// Config from lardwaz app
	// TODO: might need to augment
	Config interface {
		Env() string
		DB() *sql.DB // Note: Don't like the ref!
		BaseURL() string
	}

	config struct {
		env     string
		db      *sql.DB
		baseURL string
	}
)

// NewConfig returns a new Config
func NewConfig(env string, db *sql.DB, baseURL string) Config {
	return &config{
		env:     env,
		db:      db,
		baseURL: baseURL,
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
