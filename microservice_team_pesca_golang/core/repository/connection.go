package repository

import (
	"log"
	"sync"

	"gorm.io/gorm"
	"microservice_team_pesca.com/core/interfaces"
	"microservice_team_pesca.com/infrastructure/database"
)

type OpenConnection struct {
	connection *gorm.DB
	mux        sync.Mutex
}

var (
	_openInstance *OpenConnection
	_once         sync.Once
)

func TeamPescaInstance() interfaces.ITeamPesca {
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
