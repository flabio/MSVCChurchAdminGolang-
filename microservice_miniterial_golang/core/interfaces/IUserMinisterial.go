package interfaces

import (
	"microservice_ministerial.com/infrastructure/entities"
	"microservice_ministerial.com/usecase/dto"
)

type IUserMinisterial interface {
	GetMinisterialAndUserByIdFindAll(userId uint) ([]dto.MinisterialResponseDTO, error)
	AddUserToMinisterial(userMinisterial entities.UserMinisterial) (entities.UserMinisterial, error)
	DeleteUserMinisterial(id uint) (bool, error)
}
