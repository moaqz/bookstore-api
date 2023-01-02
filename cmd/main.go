package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/techwithmat/booki-api/config"
	"github.com/techwithmat/booki-api/pkg/database"

	// Book Imports
	"github.com/techwithmat/booki-api/internal/books/repository"
	"github.com/techwithmat/booki-api/internal/books/usecase"
	"github.com/techwithmat/booki-api/internal/books/delivery"
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

	e := echo.New()
	router := e.Group("/api/v1")

	// initialize repos, services and handlers
	bookRepository := repository.NewBookRepository(db)
	bookUseCase := usecase.NewBookUseCase(bookRepository)
	delivery.NewBookHandler(router, bookUseCase)

	e.Logger.Fatal(e.Start(":" + configuration.Port))
}
