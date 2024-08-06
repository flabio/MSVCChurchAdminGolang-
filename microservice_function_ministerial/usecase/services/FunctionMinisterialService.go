package services

import (
	"msvc_function_ministerial/core/interfaces"
	"msvc_function_ministerial/core/repositories"
	"msvc_function_ministerial/infrastructure/entities"
	"msvc_function_ministerial/infrastructure/utils"
	"msvc_function_ministerial/usecase/dto"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/ulule/deepcopier"
)

type functionMinisterialService struct {
	IfunctionMinisterial interfaces.IFunctionMinisterial
}

func NewFunctionMinisterialService() IFunctionMinisterialService {
	return &functionMinisterialService{
		IfunctionMinisterial: repositories.GetFunctionMinisterialInstance(),
	}
}

func (s *functionMinisterialService) GetFindAll(c *fiber.Ctx) error {
	result, err := s.IfunctionMinisterial.GetFindAll()
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
func (s *functionMinisterialService) GetFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.IfunctionMinisterial.GetFindById(id)

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
func (s *functionMinisterialService) Create(c *fiber.Ctx) error {
	var create entities.FunctionMinisterial
	functionMinisterial := new(dto.FunctionMinisterialDTO)

	msgError := validateRol(0, functionMinisterial, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(functionMinisterial).To(&create)
	result, err := s.IfunctionMinisterial.Create(create)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.MESSAGE: utils.CREATED,
		utils.DATA:    result,
	})
}
func (s *functionMinisterialService) Update(c *fiber.Ctx) error {

	var functionMinisterialEntity entities.FunctionMinisterial
	functionMinisterialDto := new(dto.FunctionMinisterialDTO)

	id, _ := strconv.Atoi(c.Params(utils.ID))
	functionMinisterial, _ := s.IfunctionMinisterial.GetFindById(id)
	if functionMinisterial.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}

	msgError := validateRol(functionMinisterial.Id, functionMinisterialDto, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}

	deepcopier.Copy(functionMinisterialDto).To(&functionMinisterialEntity)
	result, err := s.IfunctionMinisterial.Update(functionMinisterial.Id, functionMinisterialEntity)

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

func (s *functionMinisterialService) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	functionMinisterial, _ := s.IfunctionMinisterial.GetFindById(id)
	if functionMinisterial.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	result, err := s.IfunctionMinisterial.Delete(functionMinisterial.Id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		utils.STATUS:  http.StatusCreated,
		utils.MESSAGE: utils.REMOVED,
		utils.DATA:    result,
	})
}
func validateRol(id uint, functionMinisterialDto *dto.FunctionMinisterialDTO, s *functionMinisterialService, c *fiber.Ctx) string {
	var msg string = ""
	if err := c.BodyParser(functionMinisterialDto); err != nil {
		msg = err.Error()
	}
	v := validate.Struct(functionMinisterialDto)
	if !v.Validate() {
		msg = v.Errors.Error()
	}
	existName, _ := s.IfunctionMinisterial.GetFindByName(id, functionMinisterialDto.Name)
	if existName {
		msg = utils.NAME_ALREADY_EXIST
	}
	return msg
}
