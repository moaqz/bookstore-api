package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/techwithmat/booki-api/config"
)

// Connect to a database
func Connect(configuration *config.Database) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configuration.Host,
		configuration.Port,
		configuration.User,
		configuration.Password,
		configuration.DB,
	)

	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		return nil, err
	}

	log.Println("Database connected")

	return db, nil
}
