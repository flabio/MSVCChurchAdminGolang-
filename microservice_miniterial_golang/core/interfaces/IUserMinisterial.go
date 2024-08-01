package interfaces

import "microservice_ministerial.com/infrastructure/entities"

type IUserMinisterial interface {
	GetMinisterialAndUserByIdFindAll(userId uint) ([]entities.UserMinisterial, error)
	AddUserToMinisterial(userMinisterial entities.UserMinisterial) (entities.UserMinisterial, error)
	DeleteUserMinisterial(id uint) (bool, error)
}
