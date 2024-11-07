package church

import (
	"github.com/gofiber/fiber/v2"
)

func NewChurchRouter(app *fiber.App) {

	app.Get("/church/", func(c *fiber.Ctx) error {
		return MsvcChurch(c)
	})
	app.Get("/church/:id", func(c *fiber.Ctx) error {
		return MsvcChurch(c)
	})
	app.Post("/church", func(c *fiber.Ctx) error {
		return MsvcChurch(c)
	})
	app.Put("/church/:id", func(c *fiber.Ctx) error {
		return MsvcChurch(c)
	})
	app.Delete("/church/:id", func(c *fiber.Ctx) error {
		return MsvcChurch(c)
	})
}
