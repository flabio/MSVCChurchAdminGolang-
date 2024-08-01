package services

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
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
	return c.Status(http.StatusOK).JSON( result)
}
func (s *teamPescaService) CreateTeamPesca(c *fiber.Ctx) error {
	var createTeamPesca entities.TeamPesca
	teamPescaDTO := new(dto.TeamPescaDTO)

	msgError := validateTeamPesca(0, teamPescaDTO, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(teamPescaDTO).To(&createTeamPesca)
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
	teamPescaDTO := new(dto.TeamPescaDTO)
	result, _ := s.ITeamPesca.GetTeamPescaFindById(uint(id))

	if result.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	msgError := validateTeamPesca(uint(id), teamPescaDTO, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(teamPescaDTO).To(&updateTeamPesca)
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

func validateTeamPesca(id uint, teamPescaDto *dto.TeamPescaDTO, s *teamPescaService, c *fiber.Ctx) string {
	var msg string = ""
	if err := c.BodyParser(teamPescaDto); err != nil {
		msg = err.Error()
	}
	v := validate.Struct(teamPescaDto)
	if !v.Validate() {
		msg = v.Errors.Error()
	}
	existName, _ := s.ITeamPesca.GetTeamPescaFindByName(id, teamPescaDto.Name)
	if existName {
		msg = utils.NAME_ALREADY_EXIST
	}
	return msg
}
