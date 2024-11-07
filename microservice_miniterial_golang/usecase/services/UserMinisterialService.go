package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ulule/deepcopier"
	"microservice_ministerial.com/core/interfaces"
	"microservice_ministerial.com/core/repository"
	"microservice_ministerial.com/infrastructure/entities"
	"microservice_ministerial.com/infrastructure/utils"
	"microservice_ministerial.com/usecase/dto"
)

type userMinitsrialService struct {
	IuserMinisterial interfaces.IUserMinisterial
}

func NewUserMinistrialService() IUserMinisterialService {
	return &userMinitsrialService{IuserMinisterial: repository.UserMinisterialInstance()}
}
func (s *userMinitsrialService) GetMinisterialAndUserByIdFindAll(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.IuserMinisterial.GetMinisterialAndUserByIdFindAll(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(result)
}
func (s *userMinitsrialService) AddUserToMinisterial(c *fiber.Ctx) error {
	var ministrial entities.UserMinisterial
	var ministerialDTO dto.UserMinisterialDTO
	data, msgErr := validateUserMInisterial(ministerialDTO, c)
	if msgErr != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgErr,
		})
	}
	deepcopier.Copy(data).To(&ministrial)
	newMinistrial, err := s.IuserMinisterial.AddUserToMinisterial(ministrial)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			utils.STATUS:  http.StatusInternalServerError,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		utils.STATUS:  http.StatusCreated,
		utils.MESSAGE: utils.CREATED,
		utils.DATA:    newMinistrial,
	})
}

func (s *userMinitsrialService) DeleteUserMinisterial(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))

	result, err := s.IuserMinisterial.DeleteUserMinisterial(uint(id))
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

func validateUserMInisterial(userMinisterialDTO dto.UserMinisterialDTO, c *fiber.Ctx) (dto.UserMinisterialDTO, string) {
	var msg string = ""
	b := c.Body()
	var dataMap map[string]interface{}
	errJson := json.Unmarshal([]byte(b), &dataMap)
	if errJson != nil {
		msg = errJson.Error()
	}
	msgValid := validateFieldUserMinisterial(dataMap)
	if msgValid != "" {
		return dto.UserMinisterialDTO{}, msgValid
	}
	MapToStructUserMinisterial(&userMinisterialDTO, dataMap)
	msgRequired := validateRequiredUserMinisterial(userMinisterialDTO)
	if msgRequired != "" {
		return dto.UserMinisterialDTO{}, msgRequired
	}
	return userMinisterialDTO, msg
}
func MapToStructUserMinisterial(dataDto *dto.UserMinisterialDTO, dataMap map[string]interface{}) {
	fields := dto.UserMinisterialDTO{
		MinisterialId: dataMap["ministerial_id"].(uint),
		UserId:        dataMap["user_id"].(uint),
	}
	*dataDto = fields
}
func validateFieldUserMinisterial(value map[string]interface{}) string {
	var msg string = ""
	if value["ministerial_id"] == nil {
		msg = "The field ministerial_id is required"
	}
	if value["user_id"] == nil {
		msg = "The field user_id is required"
	}
	return msg
}
func validateRequiredUserMinisterial(field dto.UserMinisterialDTO) string {
	var msg string = ""
	if field.MinisterialId == 0 {
		msg = "The ministerial id is required"
	}
	if field.UserId == 0 {
		msg = "The user id is required"
	}
	return msg
}
