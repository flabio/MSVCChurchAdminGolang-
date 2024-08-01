package handler

import (
	"github.com/gofiber/fiber/v2"
)

type ITeamPescaHandler interface {
	GetTeamPescaFindAll(c *fiber.Ctx) error
	GetTeamPescaFindById(c *fiber.Ctx) error
	CreateTeamPesca(c *fiber.Ctx) error
	UpdateTeamPesca(c *fiber.Ctx) error
	DeleteTeamPesca(c *fiber.Ctx) error
}
