package config

import (
	"github.com/techwithmat/bookery-api/pkg/utils/env"
)

// Config is a struct that contains configuration variables
type Config struct {
	Port     string
	Database *Database
}

// Database is a struct that contains DB's configuration variables
type Database struct {
	Host     string
	Port     string
	User     string
	DB       string
	Password string
}

// NewConfig creates a new Config struct
func NewConfig() (*Config, error) {
	env.LoadEnv()

	port := env.MustGet("PORT")

	// set default PORT if missing
	if port == "" {
		port = "3000"
	}

	return &Config{
		Port: port,
		Database: &Database{
			Host:     env.MustGet("POSTGRES_HOST"),
			Port:     env.MustGet("POSTGRES_PORT"),
			User:     env.MustGet("POSTGRES_USER"),
			DB:       env.MustGet("POSTGRES_DB"),
			Password: env.MustGet("POSTGRES_PASSWORD"),
		},
	}, nil
}
