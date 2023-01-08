package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/techwithmat/bookery-api/config"
	"github.com/techwithmat/bookery-api/pkg/database"

	// Book Imports
	bookDelivery "github.com/techwithmat/bookery-api/internal/books/delivery"
	bookRepository "github.com/techwithmat/bookery-api/internal/books/repository"
	bookUseCase "github.com/techwithmat/bookery-api/internal/books/usecase"

	// User Imports
	userDelivery "github.com/techwithmat/bookery-api/internal/users/delivery"
	userRepository "github.com/techwithmat/bookery-api/internal/users/repository"
	userUseCase "github.com/techwithmat/bookery-api/internal/users/usecase"
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
	e.Use(middleware.Logger())
	router := e.Group("/api/v1")

	// initialize repos, services and handlers
	bookRepository := bookRepository.NewBookRepository(db)
	bookUseCase := bookUseCase.NewBookUseCase(bookRepository)
	bookDelivery.NewBookHandler(router, bookUseCase)

	userRepository := userRepository.NewUsersRepo(db)
	userUseCase := userUseCase.NewUserUseCase(userRepository)
	userDelivery.NewUserHandler(router, userUseCase)

	e.Logger.Fatal(e.Start(":" + configuration.Port))
}
