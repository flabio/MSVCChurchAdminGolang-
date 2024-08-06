package handler

import (
	"msvc_function_ministerial/usecase/services"

	"github.com/gofiber/fiber/v2"
)

type userFunctionMinisterialHandler struct {
	userFunctionMinisterial services.IUserFunctionMinisterialService
}

func NewUserFunctionMinisterialHandler() IUserFunctionMinisterialHandler {
	return &userFunctionMinisterialHandler{userFunctionMinisterial: services.NewUserFunctionMinistrialService()}
}

func (h *userFunctionMinisterialHandler) GetFunctionMinisterialAndUserByIdFindAll(c *fiber.Ctx) error {
	return h.userFunctionMinisterial.GetFunctionMinisterialAndUserByIdFindAll(c)
}
func (h *userFunctionMinisterialHandler) AddUserToFunctionMinisterial(c *fiber.Ctx) error {
	return h.userFunctionMinisterial.AddUserToFunctionMinisterial(c)
}
func (h *userFunctionMinisterialHandler) DeleteUserFunctionMinisterial(c *fiber.Ctx) error {
	return h.userFunctionMinisterial.DeleteUserFunctionMinisterial(c)
}
