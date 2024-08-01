package repositories

import (
	"microservice_church.com/infrastructure/database"
	"microservice_church.com/infrastructure/entities"
	"microservice_church.com/infrastructure/utils"
)

func (db *OpenConnection) GetChurchFindAll() ([]entities.Church, error) {
	var entities []entities.Church
	db.mux.Lock()
	result := db.connection.Order(utils.DB_ORDER_BY_ID).Find(&entities)
	defer database.Closedb()
	defer db.mux.Unlock()
	return entities, result.Error
}
func (db *OpenConnection) GetChurchFindById(id uint) (entities.Church, error) {
	var entities entities.Church
	db.mux.Lock()
	result := db.connection.Where(utils.DB_EQUAL_ID, id).Find(&entities)
	defer database.Closedb()
	defer db.mux.Unlock()
	return entities, result.Error
}
func (db *OpenConnection) GetChurchFindByName(id uint, name string) (bool, error) {
	var entities entities.Church
	db.mux.Lock()
	query := db.connection.Where(utils.DB_NAME, name)
	if id > 0 {
		query = query.Where(utils.DB_DIFF_ID, id)
	}
	result := query.Find(&entities)
	defer database.Closedb()
	defer db.mux.Unlock()
	if result.RowsAffected == 0 {
		return false, result.Error
	}
	return true, result.Error
}

func (db *OpenConnection) GetChurchFindByEmail(id uint, email string) (bool, error) {
	var entities entities.Church
	db.mux.Lock()
	query := db.connection.Where(utils.DB_EMAIL_EQUAL, email)
	if id > 0 {
		query = query.Where(utils.DB_DIFF_ID, id)
	}
	query = query.Find(&entities)
	defer database.Closedb()
	defer db.mux.Unlock()
	if query.RowsAffected == 0 {
		return false, query.Error
	}
	return true, query.Error
}

func (db *OpenConnection) CreateChurch(entities entities.Church) (entities.Church, error) {
	db.mux.Lock()
	err := db.connection.Create(&entities).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	return entities, err
}
func (db *OpenConnection) UpdateChurch(id uint, entities entities.Church) (entities.Church, error) {
	db.mux.Lock()
	err := db.connection.Where(utils.DB_EQUAL_ID, id).Updates(&entities).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	return entities, err
}
func (db *OpenConnection) DeleteChurch(id uint) (bool, error) {
	db.mux.Lock()
	var entities entities.Church
	err := db.connection.Where(utils.DB_EQUAL_ID, id).Delete(&entities).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	if err == nil {
		return true, err
	}
	return false, err
}
