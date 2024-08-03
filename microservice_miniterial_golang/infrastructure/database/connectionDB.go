package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"microservice_ministerial.com/infrastructure/entities"
)

func DatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading.env file")
	}
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	db.AutoMigrate(&entities.Ministerial{}, &entities.UserMinisterial{})
	return db
}

// Close database connection method is closin a connection between your app db

func CloseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic(err.Error())
	}
	dbSQL.Close()
}

func Closedb() {
	var db *gorm.DB = DatabaseConnection()
	CloseConnection(db)
}
