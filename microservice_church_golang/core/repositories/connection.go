package repositories

import (
	"log"
	"sync"

	"gorm.io/gorm"
	"microservice_church.com/core/interfaces"
	"microservice_church.com/infrastructure/database"
)

type OpenConnection struct {
	connection *gorm.DB
	mux        sync.Mutex
}

var (
	_openInstance *OpenConnection
	_once         sync.Once
)

// GetChurchInstance retorna una instancia singleton de OpenConnection que implementa interfaces.IChurch
func GetChurchInstance() interfaces.IChurch {
	_once.Do(func() {
		db, err := database.DatabaseConnection()
		if err != nil {
			log.Fatalf("Error al conectar a la base de datos: %v", err)
		}
		_openInstance = &OpenConnection{
			connection: db,
		}
	})
	return _openInstance
}
