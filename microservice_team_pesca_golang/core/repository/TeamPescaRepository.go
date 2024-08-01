package repository

import (
	"microservice_team_pesca.com/infrastructure/database"
	"microservice_team_pesca.com/infrastructure/entities"
	"microservice_team_pesca.com/infrastructure/utils"
)

func (db *OpenConnection) GetTeamPescaFindAll() ([]entities.TeamPesca, error) {
	var teamPescas []entities.TeamPesca

	db.mux.Lock()
	result := db.connection.Order(utils.DB_ORDER_DESC).Find(&teamPescas)
	defer db.mux.Unlock()
	defer database.Closedb()
	return teamPescas, result.Error
}

func (db *OpenConnection) GetTeamPescaFindById(id uint) (entities.TeamPesca, error) {
	var teamPesca entities.TeamPesca
	db.mux.Lock()
	result := db.connection.Where(utils.DB_EQUAL_ID, id).Find(&teamPesca)
	defer db.mux.Unlock()
	defer database.Closedb()
	return teamPesca, result.Error
}

func (db *OpenConnection) CreateTeamPesca(teamPesca entities.TeamPesca) (entities.TeamPesca, error) {
	db.mux.Lock()
	result := db.connection.Create(&teamPesca)
	defer db.mux.Unlock()
	defer database.Closedb()
	return teamPesca, result.Error
}

func (db *OpenConnection) UpdateTeamPesca(id uint, teamPesca entities.TeamPesca) (entities.TeamPesca, error) {
	db.mux.Lock()
	result := db.connection.Where(utils.DB_EQUAL_ID, id).Updates(&teamPesca)
	defer db.mux.Unlock()
	defer database.Closedb()
	return teamPesca, result.Error
}

func (db *OpenConnection) DeleteTeamPesca(id uint) (bool, error) {
	db.mux.Lock()
	var teamPesca entities.TeamPesca
	result := db.connection.Where(utils.DB_EQUAL_ID, id).Delete(&teamPesca)
	defer db.mux.Unlock()
	defer database.Closedb()
	return true, result.Error
}
func (db *OpenConnection) GetTeamPescaFindByName(id uint, name string) (bool, error) {
	db.mux.Lock()
	var teamPesca entities.TeamPesca
	query := db.connection.Where(utils.DB_NAME_EQUAL, name)

	if id > 0 {
		query = query.Where(utils.DB_DIFF_ID, id)
	}
	query = query.Find(&teamPesca)

	defer db.mux.Unlock()
	defer database.Closedb()
	if query.RowsAffected == 0 {
		return false, query.Error
	}
	return true, query.Error

}
