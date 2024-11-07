package repsitories

import (
	"log"
	"sync"

	"gorm.io/gorm"
	"microservice_user.com/core/interfaces"
	"microservice_user.com/infrastructure/database"
	"microservice_user.com/infrastructure/entities"
	"microservice_user.com/infrastructure/utils"
	"microservice_user.com/usecase/dto"
)

type OpenConnection struct {
	connection *gorm.DB
	mux        sync.Mutex
}

var (
	_openInstance *OpenConnection
	_once         sync.Once
)

// GetChurchInstance retorna una instancia singleton de OpenConnection que implementa interfaces.IChurch
func UserInstance() interfaces.IUser {
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
func (db *OpenConnection) GetUserFindAll() ([]entities.User, error) {
	var users []entities.User
	db.mux.Lock()
	result := db.connection.Find(&users)
	defer db.mux.Unlock()
	return users, result.Error
}
func (db *OpenConnection) GetUsersMembersFindAll() ([]entities.User, error) {
	var users []entities.User
	db.mux.Lock()
	result := db.connection.Where(utils.DB_ROL_ID_EQUAL, 6).Find(&users)
	defer db.mux.Unlock()
	return users, result.Error
}
func (db *OpenConnection) GetUserFindById(id uint) (dto.UserResposeDTO, error) {
	var user dto.UserResposeDTO
	db.mux.Lock()
	result := db.connection.Table("users").Find(&user, id)
	defer db.mux.Unlock()
	return user, result.Error
}
func (db *OpenConnection) GetUserFindByEmail(id uint, email string) (bool, error) {
	db.mux.Lock()
	query := db.connection.Where(utils.DB_EMAIL_EQUAL, email)
	if id > 0 {
		query = query.Where(utils.DB_DIFF_ID, id)
	}
	query = query.Find(&entities.User{})

	defer db.mux.Unlock()

	if query.RowsAffected == 0 {
		return false, query.Error
	}
	return true, query.Error
}
func (db *OpenConnection) GetUserFindByName(id uint, name string) (bool, error) {
	db.mux.Lock()
	query := db.connection.Where(utils.DB_NAME, name)
	if id > 0 {
		query = query.Where(utils.DB_DIFF_ID, id)
	}
	query = query.Find(&entities.User{})
	defer db.mux.Unlock()

	if query.RowsAffected == 0 {
		return false, query.Error
	}
	return true, query.Error
}
func (db *OpenConnection) GetUserFindByIdentification(id uint, identification string) (bool, error) {
	db.mux.Lock()
	query := db.connection.Where(utils.DB_IDENTIFICATION_EQUAL, identification)
	if id > 0 {
		query = query.Where(utils.DB_DIFF_ID, id)
	}
	query = query.Find(&entities.User{})

	defer db.mux.Unlock()

	if query.RowsAffected == 0 {
		return false, query.Error
	}
	return true, query.Error
}
func (db *OpenConnection) GetUserFindByUsername(id uint, username string) (bool, error) {
	db.mux.Lock()
	query := db.connection.Where(utils.DB_USERNAME_EQUAL, username)
	if id > 0 {
		query = query.Where(utils.DB_DIFF_ID, id)
	}
	query = query.Find(&entities.User{})

	defer db.mux.Unlock()

	if query.RowsAffected == 0 {
		return false, query.Error
	}
	return true, query.Error
}

func (db *OpenConnection) CreateUser(user entities.User) (entities.User, error) {
	db.mux.Lock()
	err := db.connection.Create(&user).Error
	defer db.mux.Unlock()
	return user, err
}
func (db *OpenConnection) UpdateUser(id uint, user entities.User) (entities.User, error) {
	db.mux.Lock()
	query := db.connection.Where(utils.DB_EQUAL_ID, id).Updates(&user)
	db.connection.Find(&user, id)
	defer db.mux.Unlock()
	return user, query.Error
}
func (db *OpenConnection) DeleteUser(id uint) (bool, error) {
	db.mux.Lock()
	query := db.connection.Delete(&entities.User{}, id)
	defer db.mux.Unlock()
	return true, query.Error
}
