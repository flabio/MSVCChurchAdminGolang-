package handler

import "github.com/gofiber/fiber/v2"

type IChurchHandler interface {
	GetChurchFindAll(c *fiber.Ctx) error
	GetChurchFindById(c *fiber.Ctx) error
	CreateChurch(c *fiber.Ctx) error
	UpdateChurch(c *fiber.Ctx) error
	DeleteChurch(c *fiber.Ctx) error
}
