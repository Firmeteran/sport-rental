package config

import (
	"fmt"
	"os"
	"sport-rental/models"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require default_query_exec_mode=simple_protocol",
		dbHost, dbUser, dbPass, dbName, dbPort,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		PrepareStmt:            false,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	err = db.AutoMigrate(&models.User{}, &models.Equipment{}, &models.TopUp{}, &models.RentalHistory{})
	if err != nil {
		fmt.Println("Migration Error:", err)
	}

	DB = db
	fmt.Println("Connection and migration is successful.")
}
