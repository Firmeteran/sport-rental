package main

import (
	"os"
	"sport-rental/config"

	"github.com/labstack/echo/v4"
)

func main() {
	// Connect to database
	config.InitDB()

	e := echo.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
