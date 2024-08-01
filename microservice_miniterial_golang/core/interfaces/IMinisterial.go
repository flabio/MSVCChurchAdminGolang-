package interfaces

import "microservice_ministerial.com/infrastructure/entities"

type IMinisterial interface {
	GetMinisterialAll() ([]entities.Ministerial, error)
	GetMinisterialFindById(id uint) (entities.Ministerial, error)
	GetMinisterialFindByName(id uint, name string) (bool, error)
	CreateMinisterial(ministerial entities.Ministerial) (entities.Ministerial, error)
	UpdateMinisterial(id uint, ministerial entities.Ministerial) (entities.Ministerial, error)
	DeleteMinisterial(id uint) (bool, error)
}
