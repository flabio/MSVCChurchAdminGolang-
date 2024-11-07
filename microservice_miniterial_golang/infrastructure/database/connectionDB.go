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

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error al cargar el archivo .env:", err.Error())
	}
}
func CreateDatabase() (string, error) {
	EnvLoad()
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_SSLMODE := os.Getenv("DB_SSLMODE")
	DB_NAME := os.Getenv("DB_NAME")
	strConnection := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=%s dbname=postgres", DB_HOST, DB_USER, DB_PASSWORD, DB_PORT, DB_SSLMODE)
	db, err := gorm.Open(postgres.Open(strConnection), &gorm.Config{})
	if err != nil {
		return "", fmt.Errorf("error al conectar a postgres: %w", err)
	}
	var exists int
	query := fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname ='%s'", DB_NAME)
	db.Raw(query).Scan(&exists)
	if exists == 0 {
		createDBSQL := fmt.Sprintf("CREATE DATABASE %s", DB_NAME)
		if err := db.Exec(createDBSQL).Error; err != nil {
			return "", fmt.Errorf("error al crear la base de datos: %w", err)
		}
		log.Printf("Base de datos '%s' creada exitosamente.\n", DB_NAME)
	} else {
		log.Printf("La base de datos '%s' ya existe.\n", DB_NAME)
	}
	return strConnection, nil
}
func DatabaseConnection() (*gorm.DB, error) {
	EnvLoad()
	DB_NAME := os.Getenv("DB_NAME")
	strConnection, err := CreateDatabase()
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("%s dbname=%s", strConnection, DB_NAME)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error al conectar a la base de datos %s: %w", DB_NAME, err)
	}
	if err := db.AutoMigrate(&entities.Ministerial{}, &entities.UserMinisterial{}); err != nil {
		return nil, fmt.Errorf("error al ejecutar AutoMigrate: %w", err)
	}
	return db, nil
}
func CloseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		log.Println("Error al obtener la conexión SQL:", err.Error())
		return
	}
	if err := dbSQL.Close(); err != nil {
		log.Println("Error al cerrar la conexión SQL:", err.Error())
	}
}
func Closedb() {
	db, err := DatabaseConnection()
	if err != nil {
		log.Println("Error al conectar a la base de datos:", err)
		return
	}
	CloseConnection(db)
}
