package services

import "github.com/gofiber/fiber/v2"

type IUserService interface {
	GetUserFindAll(c *fiber.Ctx) error
	GetUsersMembersFindAll(c *fiber.Ctx) error
	GetUserFindById(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}
