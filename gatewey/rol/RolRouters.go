package rol

import (
	"github.com/gofiber/fiber/v2"
)

func NewRolRouter(app *fiber.App) {

	app.Get("/rol/", func(c *fiber.Ctx) error {
		return MsvcRol(c)
	})
	app.Get("/rol/:id", func(c *fiber.Ctx) error {
		return MsvcRol(c)
	})
	app.Post("/rol", func(c *fiber.Ctx) error {
		return MsvcRol(c)
	})
	app.Put("/rol/:id", func(c *fiber.Ctx) error {
		return MsvcRol(c)
	})
	app.Delete("/rol/:id", func(c *fiber.Ctx) error {
		return MsvcRol(c)
	})
}
