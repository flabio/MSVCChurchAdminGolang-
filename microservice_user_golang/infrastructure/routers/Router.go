package routers

import (
	"github.com/gofiber/fiber/v2"
	"microservice_user.com/application/handler"
)

var (
	userHandler handler.IUserHandler = handler.NewUserHandler()
)

func NewRouter(app *fiber.App) {
	apiGroup := app.Group("/api/user")
	apiGroup.Get("/", func(c *fiber.Ctx) error {
		return userHandler.GetUserFindAll(c)
	})
	apiGroup.Get("/members", func(c *fiber.Ctx) error {
		return userHandler.GetUsersMembersFindAll(c)
	})
	apiGroup.Get("/:id", func(c *fiber.Ctx) error {
		return userHandler.GetUserFindById(c)
	})
	apiGroup.Post("/", func(c *fiber.Ctx) error {
		return userHandler.CreateUser(c)
	})
	apiGroup.Put("/:id", func(c *fiber.Ctx) error {
		return userHandler.UpdateUser(c)
	})
	apiGroup.Delete("/:id", func(c *fiber.Ctx) error {
		return userHandler.DeleteUser(c)
	})
}
