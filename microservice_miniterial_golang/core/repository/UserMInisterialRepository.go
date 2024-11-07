package repository

import (
	"log"
	"sync"

	"microservice_ministerial.com/core/interfaces"
	"microservice_ministerial.com/infrastructure/database"
	"microservice_ministerial.com/infrastructure/entities"
	"microservice_ministerial.com/infrastructure/utils"
	"microservice_ministerial.com/usecase/dto"
)

var (
	_openInstance *OpenConnection
	_once         sync.Once
)

func UserMinisterialInstance() interfaces.IUserMinisterial {
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

func (db *OpenConnection) GetMinisterialAndUserByIdFindAll(userId uint) ([]dto.MinisterialResponseDTO, error) {

	var ministerials []dto.MinisterialResponseDTO
	db.mux.Lock()

	result := db.connection.Table("user_ministerials").Select("user_ministerials.Id", "ministerials.Name", "user_ministerials.ministerial_id ", "ministerials.created_at", "ministerials.updated_at", "ministerials.active").Joins("left join ministerials on user_ministerials.ministerial_id = ministerials.id").Where(utils.DB_EQUAL_USER_ID, userId).Find(&ministerials)
	defer db.mux.Unlock()
	defer database.Closedb()
	return ministerials, result.Error
}
func (db *OpenConnection) AddUserToMinisterial(userMinisterial entities.UserMinisterial) (entities.UserMinisterial, error) {
	db.mux.Lock()
	result := db.connection.Create(&userMinisterial)
	defer db.mux.Unlock()
	defer database.Closedb()
	return userMinisterial, result.Error
}
func (db *OpenConnection) DeleteUserMinisterial(id uint) (bool, error) {
	db.mux.Lock()
	result := db.connection.Where(utils.DB_EQUAL_ID, id).Delete(&entities.UserMinisterial{})
	defer db.mux.Unlock()
	defer database.Closedb()
	return true, result.Error
}
