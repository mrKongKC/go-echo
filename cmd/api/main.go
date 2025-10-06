package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	err := godotenv.Load()
	port := os.Getenv("APP_PORT")
	appAddress := fmt.Sprintf("localhost:%s", port)

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo 👋")
	})

	e.Logger.Fatal(e.Start(appAddress))
}
