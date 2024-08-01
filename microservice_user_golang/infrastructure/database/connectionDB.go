package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"microservice_user.com/infrastructure/entities"
)

func DatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic(errEnv.Error())
	}
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&entities.User{})
	return db
}

// Close database connection method is closin a connection between your app db
func CloseConnection() {
	var db *gorm.DB = DatabaseConnection()
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
