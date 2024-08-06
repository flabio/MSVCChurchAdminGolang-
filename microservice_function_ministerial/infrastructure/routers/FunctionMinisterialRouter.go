package routers

import (
	"msvc_function_ministerial/application/handler"

	"github.com/gofiber/fiber/v2"
)

var (
	functionMinisterial handler.IFunctionMinisterialHandler = handler.NewFunctionMinisterialHandler()
)



func NewRouter(app *fiber.App) {

	api := app.Group("/api/function_ministerial")
	api.Get("/", func(c *fiber.Ctx) error {
		return functionMinisterial.GetFindAll(c)
	})
	api.Get("/:id", func(c *fiber.Ctx) error {
		return functionMinisterial.GetFindById(c)
	})
	api.Post("/", func(c *fiber.Ctx) error {
		return functionMinisterial.Create(c)
	})
	api.Put("/:id", func(c *fiber.Ctx) error {
		return functionMinisterial.Update(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return functionMinisterial.Delete(c)
	})

}
