package interfaces

import "microservice_user.com/infrastructure/entities"

type IUser interface {
	GetUserFindAll() ([]entities.User, error)
	GetUserFindById(id uint) (entities.User, error)
	GetUsersMembersFindAll() ([]entities.User, error)
	GetUserFindByName(id uint, name string) (bool, error)
	GetUserFindByEmail(id uint, email string) (bool, error)
	GetUserFindByIdentification(id uint, identification string) (bool, error)
	GetUserFindByUsername(id uint, username string) (bool, error)
	CreateUser(user entities.User) (entities.User, error)
	UpdateUser(id uint, user entities.User) (entities.User, error)
	DeleteUser(id uint) (bool, error)
}
