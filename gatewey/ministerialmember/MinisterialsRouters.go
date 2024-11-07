package ministerialmember

import "github.com/gofiber/fiber/v2"

func NewMinisterialMemberRouter(app *fiber.App) {
	app.Get("/ministerial/", func(c *fiber.Ctx) error {
		return MsvcMinisterialMember(c)
	})
	app.Get("/ministerial/:id", func(c *fiber.Ctx) error {
		return MsvcMinisterialMember(c)
	})

	app.Post("/ministerial", func(c *fiber.Ctx) error {
		return MsvcMinisterialMember(c)
	})
	app.Put("/ministerial/:id", func(c *fiber.Ctx) error {
		return MsvcMinisterialMember(c)
	})
	app.Delete("/ministerial/:id", func(c *fiber.Ctx) error {
		return MsvcMinisterialMember(c)
	})

	app.Get("/ministerial/user_ministerial/:id", func(c *fiber.Ctx) error {
		return MsvcMinisterialUserMinisterial(c)
	})
	app.Post("/ministerial/user_ministerial", func(c *fiber.Ctx) error {
		return MsvcMinisterialUserMinisterial(c)
	})
	app.Delete("/ministerial/user_ministerial/:id", func(c *fiber.Ctx) error {
		return MsvcMinisterialUserMinisterial(c)
	})
}
