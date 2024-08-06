package interfaces

import "msvc_function_ministerial/infrastructure/entities"

type IFunctionMinisterial interface {
	GetFindAll() ([]entities.FunctionMinisterial, error)
	GetFindById(id int) (entities.FunctionMinisterial, error)
	GetFindByName(id uint, name string) (bool, error)
	Create(FunctionMinisterial entities.FunctionMinisterial) (entities.FunctionMinisterial, error)
	Update(id uint, FunctionMinisterial entities.FunctionMinisterial) (entities.FunctionMinisterial, error)
	Delete(id uint) (bool, error)
}
