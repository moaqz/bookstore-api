package main

import (
	"fmt"

	"github.com/techwithmat/booki-api/config"
	"github.com/techwithmat/booki-api/data/database"
)

func main() {
	// get configuration stucts
	configuration, err := config.NewConfig()

	if err != nil {
		fmt.Printf("Error loading configuration file: %v", err)
	}

	// establish DB connection
	db, err := database.Connect(configuration.Database)

	if err != nil {
		fmt.Printf("Unable to connect to database: %v", err)
	}

	defer db.Close()
}
