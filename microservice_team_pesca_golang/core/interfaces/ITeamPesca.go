package interfaces

import "microservice_team_pesca.com/infrastructure/entities"

type ITeamPesca interface {
	GetTeamPescaFindAll() ([]entities.TeamPesca, error)
	GetTeamPescaFindById(id uint) (entities.TeamPesca, error)
	GetTeamPescaFindByName(id uint, name string) (bool, error)
	CreateTeamPesca(teamPesca entities.TeamPesca) (entities.TeamPesca, error)
	UpdateTeamPesca(id uint, teamPesca entities.TeamPesca) (entities.TeamPesca, error)
	DeleteTeamPesca(id uint) (bool, error)
}
