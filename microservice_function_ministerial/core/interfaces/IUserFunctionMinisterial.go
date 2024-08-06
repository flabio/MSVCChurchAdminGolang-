package interfaces

import (
	"msvc_function_ministerial/infrastructure/entities"
	"msvc_function_ministerial/usecase/dto"
)

type IUserFunctionMinisterial interface {
	GetFunctionMinisterialAndUserByIdFindAll(userId uint) ([]dto.FunctionMinisterialResponseDTO, error)
	AddUserToFunctionMinisterial(userMinisterial entities.UserFunctionMinisterial) (entities.UserFunctionMinisterial, error)
	DeleteUserFunctionMinisterial(id uint) (bool, error)
}
