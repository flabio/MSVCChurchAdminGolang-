package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ulule/deepcopier"
	"microservice_team_pesca.com/core/interfaces"
	"microservice_team_pesca.com/core/repository"
	"microservice_team_pesca.com/infrastructure/entities"
	"microservice_team_pesca.com/infrastructure/utils"
	"microservice_team_pesca.com/usecase/dto"
)

type teamPescaService struct {
	ITeamPesca interfaces.ITeamPesca
}

func NewTeamPescaService() ITeamPescaService {
	return &teamPescaService{ITeamPesca: repository.TeamPescaInstance()}
}

func (s *teamPescaService) GetTeamPescaFindAll(c *fiber.Ctx) error {
	result, err := s.ITeamPesca.GetTeamPescaFindAll()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   result,
	})
}
func (s *teamPescaService) GetTeamPescaFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.ITeamPesca.GetTeamPescaFindById(uint(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	if result.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	return c.Status(http.StatusOK).JSON(result)
}
func (s *teamPescaService) CreateTeamPesca(c *fiber.Ctx) error {
	var createTeamPesca entities.TeamPesca
	var teamPescaDTO dto.TeamPescaDTO

	data, msgError := validateTeamPesca(0, teamPescaDTO, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(data).To(&createTeamPesca)
	result, err := s.ITeamPesca.CreateTeamPesca(createTeamPesca)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		utils.STATUS:  http.StatusCreated,
		utils.MESSAGE: utils.CREATED,
		utils.DATA:    result,
	})
}
func (s *teamPescaService) UpdateTeamPesca(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	var updateTeamPesca entities.TeamPesca
	var teamPescaDTO dto.TeamPescaDTO
	result, _ := s.ITeamPesca.GetTeamPescaFindById(uint(id))

	if result.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	data, msgError := validateTeamPesca(uint(id), teamPescaDTO, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(data).To(&updateTeamPesca)
	updateTeamPesca.Id = uint(id)
	result, err := s.ITeamPesca.UpdateTeamPesca(uint(id), updateTeamPesca)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.MESSAGE: utils.UPDATED,
		utils.DATA:    result,
	})
}
func (s *teamPescaService) DeleteTeamPesca(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, _ := s.ITeamPesca.GetTeamPescaFindById(uint(id))
	if result.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	teamPesca, err := s.ITeamPesca.DeleteTeamPesca(uint(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.MESSAGE: utils.REMOVED,
		utils.DATA:    teamPesca,
	})
}

func validateTeamPesca(id uint, teamPescaDto dto.TeamPescaDTO, s *teamPescaService, c *fiber.Ctx) (dto.TeamPescaDTO, string) {
	var msg string = ""
	b := c.Body()
	var dataMap map[string]interface{}
	errJson := json.Unmarshal([]byte(b), &dataMap)
	if errJson != nil {
		msg = errJson.Error()
	}
	msgValid := validateField(dataMap)
	if msgValid != "" {
		return dto.TeamPescaDTO{}, msgValid
	}
	MapToStruct(&teamPescaDto, dataMap)
	msgRequired := validateRequired(teamPescaDto)
	if msgRequired != "" {
		return dto.TeamPescaDTO{}, msgRequired
	}
	existName, _ := s.ITeamPesca.GetTeamPescaFindByName(id, teamPescaDto.Name)
	if existName {
		msg = utils.NAME_ALREADY_EXIST
	}
	return teamPescaDto, msg
}

func MapToStruct(dataDto *dto.TeamPescaDTO, dataMap map[string]interface{}) {
	fields := dto.TeamPescaDTO{
		Name:   dataMap["name"].(string),
		Active: dataMap["active"].(bool),
	}
	*dataDto = fields
}
func validateField(value map[string]interface{}) string {
	var msg string = ""
	if value["name"] == nil {
		msg = "The field name is required"
	}
	if value["active"] == nil {
		msg = "The field active is required"
	}
	return msg
}
func validateRequired(field dto.TeamPescaDTO) string {
	var msg string = ""
	if field.Name == "" {
		msg = "The name is required"
	}
	return msg
}
