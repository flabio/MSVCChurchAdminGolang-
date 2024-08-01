package handler

import (
	"github.com/gofiber/fiber/v2"
	"microservice_church.com/usecases/services"
)

type churchHandler struct {
	churchService services.IChurchService
}

func NewChurchHandler() IChurchHandler {
	return &churchHandler{
		churchService: services.NewChurchService(),
	}
}

func (h *churchHandler) GetChurchFindAll(c *fiber.Ctx) error {
	return h.churchService.GetChurchFindAll(c)
}
func (h *churchHandler) GetChurchFindById(c *fiber.Ctx) error {
	return h.churchService.GetChurchFindById(c)
}
func (h *churchHandler) CreateChurch(c *fiber.Ctx) error {
	return h.churchService.CreateChurch(c)
}
func (h *churchHandler) UpdateChurch(c *fiber.Ctx) error {
	return h.churchService.UpdateChurch(c)
}
func (h *churchHandler) DeleteChurch(c *fiber.Ctx) error {
	return h.churchService.DeleteChurch(c)
}
