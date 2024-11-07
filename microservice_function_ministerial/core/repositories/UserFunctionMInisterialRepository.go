package repositories

import (
	"log"
	"msvc_function_ministerial/core/interfaces"
	"msvc_function_ministerial/infrastructure/database"
	"msvc_function_ministerial/infrastructure/entities"
	"msvc_function_ministerial/infrastructure/utils"
	"msvc_function_ministerial/usecase/dto"
	"sync"
)

var (
	_openInstance *OpenConnection
	_once         sync.Once
)

func UserFunctionMinisterialInstance() interfaces.IUserFunctionMinisterial {
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
func (db *OpenConnection) GetFunctionMinisterialAndUserByIdFindAll(userId uint) ([]dto.FunctionMinisterialResponseDTO, error) {
	var ministerials []dto.FunctionMinisterialResponseDTO
	db.mux.Lock()
	result := db.connection.Table("user_function_ministerials").Select("user_function_ministerials.Id", "function_ministerials.Name", "user_function_ministerials.function_ministerial_id ", "function_ministerials.created_at", "function_ministerials.updated_at", "function_ministerials.active").Joins("left join function_ministerials on user_function_ministerials.function_ministerial_id = function_ministerials.id").Where(utils.DB_EQUAL_USER_ID, userId).Find(&ministerials)
	defer db.mux.Unlock()
	return ministerials, result.Error
}
func (db *OpenConnection) AddUserToFunctionMinisterial(userMinisterial entities.UserFunctionMinisterial) (entities.UserFunctionMinisterial, error) {
	db.mux.Lock()
	result := db.connection.Create(&userMinisterial)
	defer db.mux.Unlock()
	return userMinisterial, result.Error
}
func (db *OpenConnection) DeleteUserFunctionMinisterial(id uint) (bool, error) {
	db.mux.Lock()
	result := db.connection.Where(utils.DB_EQUAL_ID, id).Delete(&entities.UserFunctionMinisterial{})
	defer db.mux.Unlock()
	return true, result.Error
}
