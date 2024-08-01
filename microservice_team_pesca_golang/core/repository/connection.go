package repository

import (
	"sync"

	"gorm.io/gorm"
	"microservice_team_pesca.com/core/interfaces"
	"microservice_team_pesca.com/infrastructure/database"
)

type OpenConnection struct {
	connection *gorm.DB
	mux        sync.Mutex
}

func TeamPescaInstance() interfaces.ITeamPesca {
	var (
		_OPEN *OpenConnection
		_ONCE sync.Once
	)
	_ONCE.Do(func() {
		_OPEN = &OpenConnection{
			connection: database.DatabaseConnection(),
		}
	})
	return _OPEN
}
