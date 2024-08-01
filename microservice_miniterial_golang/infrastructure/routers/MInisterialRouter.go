package routers

import (
	"github.com/gofiber/fiber/v2"
	"microservice_ministerial.com/application/handler"
)

var (
	ministerialHandler     handler.IMinisterialHandler     = handler.NewMinisterialHandler()
	userMinisterialHandler handler.IUserMinisterialHandler = handler.NewUserMinisterialHandler()
)

func NewMinisterialRouter(app *fiber.App) {
	ministerial := app.Group("/api/ministerial")

	ministerial.Get("/", func(c *fiber.Ctx) error {
		return ministerialHandler.GetMinisterialFindAll(c)
	})
	ministerial.Get("/:id", func(c *fiber.Ctx) error {
		return ministerialHandler.GetMinisterialFindById(c)
	})
	ministerial.Post("/", func(c *fiber.Ctx) error {
		return ministerialHandler.CreateMinisterial(c)
	})
	ministerial.Put("/:id", func(c *fiber.Ctx) error {
		return ministerialHandler.UpdateMinisterial(c)
	})
	ministerial.Delete("/:id", func(c *fiber.Ctx) error {
		return ministerialHandler.DeleteMinisterial(c)
	})

	userMinisterial := app.Group("/api/ministerial/user_ministerial")
	userMinisterial.Get("/:id", func(c *fiber.Ctx) error {
		return userMinisterialHandler.GetMinisterialAndUserByIdFindAll(c)
	})
	userMinisterial.Post("/", func(c *fiber.Ctx) error {
		return userMinisterialHandler.AddUserToMinisterial(c)
	})
	userMinisterial.Delete("/:id", func(c *fiber.Ctx) error {
		return userMinisterialHandler.DeleteUserMinisterial(c)
	})
}
