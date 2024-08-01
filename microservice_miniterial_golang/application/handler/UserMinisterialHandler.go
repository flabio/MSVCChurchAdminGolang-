package handler

import (
	"github.com/gofiber/fiber/v2"
	"microservice_ministerial.com/usecase/services"
)

type userMinisterialHandler struct {
	userMinisterialHandler services.IUserMinisterialService
}

func NewUserMinisterialHandler() IUserMinisterialHandler {
	return &userMinisterialHandler{userMinisterialHandler: services.NewUserMinistrialService()}
}

func (h *userMinisterialHandler) GetMinisterialAndUserByIdFindAll(c *fiber.Ctx) error {
	return h.userMinisterialHandler.GetMinisterialAndUserByIdFindAll(c)
}
func (h *userMinisterialHandler) AddUserToMinisterial(c *fiber.Ctx) error {
	return h.userMinisterialHandler.AddUserToMinisterial(c)
}
func (h *userMinisterialHandler) DeleteUserMinisterial(c *fiber.Ctx) error {
	return h.userMinisterialHandler.DeleteUserMinisterial(c)
}
