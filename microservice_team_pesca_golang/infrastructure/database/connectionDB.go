package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"microservice_team_pesca.com/infrastructure/entities"
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

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s  sslmode=disable ",
		DB_HOST,
		DB_USER,
		DB_PASSWORD,
		DB_NAME,
		DB_PORT,
	)
	//dsn := "host=localhost user=postgres password=12345678 dbname=msvc_team_pesca_db port=5432  sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entities.TeamPesca{})
	return db
}

// Close database connection method is closin a connection between your app db
func CloseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()

}

// Closedb
func Closedb() {
	var db *gorm.DB = DatabaseConnection()
	CloseConnection(db)

}
