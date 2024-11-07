package main

import (
	"github.com/gofiber/fiber/v2"
	"microservice_team_pesca.com/infrastructure/routers"
)

func main() {
	app := fiber.New()

	routers.NewRouter(app)
	app.Listen(":3003") // Start server on port 3000
}
