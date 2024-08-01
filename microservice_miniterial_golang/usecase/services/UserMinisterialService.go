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

type userMinitsrialService struct {
	IuserMinisterial interfaces.IUserMinisterial
}

func NewUserMinistrialService() IUserMinisterialService {
	return &userMinitsrialService{IuserMinisterial: repository.UserMinisterialInstance()}
}
func (s *userMinitsrialService) GetMinisterialAndUserByIdFindAll(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))

	results, err := s.IuserMinisterial.GetMinisterialAndUserByIdFindAll(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(results)

}

func (s *userMinitsrialService) AddUserToMinisterial(c *fiber.Ctx) error {
	var ministrial entities.UserMinisterial
	ministerialDTO := new(dto.UserMinisterialDTO)

	msgErr := validateUserMInisterial(ministerialDTO, c)
	if msgErr != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgErr,
		})
	}
	deepcopier.Copy(ministerialDTO).To(&ministrial)
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

func validateUserMInisterial(userMinisterialDTO *dto.UserMinisterialDTO, c *fiber.Ctx) string {
	var msg string = ""
	if err := c.BodyParser(userMinisterialDTO); err != nil {
		msg = err.Error()
	}
	v := validate.Struct(userMinisterialDTO)
	if !v.Validate() {
		msg = v.Errors.Error()
	}
	return msg
}
