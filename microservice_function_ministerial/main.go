package main

import (
	"msvc_function_ministerial/infrastructure/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// Custom CORS configuration
	routers.NewRouter(app)
	app.Listen(":3005")
}
