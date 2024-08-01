package handler

import (
	"github.com/gofiber/fiber/v2"
	"microservice_user.com/usecase/services"
)

type userHandler struct {
	user services.IUserService
}

func NewUserHandler() IUserHandler {
	return &userHandler{user: services.NewUserService()}
}

func (h *userHandler) GetUserFindAll(c *fiber.Ctx) error {
	return h.user.GetUserFindAll(c)
}
func (h *userHandler) GetUsersMembersFindAll(c *fiber.Ctx) error {
	return h.user.GetUsersMembersFindAll(c)
}
func (h *userHandler) GetUserFindById(c *fiber.Ctx) error {
	return h.user.GetUserFindById(c)
}
func (h *userHandler) CreateUser(c *fiber.Ctx) error {
	return h.user.CreateUser(c)
}
func (h *userHandler) UpdateUser(c *fiber.Ctx) error {
	return h.user.UpdateUser(c)
}

func (h *userHandler) DeleteUser(c *fiber.Ctx) error {
	return h.user.DeleteUser(c)
}
