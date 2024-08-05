package services

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/ulule/deepcopier"
	"microservice_user.com/core/interfaces"
	"microservice_user.com/core/repsitories"
	"microservice_user.com/infrastructure/entities"
	"microservice_user.com/infrastructure/utils"
	"microservice_user.com/usecase/client"
	"microservice_user.com/usecase/client/church"
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
	userDTO := new(dto.UserDTO)
	msgError := vaidateCreate(userDTO, s, c)
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

	deepcopier.Copy(userDTO).To(&user)
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
	userDTO := new(dto.UserUpdateDTO)

	findById, _ := s.IUser.GetUserFindById(uint(id))
	if findById.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}

	msgError := validateUpdate(uint(id), userDTO, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(userDTO).To(&user)
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
func vaidateCreate(userDTO *dto.UserDTO, s *userService, c *fiber.Ctx) string {
	var msg string = ""
	if err := c.BodyParser(userDTO); err != nil {
		msg = err.Error()
	}
	v := validate.Struct(userDTO)
	if !v.Validate() {
		msg = v.Errors.Error()
	}
	existIdentification, _ := s.IUser.GetUserFindByIdentification(0, userDTO.Identification)
	if existIdentification {
		msg = utils.IDENTIFICATION_ALREADY_EXIST
	}
	existEmail, _ := s.IUser.GetUserFindByEmail(0, userDTO.Email)
	if existEmail {
		msg = utils.EMAIL_ALREADY_EXIST
	}
	existUsername, _ := s.IUser.GetUserFindByUsername(0, userDTO.Username)
	if existUsername {
		msg = utils.USERNAME_ALREADY_EXIST
	}
	return msg
}
func validateUpdate(id uint, userDTO *dto.UserUpdateDTO, s *userService, c *fiber.Ctx) string {
	var msg string = ""
	if err := c.BodyParser(userDTO); err != nil {
		msg = err.Error()
	}
	v := validate.Struct(userDTO)
	if !v.Validate() {
		msg = v.Errors.Error()
	}
	existIdentification, _ := s.IUser.GetUserFindByIdentification(id, userDTO.Identification)
	if existIdentification {
		msg = utils.IDENTIFICATION_ALREADY_EXIST
	}
	existEmail, _ := s.IUser.GetUserFindByEmail(id, userDTO.Email)
	if existEmail {
		msg = utils.EMAIL_ALREADY_EXIST
	}
	return msg
}
