package repository

import (
	"sync"

	"microservice_ministerial.com/core/interfaces"
	"microservice_ministerial.com/infrastructure/database"
	"microservice_ministerial.com/infrastructure/entities"
	"microservice_ministerial.com/infrastructure/utils"
)

func Instance() interfaces.IMinisterial {
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

func (db *OpenConnection) GetMinisterialAll() ([]entities.Ministerial, error) {

	var ministerials []entities.Ministerial
	db.mux.Lock()
	result := db.connection.Order(utils.DB_ORDER_DESC).Find(&ministerials)
	defer db.mux.Unlock()
	defer database.CloseConnection()
	return ministerials, result.Error
}
func (db *OpenConnection) GetMinisterialFindById(id uint) (entities.Ministerial, error) {
	var ministerial entities.Ministerial
	db.mux.Lock()
	result := db.connection.Find(&ministerial, id)
	defer db.mux.Unlock()
	defer database.CloseConnection()
	return ministerial, result.Error
}
func (db *OpenConnection) GetMinisterialFindByName(id uint, name string) (bool, error) {
	db.mux.Lock()
	query := db.connection.Where(utils.DB_NAME_EQUAL, name)
	if id > 0 {
		query = query.Where(utils.DB_DIFF_ID, id).Find(&entities.Ministerial{})
	} else {
		query = query.Find(&entities.Ministerial{})
	}

	defer db.mux.Unlock()
	defer database.CloseConnection()
	if query.RowsAffected == 1 {
		return true, query.Error
	}
	return false, query.Error
}
func (db *OpenConnection) CreateMinisterial(ministerial entities.Ministerial) (entities.Ministerial, error) {
	db.mux.Lock()
	result := db.connection.Create(&ministerial)
	defer db.mux.Unlock()
	defer database.CloseConnection()
	return ministerial, result.Error
}
func (db *OpenConnection) UpdateMinisterial(id uint, ministerial entities.Ministerial) (entities.Ministerial, error) {
	db.mux.Lock()
	result := db.connection.Where(utils.DB_EQUAL_ID, id).Updates(&ministerial)
	db.connection.Find(&ministerial, id)
	defer db.mux.Unlock()
	defer database.CloseConnection()
	return ministerial, result.Error
}
func (db *OpenConnection) DeleteMinisterial(id uint) (bool, error) {
	db.mux.Lock()
	result := db.connection.Where(utils.DB_EQUAL_ID, id).Delete(&entities.Ministerial{})
	defer db.mux.Unlock()
	defer database.CloseConnection()
	return true, result.Error
}
