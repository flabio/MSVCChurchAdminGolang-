package services

import (
	"github.com/gofiber/fiber/v2"
)

type IFunctionMinisterialService interface {
	GetFindAll(c *fiber.Ctx) error
	GetFindById(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
