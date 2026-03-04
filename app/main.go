package main

import (
	"os"
	"sport-rental/config"
	"sport-rental/controller"
	"sport-rental/models"
	"sport-rental/repository"
	"sport-rental/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Connect to database
	config.InitDB()
	config.DB.AutoMigrate(&models.User{}, &models.Equipment{}, &models.RentalHistory{})

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Dependency Injection
	userRepo := repository.NewUserRepo(config.DB)
	userSvc := service.NewUserService(userRepo)
	userCtrl := controller.NewUserController(userSvc)

	// Routes
	e.POST("/register", userCtrl.Register)
	e.POST("/login", userCtrl.Login)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
