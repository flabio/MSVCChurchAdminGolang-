package services

import "github.com/gofiber/fiber/v2"

type IUserMinisterialService interface {
	GetMinisterialAndUserByIdFindAll(c *fiber.Ctx) error
	AddUserToMinisterial(c *fiber.Ctx) error
	DeleteUserMinisterial(c *fiber.Ctx) error
}
