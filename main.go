package main

import (
	"fmt"
	"os"

	"github.com/enylvia/shorten-link/handler"
	"github.com/enylvia/shorten-link/repository"
	"github.com/enylvia/shorten-link/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	// dependencies injection

	shortRepository := repository.NewShortenRepository()
	shortService := service.NewShortenService(shortRepository)
	shortHandler := handler.NewShortenHandler(shortService)

	// initialize database connection
	router := echo.New()

	// initialize routes
	router.GET("/:url", shortHandler.ResolveURL)
	router.POST("/api/v1", shortHandler.ShortenURL)

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	router.Logger.Fatal(router.Start(":" + os.Getenv("APP_PORT")))

}
