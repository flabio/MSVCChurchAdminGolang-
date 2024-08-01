package handler

import (
	"github.com/gofiber/fiber/v2"
	"microservice_ministerial.com/usecase/services"
)

type ministerialHandler struct {
	ministerialHandler services.IMinisterialService
}

func NewMinisterialHandler() IMinisterialHandler {
	return &ministerialHandler{ministerialHandler: services.NewMinistrialService()}
}
func (h *ministerialHandler) GetMinisterialFindAll(c *fiber.Ctx) error {
	return h.ministerialHandler.GetMinisterialAll(c)
}

func (h *ministerialHandler) GetMinisterialFindById(c *fiber.Ctx) error {
	return h.ministerialHandler.GetMinisterialFindById(c)
}

func (h *ministerialHandler) CreateMinisterial(c *fiber.Ctx) error {
	return h.ministerialHandler.CreateMinisterial(c)
}

func (h *ministerialHandler) UpdateMinisterial(c *fiber.Ctx) error {
	return h.ministerialHandler.UpdateMinisterial(c)
}
func (h *ministerialHandler) DeleteMinisterial(c *fiber.Ctx) error {
	return h.ministerialHandler.DeleteMinisterial(c)
}
