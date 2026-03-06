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
	config.DB.AutoMigrate(&models.User{}, &models.Equipment{}, &models.RentalHistory{}, &models.TopUp{})

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Dependency Injection
	// User
	userRepo := repository.NewUserRepo(config.DB)
	userSvc := service.NewUserService(userRepo)
	userCtrl := controller.NewUserController(userSvc)

	// Equipment
	eqRepo := repository.NewEquipmentRepo(config.DB)
	eqSvc := service.NewEquipmentService(eqRepo)
	eqCtrl := controller.NewEquipmentController(eqSvc)

	// Rental
	rentalRepo := repository.NewRentalRepo(config.DB)
	rentalSvc := service.NewRentalService(rentalRepo, userRepo, eqRepo)
	rentalCtrl := controller.NewRentalController(rentalSvc)

	// Top Up
	topUpRepo := repository.NewTopUpRepo(config.DB)
	topUpSvc := service.NewTopUpService(topUpRepo)
	topUpCtrl := controller.NewTopUpController(topUpSvc)

	// Routes
	e.POST("/register", userCtrl.Register)
	e.POST("/login", userCtrl.Login)

	e.POST("/equipments", eqCtrl.Create)
	e.GET("/equipments", eqCtrl.GetAll)

	e.POST("/rentals", rentalCtrl.CreateRental)
	e.PUT("/rentals/return", rentalCtrl.ReturnRental)

	e.POST("/topup", topUpCtrl.RequestTopUp)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
