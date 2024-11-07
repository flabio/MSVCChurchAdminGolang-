package ministerial

import "github.com/gofiber/fiber/v2"

func NewMinisterialRouter(app *fiber.App) {
	app.Get("/function_ministerial/", func(c *fiber.Ctx) error {
		return MsvcMinisterial(c)
	})
	app.Get("/function_ministerial/:id", func(c *fiber.Ctx) error {
		return MsvcMinisterial(c)
	})

	app.Post("/function_ministerial", func(c *fiber.Ctx) error {
		return MsvcMinisterial(c)
	})
	app.Put("/function_ministerial/:id", func(c *fiber.Ctx) error {
		return MsvcMinisterial(c)
	})
	app.Delete("/function_ministerial/:id", func(c *fiber.Ctx) error {
		return MsvcMinisterial(c)
	})

	app.Get("/function_ministerial/user/:id", func(c *fiber.Ctx) error {
		return MsvcMinisterialUser(c)
	})
	app.Post("/function_ministerial/user", func(c *fiber.Ctx) error {
		return MsvcMinisterialUser(c)
	})
	app.Delete("/function_ministerial/user/:id", func(c *fiber.Ctx) error {
		return MsvcMinisterialUser(c)
	})
}
