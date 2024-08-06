package repositories

import (
	"msvc_function_ministerial/core/interfaces"
	"msvc_function_ministerial/infrastructure/database"
	"msvc_function_ministerial/infrastructure/entities"
	"msvc_function_ministerial/infrastructure/utils"
	"msvc_function_ministerial/usecase/dto"
	"sync"
)

func UserFunctionMinisterialInstance() interfaces.IUserFunctionMinisterial {
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

func (db *OpenConnection) GetFunctionMinisterialAndUserByIdFindAll(userId uint) ([]dto.FunctionMinisterialResponseDTO, error) {

	var ministerials []dto.FunctionMinisterialResponseDTO
	db.mux.Lock()

	result := db.connection.Table("user_function_ministerials").Select("user_function_ministerials.Id", "function_ministerials.Name", "user_function_ministerials.function_ministerial_id ", "function_ministerials.created_at", "function_ministerials.updated_at", "function_ministerials.active").Joins("left join function_ministerials on user_function_ministerials.function_ministerial_id = function_ministerials.id").Where(utils.DB_EQUAL_USER_ID, userId).Find(&ministerials)
	defer db.mux.Unlock()
	defer database.Closedb()
	return ministerials, result.Error
}
func (db *OpenConnection) AddUserToFunctionMinisterial(userMinisterial entities.UserFunctionMinisterial) (entities.UserFunctionMinisterial, error) {
	db.mux.Lock()
	result := db.connection.Create(&userMinisterial)
	defer db.mux.Unlock()
	defer database.Closedb()
	return userMinisterial, result.Error
}
func (db *OpenConnection) DeleteUserFunctionMinisterial(id uint) (bool, error) {
	db.mux.Lock()
	result := db.connection.Where(utils.DB_EQUAL_ID, id).Delete(&entities.UserFunctionMinisterial{})
	defer db.mux.Unlock()
	defer database.Closedb()
	return true, result.Error
}
