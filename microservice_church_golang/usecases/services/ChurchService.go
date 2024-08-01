package services

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/ulule/deepcopier"
	"microservice_church.com/core/interfaces"
	"microservice_church.com/core/repositories"
	"microservice_church.com/infrastructure/entities"
	"microservice_church.com/infrastructure/utils"

	"microservice_church.com/usecases/dto"
)

type churchService struct {
	IChurch interfaces.IChurch
}

func NewChurchService() IChurchService {
	return &churchService{
		IChurch: repositories.GetChurchInstance(),
	}
}

func (s *churchService) GetChurchFindAll(c *fiber.Ctx) error {
	result, err := s.IChurch.GetChurchFindAll()
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
func (s *churchService) GetChurchFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.IChurch.GetChurchFindById(uint(id))
	if result.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(result)
}
func (s *churchService) CreateChurch(c *fiber.Ctx) error {
	var churchCreate entities.Church
	churchDto := new(dto.ChurchDTO)
	msgError := validateChurch(0, churchDto, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(churchDto).To(&churchCreate)
	result, err := s.IChurch.CreateChurch(churchCreate)
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

func (s *churchService) UpdateChurch(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	var churchUpdate entities.Church
	churchDto := new(dto.ChurchDTO)

	findById, _ := s.IChurch.GetChurchFindById(uint(id))
	if findById.Id == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}

	msgError := validateChurch(uint(id), churchDto, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}

	deepcopier.Copy(churchDto).To(&churchUpdate)

	result, err := s.IChurch.UpdateChurch(uint(id), churchUpdate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		utils.STATUS:  http.StatusCreated,
		utils.MESSAGE: utils.UPDATED,
		utils.DATA:    result,
	})
}

func (s *churchService) DeleteChurch(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))

	findById, _ := s.IChurch.GetChurchFindById(uint(id))
	if findById.Id == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	result, err := s.IChurch.DeleteChurch(uint(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.MESSAGE: utils.REMOVED,
		utils.DATA:    result,
	})
}

func validateChurch(id uint, churchDto *dto.ChurchDTO, s *churchService, c *fiber.Ctx) string {
	var msg string = ""
	if err := c.BodyParser(churchDto); err != nil {
		msg = err.Error()
	}
	v := validate.Struct(churchDto)
	if !v.Validate() {
		msg = v.Errors.Error()
	}
	existName, _ := s.IChurch.GetChurchFindByName(id, churchDto.Name)
	if existName {
		msg = utils.NAME_ALREADY_EXIST
	}
	existEmail, _ := s.IChurch.GetChurchFindByEmail(id, churchDto.Email)
	if existEmail {
		msg = utils.EMAIL_ALREADY_EXIST
	}
	return msg
}
