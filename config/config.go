package config

import (
	"github.com/techwithmat/bookery-api/pkg/utils"
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
	utils.LoadEnv()

	port := utils.MustGet("PORT")

	// set default PORT if missing
	if port == "" {
		port = "3000"
	}

	return &Config{
		Port: port,
		Database: &Database{
			Host:     utils.MustGet("POSTGRES_HOST"),
			Port:     utils.MustGet("POSTGRES_PORT"),
			User:     utils.MustGet("POSTGRES_USER"),
			DB:       utils.MustGet("POSTGRES_DB"),
			Password: utils.MustGet("POSTGRES_PASSWORD"),
		},
	}, nil
}
