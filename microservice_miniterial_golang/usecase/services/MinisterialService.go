package services

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/ulule/deepcopier"
	"microservice_ministerial.com/core/interfaces"
	"microservice_ministerial.com/core/repository"
	"microservice_ministerial.com/infrastructure/entities"
	"microservice_ministerial.com/infrastructure/utils"
	"microservice_ministerial.com/usecase/dto"
)

type minitsrialService struct {
	Iministerial interfaces.IMinisterial
}

func NewMinistrialService() IMinisterialService {
	return &minitsrialService{Iministerial: repository.Instance()}
}
func (s *minitsrialService) GetMinisterialAll(c *fiber.Ctx) error {

	ministrials, err := s.Iministerial.GetMinisterialAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   ministrials,
	})
}
func (s *minitsrialService) GetMinisterialFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	findById, _ := s.Iministerial.GetMinisterialFindById(uint(id))
	if findById.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   findById,
	})
}
func (s *minitsrialService) CreateMinisterial(c *fiber.Ctx) error {
	var ministrial entities.Ministerial
	ministerialDTO := new(dto.MinisterialDTO)

	msgErr := validateMInisterial(0, ministerialDTO, s, c)
	if msgErr != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgErr,
		})
	}
	deepcopier.Copy(ministerialDTO).To(&ministrial)
	newMinistrial, err := s.Iministerial.CreateMinisterial(ministrial)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			utils.STATUS:  http.StatusInternalServerError,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		utils.STATUS: http.StatusCreated,
		utils.DATA:   newMinistrial,
	})
}

func (s *minitsrialService) UpdateMinisterial(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	var ministrial entities.Ministerial
	ministrialDTO := new(dto.MinisterialDTO)

	findById, _ := s.Iministerial.GetMinisterialFindById(uint(id))
	if findById.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	msgErr := validateMInisterial(uint(id), ministrialDTO, s, c)
	if msgErr != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgErr,
		})
	}
	deepcopier.Copy(ministrialDTO).To(&ministrial)

	updatedMinistrial, err := s.Iministerial.UpdateMinisterial(uint(id), ministrial)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			utils.STATUS:  http.StatusInternalServerError,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   updatedMinistrial,
	})
}
func (s *minitsrialService) DeleteMinisterial(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	findById, _ := s.Iministerial.GetMinisterialFindById(uint(id))
	if findById.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	result, err := s.Iministerial.DeleteMinisterial(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			utils.STATUS:  http.StatusInternalServerError,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.DATA:    result,
		utils.MESSAGE: utils.REMOVED,
	})
}

func validateMInisterial(id uint, ministerialDTO *dto.MinisterialDTO, s *minitsrialService, c *fiber.Ctx) string {
	var msg string = ""
	if err := c.BodyParser(ministerialDTO); err != nil {
		msg = err.Error()
	}
	v := validate.Struct(ministerialDTO)
	if !v.Validate() {
		msg = v.Errors.Error()
	}
	existName, _ := s.Iministerial.GetMinisterialFindByName(id, ministerialDTO.Name)
	if existName {
		msg = utils.NAME_ALREADY_EXIST
	}
	return msg
}
