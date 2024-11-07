package services

import (
	"encoding/json"
	"msvc_function_ministerial/core/interfaces"
	"msvc_function_ministerial/core/repositories"
	"msvc_function_ministerial/infrastructure/entities"
	"msvc_function_ministerial/infrastructure/utils"
	"msvc_function_ministerial/usecase/dto"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ulule/deepcopier"
)

type userFunctionMinitsrialService struct {
	IuserFunctionMinisterial interfaces.IUserFunctionMinisterial
}

func NewUserFunctionMinistrialService() IUserFunctionMinisterialService {
	return &userFunctionMinitsrialService{IuserFunctionMinisterial: repositories.UserFunctionMinisterialInstance()}
}
func (s *userFunctionMinitsrialService) GetFunctionMinisterialAndUserByIdFindAll(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.IuserFunctionMinisterial.GetFunctionMinisterialAndUserByIdFindAll(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(result)
}
func (s *userFunctionMinitsrialService) AddUserToFunctionMinisterial(c *fiber.Ctx) error {
	var ministrial entities.UserFunctionMinisterial
	var ministerialDTO dto.UserFunctionMinisterialDTO
	data, msgErr := validateUserMInisterial(ministerialDTO, c)
	if msgErr != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgErr,
		})
	}
	deepcopier.Copy(data).To(&ministrial)
	_, err := s.IuserFunctionMinisterial.AddUserToFunctionMinisterial(ministrial)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			utils.STATUS:  http.StatusInternalServerError,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		utils.STATUS:  http.StatusCreated,
		utils.MESSAGE: utils.CREATED,
	})
}

func (s *userFunctionMinitsrialService) DeleteUserFunctionMinisterial(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.IuserFunctionMinisterial.DeleteUserFunctionMinisterial(uint(id))
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

func validateUserMInisterial(userMinisterialDTO dto.UserFunctionMinisterialDTO, c *fiber.Ctx) (dto.UserFunctionMinisterialDTO, string) {
	var msg string = ""
	b := c.Body()
	var dataMap map[string]interface{}
	errJson := json.Unmarshal([]byte(b), &dataMap)
	if errJson != nil {
		msg = errJson.Error()
	}
	msgValid := validateField(dataMap)
	if msgValid != "" {
		return dto.UserFunctionMinisterialDTO{}, msgValid
	}
	MapToStruct(&userMinisterialDTO, dataMap)
	msgRequired := validateRequired(userMinisterialDTO)
	if msgRequired != "" {
		return dto.UserFunctionMinisterialDTO{}, msgRequired
	}
	return userMinisterialDTO, msg
}

func MapToStruct(dataDto *dto.UserFunctionMinisterialDTO, dataMap map[string]interface{}) {
	fiels := dto.UserFunctionMinisterialDTO{
		FunctionMinisterialId: dataMap["function_ministerial_id"].(uint),
		UserId:                dataMap["user_id"].(uint),
	}
	*dataDto = fiels
}
func validateField(value map[string]interface{}) string {
	var msg string = ""
	if value["function_ministerial_id"] == nil {
		msg = "The field function_ministerial_id is required"
	}
	if value["user_id"] == nil {
		msg = "The field user id is required"
	}
	return msg
}
func validateRequired(field dto.UserFunctionMinisterialDTO) string {
	var msg string = ""
	if field.FunctionMinisterialId == 0 {
		msg = "The FunctionMinisterialId is required"
	}
	if field.UserId == 0 {
		msg = "The UserId is required"
	}
	return msg
}
