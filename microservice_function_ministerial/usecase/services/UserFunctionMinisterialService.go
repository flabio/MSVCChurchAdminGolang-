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
	ministerialDTO := new(dto.UserFunctionMinisterialDTO)

	msgErr := validateUserMInisterial(ministerialDTO, c)
	if msgErr != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgErr,
		})
	}
	deepcopier.Copy(ministerialDTO).To(&ministrial)
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

func validateUserMInisterial(userMinisterialDTO *dto.UserFunctionMinisterialDTO, c *fiber.Ctx) string {
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
