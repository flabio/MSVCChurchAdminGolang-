package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"microservice_ministerial.com/infrastructure/routers"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080", // Specify allowed origins
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Content-Type, Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true, // Allow credentials
	}))

	routers.NewMinisterialRouter(app)
	app.Listen(":3004") // Start server on port 3000
}
