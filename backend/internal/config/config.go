// /////////////////////////////////////////////////////////
// src: ./internal/config/config.go 					//
// desc: config package provides configuration loading.//
// //////////////////////////////////////////////////////

package config

import (
	"os"
	"strings"
)

// Config holds the application configuration.
type Config struct {
	DSN      string // holds the database connection values
	Port     string
	LogLevel string
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {

	dsn := os.Getenv("DATABASE_URL")
	dsn = strings.TrimSpace(dsn)

	return &Config{
		DSN:      dsn,
		Port:     os.Getenv("PORT"),
		LogLevel: os.Getenv("LOG_LEVEL"),
	}, nil
}
