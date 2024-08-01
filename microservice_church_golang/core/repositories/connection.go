package repositories

import (
	"sync"

	"gorm.io/gorm"
	"microservice_church.com/core/interfaces"
	"microservice_church.com/infrastructure/database"
)

type OpenConnection struct {
	connection *gorm.DB
	mux        sync.Mutex
}

func GetChurchInstance() interfaces.IChurch {
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
