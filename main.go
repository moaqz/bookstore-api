package main

import (
	"fmt"
	"log"

	"github.com/techwithmat/booki-api/config"
)

func main() {
	// get configuration stucts
	configuration, err := config.NewConfig()

	if err != nil {
		log.Println("Error loading configuration file", err.Error())
	}

	fmt.Println(configuration.Database)
}
