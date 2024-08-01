package handler

import "github.com/gofiber/fiber/v2"

type IUserMinisterialHandler interface {
	GetMinisterialAndUserByIdFindAll(c *fiber.Ctx) error
	AddUserToMinisterial(c *fiber.Ctx) error
	DeleteUserMinisterial(c *fiber.Ctx) error
}
