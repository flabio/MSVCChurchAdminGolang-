package handler

import "github.com/gofiber/fiber/v2"

type IUserFunctionMinisterialHandler interface {
	GetFunctionMinisterialAndUserByIdFindAll(c *fiber.Ctx) error
	AddUserToFunctionMinisterial(c *fiber.Ctx) error
	DeleteUserFunctionMinisterial(c *fiber.Ctx) error
}
