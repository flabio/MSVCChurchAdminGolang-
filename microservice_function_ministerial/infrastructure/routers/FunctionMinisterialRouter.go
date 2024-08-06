package routers

import (
	"msvc_function_ministerial/application/handler"

	"github.com/gofiber/fiber/v2"
)

var (
	functionMinisterial     handler.IFunctionMinisterialHandler     = handler.NewFunctionMinisterialHandler()
	userFunctionMinisterial handler.IUserFunctionMinisterialHandler = handler.NewUserFunctionMinisterialHandler()
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

	userGroupFunctionMinisterial := app.Group("/api/function_ministerial/user")
	userGroupFunctionMinisterial.Get("/:id", func(c *fiber.Ctx) error {
		return userFunctionMinisterial.GetFunctionMinisterialAndUserByIdFindAll(c)
	})
	userGroupFunctionMinisterial.Post("/", func(c *fiber.Ctx) error {
		return userFunctionMinisterial.AddUserToFunctionMinisterial(c)
	})
	userGroupFunctionMinisterial.Delete("/:id", func(c *fiber.Ctx) error {
		return userFunctionMinisterial.DeleteUserFunctionMinisterial(c)
	})
}
