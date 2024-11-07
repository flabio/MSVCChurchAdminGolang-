package main

import (
	"github.com/gofiber/fiber/v2"
	"microservice_user.com/infrastructure/routers"
)

func main() {
	app := fiber.New()
	// Custom CORS configuration

	routers.NewRouter(app)
	app.Listen(":3006") // Start server on port 3000
}
