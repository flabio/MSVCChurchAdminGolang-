package main

import (
	"github.com/gofiber/fiber/v2"
	"microservice_church.com/infrastructure/routers"
)

func main() {
	app := fiber.New()
	// Custom CORS configuration

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "*", // Specify allowed origins
	// 	AllowMethods:     "GET,POST,PUT,DELETE",
	// 	AllowHeaders:     "Content-Type, Authorization",
	// 	ExposeHeaders:    "Content-Length",
	// 	AllowCredentials: true, // Allow credentials
	// }))
	routers.NewChurchRouter(app)
	app.Listen(":3002")
}
