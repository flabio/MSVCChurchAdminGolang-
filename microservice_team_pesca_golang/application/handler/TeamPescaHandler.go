package handler

import (
	"github.com/gofiber/fiber/v2"
	"microservice_team_pesca.com/usecase/services"
)

type teamPescaHandler struct {
	teamPesca services.ITeamPescaService
}

func NewTeamPescaHandler() ITeamPescaHandler {
	return &teamPescaHandler{teamPesca: services.NewTeamPescaService()}
}

func (h *teamPescaHandler) GetTeamPescaFindAll(c *fiber.Ctx) error {
	return h.teamPesca.GetTeamPescaFindAll(c)
}
func (h *teamPescaHandler) GetTeamPescaFindById(c *fiber.Ctx) error {
	return h.teamPesca.GetTeamPescaFindById(c)
}
func (h *teamPescaHandler) CreateTeamPesca(c *fiber.Ctx) error {
	return h.teamPesca.CreateTeamPesca(c)
}
func (h *teamPescaHandler) UpdateTeamPesca(c *fiber.Ctx) error {
	return h.teamPesca.UpdateTeamPesca(c)
}

func (h *teamPescaHandler) DeleteTeamPesca(c *fiber.Ctx) error {
	return h.teamPesca.DeleteTeamPesca(c)
}
