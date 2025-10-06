package main

import (
	"fmt"
	"go-echo-app/cmd/api/handlers"
	"go-echo-app/common"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type application struct {
	logger   echo.Logger
	server   *echo.Echo
	handlers handlers.Handler
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	appAddress := fmt.Sprintf(":%s", port)

	// เชื่อม DB
	common.Connect()

	e := echo.New()

	h := handlers.Handler{DB: common.DB}

	app := application{
		logger:   e.Logger,
		server:   e,
		handlers: h,
	}

	fmt.Println(app)

	e.GET("/customers", h.GetCustomers)
	e.POST("/customers", h.CreateCustomer)

	e.Logger.Fatal(e.Start(appAddress))
}
