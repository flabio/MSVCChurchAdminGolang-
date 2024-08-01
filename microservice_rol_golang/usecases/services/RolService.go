package services

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/ulule/deepcopier"
	"microservice_rol.com/core/interfaces"
	"microservice_rol.com/core/repositories"
	"microservice_rol.com/infrastructure/entities"
	"microservice_rol.com/infrastructure/utils"
	"microservice_rol.com/usecases/dto"
)

type rolService struct {
	Irol interfaces.IRol
}

func NewRolService() IRolService {
	return &rolService{
		Irol: repositories.GetRolInstance(),
	}
}

func (rolService *rolService) GetFindAll(c *fiber.Ctx) error {
	result, err := rolService.Irol.GetFindAll()
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
func (rolService *rolService) GetFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := rolService.Irol.GetFindById(id)

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
func (rolService *rolService) Create(c *fiber.Ctx) error {
	var rolCreate entities.Rol
	rol := new(dto.RolDTO)

	msgError := validateRol(0, rol, rolService, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(rol).To(&rolCreate)
	result, err := rolService.Irol.Create(rolCreate)
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
func (rolService *rolService) Update(c *fiber.Ctx) error {

	var rolEntity entities.Rol
	rolDto := new(dto.RolDTO)

	id, _ := strconv.Atoi(c.Params(utils.ID))
	rol, _ := rolService.Irol.GetFindById(id)
	if rol.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}

	msgError := validateRol(rol.Id, rolDto, rolService, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}

	deepcopier.Copy(rolDto).To(&rolEntity)
	result, err := rolService.Irol.Update(rol.Id, rolEntity)

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

func (rolService *rolService) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	rol, _ := rolService.Irol.GetFindById(id)
	if rol.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	result, err := rolService.Irol.Delete(rol.Id)
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
func validateRol(id uint, rolDto *dto.RolDTO, rolService *rolService, c *fiber.Ctx) string {
	var msg string = ""
	if err := c.BodyParser(rolDto); err != nil {
		msg = err.Error()
	}
	v := validate.Struct(rolDto)
	if !v.Validate() {
		msg = v.Errors.Error()
	}
	existName, _ := rolService.Irol.GetFindByName(id, rolDto.Name)
	if existName {
		msg = utils.NAME_ALREADY_EXIST
	}
	return msg
}
