package routers

import (
	"github.com/gofiber/fiber/v2"
	"microservice_team_pesca.com/application/handler"
)

var (
	teamPescaHandler handler.ITeamPescaHandler = handler.NewTeamPescaHandler()
)

func NewRouter(app *fiber.App) {
	apiGroup := app.Group("/api/team_pesca")
	apiGroup.Get("/", func(c *fiber.Ctx) error {
		return teamPescaHandler.GetTeamPescaFindAll(c)
	})
	apiGroup.Get("/:id", func(c *fiber.Ctx) error {
		return teamPescaHandler.GetTeamPescaFindById(c)
	})
	apiGroup.Post("/", func(c *fiber.Ctx) error {
		return teamPescaHandler.CreateTeamPesca(c)
	})
	apiGroup.Put("/:id", func(c *fiber.Ctx) error {
		return teamPescaHandler.UpdateTeamPesca(c)
	})
	apiGroup.Delete("/:id", func(c *fiber.Ctx) error {
		return teamPescaHandler.DeleteTeamPesca(c)
	})
}
