package repositories

import (
	"log"
	"msvc_function_ministerial/core/interfaces"
	"msvc_function_ministerial/infrastructure/database"
	"msvc_function_ministerial/infrastructure/entities"
	"msvc_function_ministerial/infrastructure/utils"
)

func GetFunctionMinisterialInstance() interfaces.IFunctionMinisterial {
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
func (db *OpenConnection) GetFindAll() ([]entities.FunctionMinisterial, error) {
	var roles []entities.FunctionMinisterial
	db.mux.Lock()
	result := db.connection.Order(utils.DB_ORDER_DESC).Find(&roles)
	defer database.Closedb()
	defer db.mux.Unlock()
	return roles, result.Error
}

func (db *OpenConnection) GetFindById(id int) (entities.FunctionMinisterial, error) {
	db.mux.Lock()
	var rol entities.FunctionMinisterial
	result := db.connection.Find(&rol, id)
	defer database.Closedb()
	defer db.mux.Unlock()
	return rol, result.Error
}

/*
param:rol is a struct
*/
func (db *OpenConnection) Create(rol entities.FunctionMinisterial) (entities.FunctionMinisterial, error) {
	db.mux.Lock()
	err := db.connection.Create(&rol).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	return rol, err
}

/*
@Params: rol Rol is a struct, id is an integer
*/
func (db *OpenConnection) Update(id uint, rol entities.FunctionMinisterial) (entities.FunctionMinisterial, error) {
	db.mux.Lock()
	result := db.connection.Where(utils.DB_EQUAL_ID, id).Updates(&rol)
	defer database.Closedb()
	defer db.mux.Unlock()
	return rol, result.Error

}

/*
@param: id is an int
*/
func (db *OpenConnection) Delete(id uint) (bool, error) {
	db.mux.Lock()
	var rol entities.FunctionMinisterial
	result := db.connection.Where(utils.DB_EQUAL_ID, id).Delete(&rol)

	defer database.Closedb()
	defer db.mux.Unlock()
	if result.RowsAffected == 0 {
		return true, result.Error
	}
	return false, result.Error
}

/*
@params: id is an uint number and name is a string
*/
func (db *OpenConnection) GetFindByName(id uint, name string) (bool, error) {
	db.mux.Lock()
	var rol entities.FunctionMinisterial

	query := db.connection.Where(utils.DB_NAME, name)
	if id > 0 {
		query = query.Where(utils.DB_DIFF_ID, id)
	}
	query = query.Find(&rol)

	defer database.Closedb()
	defer db.mux.Unlock()
	if query.RowsAffected == 1 {
		return true, query.Error
	}
	return false, query.Error

}
