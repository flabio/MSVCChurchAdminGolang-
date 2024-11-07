package main

import (
	"github.com/gofiber/fiber/v2"
	"microservice_ministerial.com/infrastructure/routers"
)

func main() {
	app := fiber.New()

	routers.NewMinisterialRouter(app)
	app.Listen(":3004") // Start server on port 3000
}
