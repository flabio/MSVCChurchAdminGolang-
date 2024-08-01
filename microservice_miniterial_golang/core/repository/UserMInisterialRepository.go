package repository

import (
	"sync"

	"microservice_ministerial.com/core/interfaces"
	"microservice_ministerial.com/infrastructure/database"
	"microservice_ministerial.com/infrastructure/entities"
	"microservice_ministerial.com/infrastructure/utils"
)

func UserMinisterialInstance() interfaces.IUserMinisterial {
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

func (db *OpenConnection) GetMinisterialAndUserByIdFindAll(userId uint) ([]entities.UserMinisterial, error) {

	var ministerials []entities.UserMinisterial
	db.mux.Lock()

	result := db.connection.Preload(utils.DB_MINISTERIAL).Where(utils.DB_EQUAL_USER_ID, userId).Find(&ministerials)
	defer db.mux.Unlock()
	defer database.CloseConnection()
	return ministerials, result.Error
}
func (db *OpenConnection) AddUserToMinisterial(userMinisterial entities.UserMinisterial) (entities.UserMinisterial, error) {
	db.mux.Lock()
	result := db.connection.Create(&userMinisterial)
	defer db.mux.Unlock()
	defer database.CloseConnection()
	return userMinisterial, result.Error
}
func (db *OpenConnection) DeleteUserMinisterial(id uint) (bool, error) {
	db.mux.Lock()
	result := db.connection.Where(utils.DB_EQUAL_ID, id).Delete(&entities.UserMinisterial{})
	defer db.mux.Unlock()
	defer database.CloseConnection()
	return true, result.Error
}
