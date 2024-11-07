package user

import (
	"github.com/gofiber/fiber/v2"
)

func NewUserRouter(app *fiber.App) {

	app.Get("/user/", func(c *fiber.Ctx) error {
		return MsvcUser(c)
	})
	app.Get("/user/:id", func(c *fiber.Ctx) error {
		return MsvcUser(c)
	})
	app.Post("/user", func(c *fiber.Ctx) error {
		return MsvcUser(c)
	})
	app.Put("/user/:id", func(c *fiber.Ctx) error {
		return MsvcUser(c)
	})
	app.Delete("/user/:id", func(c *fiber.Ctx) error {
		return MsvcUser(c)
	})

}
