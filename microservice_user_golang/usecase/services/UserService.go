package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ulule/deepcopier"
	"microservice_user.com/core/interfaces"
	"microservice_user.com/core/repsitories"
	"microservice_user.com/infrastructure/entities"
	"microservice_user.com/infrastructure/utils"
	"microservice_user.com/usecase/client"
	"microservice_user.com/usecase/client/church"
	"microservice_user.com/usecase/client/functionministerial"
	"microservice_user.com/usecase/client/ministerial"

	"microservice_user.com/usecase/client/rol"
	"microservice_user.com/usecase/client/teams"
	"microservice_user.com/usecase/dto"
)

type userService struct {
	IUser interfaces.IUser
}

func NewUserService() IUserService {
	return &userService{IUser: repsitories.UserInstance()}
}

func (s *userService) GetUserFindAll(c *fiber.Ctx) error {
	result, err := s.IUser.GetUserFindAll()

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   result,
	})
}
func (s *userService) GetUsersMembersFindAll(c *fiber.Ctx) error {
	result, err := s.IUser.GetUsersMembersFindAll()
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: err.Error(),
		})
	}
	var userR []dto.UserReponse
	for _, item := range result {
		var userResponseDate dto.UserReponse
		var churchs []church.Church
		var teamsPesca []teams.Team
		var rols []rol.Rol
		userResponseDate.Id = item.Id
		userResponseDate.FirstName = item.FirstName
		userResponseDate.LastName = item.LastName
		userResponseDate.Email = item.Email
		userResponseDate.Phone = item.Phone
		dataRol := rol.MsvcRolFindById(item.RolId)
		rols = append(rols, dataRol)
		userResponseDate.Rol = rols
		dataChurch := church.MsvcChurchFindById(item.ChurchId)
		churchs = append(churchs, dataChurch)
		userResponseDate.Churchs = churchs
		dataTeam := teams.MsvcTeamFindById(item.TeamPescaId)
		teamsPesca = append(teamsPesca, dataTeam)
		userResponseDate.Team = teamsPesca
		dataMinisterials := ministerial.MsvcUserMInisterialFindById(item.Id)
		userResponseDate.UserMinisterial = dataMinisterials
		dataFunctionMinisterials := functionministerial.MsvcUserFunctionMInisterialFindById(item.Id)
		userResponseDate.FunctionMinisterial = dataFunctionMinisterials
		userR = append(userR, userResponseDate)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   userR,
	})
}
func (s *userService) GetUserFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.IUser.GetUserFindById(uint(id))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
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
func (s *userService) CreateUser(c *fiber.Ctx) error {
	var user entities.User
	var userDTO dto.UserDTO
	userCreate, msgError := validateUser(0, userDTO, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	msgMsvcRol := rol.MsvcRolById(userDTO.RolId)
	if msgMsvcRol != "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			utils.STATUS:  http.StatusInternalServerError,
			utils.MESSAGE: msgMsvcRol,
		})
	}
	msgMsvcChurch := client.MsvcChurchById(uint(userDTO.ChurchId))
	if msgMsvcChurch != "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			utils.STATUS:  http.StatusInternalServerError,
			utils.MESSAGE: msgMsvcChurch,
		})
	}
	deepcopier.Copy(userCreate).To(&user)
	result, err := s.IUser.CreateUser(user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			utils.STATUS:  http.StatusInternalServerError,
			utils.MESSAGE: utils.INTERNAL_SERVER_ERROR,
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		utils.STATUS:  http.StatusCreated,
		utils.MESSAGE: utils.CREATED,
		utils.DATA:    result,
	})
}

func (s *userService) UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	var user entities.User
	var userDTO dto.UserDTO
	findById, _ := s.IUser.GetUserFindById(uint(id))
	if findById.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	userUpdate, msgError := validateUser(uint(id), userDTO, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(userUpdate).To(&user)
	result, err := s.IUser.UpdateUser(uint(id), user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			utils.STATUS:  http.StatusInternalServerError,
			utils.MESSAGE: utils.INTERNAL_SERVER_ERROR,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.MESSAGE: utils.UPDATED,
		utils.DATA:    result,
	})
}
func (s *userService) DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	findById, _ := s.IUser.GetUserFindById(uint(id))
	if findById.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	result, err := s.IUser.DeleteUser(uint(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			utils.STATUS:  http.StatusInternalServerError,
			utils.MESSAGE: utils.INTERNAL_SERVER_ERROR,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.MESSAGE: utils.REMOVED,
		utils.DATA:    result,
	})
}

func validateUser(id uint, userDTO dto.UserDTO, s *userService, c *fiber.Ctx) (dto.UserDTO, string) {
	var msg string = ""
	b := c.Body()
	var dataMap map[string]interface{}
	errJson := json.Unmarshal([]byte(b), &dataMap)
	if errJson != nil {
		msg = errJson.Error()
	}
	msgValid := validateField(dataMap)
	if msgValid != "" {
		return dto.UserDTO{}, msgValid
	}
	MapToStruct(&userDTO, dataMap)
	msgRequired := validateRequired(userDTO)
	if msgRequired != "" {
		return dto.UserDTO{}, msgRequired
	}
	existIdentification, _ := s.IUser.GetUserFindByIdentification(id, userDTO.Identification)
	if existIdentification {
		msg = utils.IDENTIFICATION_ALREADY_EXIST
	}
	existEmail, _ := s.IUser.GetUserFindByEmail(id, userDTO.Email)
	if existEmail {
		msg = utils.EMAIL_ALREADY_EXIST
	}
	return userDTO, msg
}

func MapToStruct(dataDto *dto.UserDTO, dataMap map[string]interface{}) {
	user := dto.UserDTO{
		FirstName:          dataMap["first_name"].(string),
		LastName:           dataMap["last_name"].(string),
		Identification:     dataMap["identification"].(string),
		Address:            dataMap["address"].(string),
		Phone:              dataMap["phone"].(string),
		TypeIdentification: dataMap["type_identification"].(string),
	}
	*dataDto = user
}
func validateField(value map[string]interface{}) string {
	var msg string = ""
	if value["first_name"] == nil {
		msg = "The field name is required"
	}
	if value["last_name"] == nil {
		msg = "The field last name is required"
	}
	if value["identification"] == nil {
		msg = "The field identification is required"
	}
	if value["address"] == nil {
		msg = "The field address is required"
	}
	if value["phone"] == nil {
		msg = "The field phone is required"
	}
	if value["type_identification"] == nil {
		msg = "The field type identification is required"
	}
	return msg
}
func validateRequired(field dto.UserDTO) string {
	var msg string = ""
	if field.FirstName == "" {
		msg = "The first name is required"
	}
	if field.LastName == "" {
		msg = "The last name is required"
	}
	if field.Identification == "" {
		msg = "The identification is required"
	}
	if field.Address == "" {
		msg = "The address is required"
	}
	if field.Phone == "" {
		msg = "The phone is required"
	}
	if field.TypeIdentification == "" {
		msg = "The sex is required"
	}
	return msg
}
