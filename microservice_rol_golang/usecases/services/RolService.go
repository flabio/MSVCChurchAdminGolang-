package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
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
	var rol dto.RolDTO

	rolDto, msgError := validateRol(0, rol, rolService, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(rolDto).To(&rolCreate)
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
	var rolDto dto.RolDTO
	id, _ := strconv.Atoi(c.Params(utils.ID))
	rol, _ := rolService.Irol.GetFindById(id)
	if rol.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	rolUpdate, msgError := validateRol(rol.Id, rolDto, rolService, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(rolUpdate).To(&rolEntity)
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
func validateRol(id uint, rolDto dto.RolDTO, s *rolService, c *fiber.Ctx) (dto.RolDTO, string) {
	var msg string = ""
	b := c.Body()
	var dataMap map[string]interface{}
	errJson := json.Unmarshal([]byte(b), &dataMap)
	if errJson != nil {
		msg = errJson.Error()
	}
	msgValid := validateField(dataMap)
	if msgValid != "" {
		return dto.RolDTO{}, msgValid
	}
	MapToStruct(&rolDto, dataMap)
	msgRequired := validateRequired(rolDto)
	if msgRequired != "" {
		return dto.RolDTO{}, msgRequired
	}
	existName, _ := s.Irol.GetFindByName(id, rolDto.Name)
	if existName {
		msg = utils.NAME_ALREADY_EXIST
	}
	return rolDto, msg
}

func MapToStruct(dataDto *dto.RolDTO, dataMap map[string]interface{}) {
	rol := dto.RolDTO{
		Name:   dataMap["name"].(string),
		Active: dataMap["active"].(bool),
	}
	*dataDto = rol
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
func validateRequired(field dto.RolDTO) string {
	var msg string = ""
	if field.Name == "" {
		msg = "The name is required"
	}
	return msg
}
