package main

import (
	"log"

	"github.com/safe/church"
	_ "github.com/safe/docs"
	"github.com/safe/ministerial"
	"github.com/safe/ministerialmember"
	"github.com/safe/rol"
	"github.com/safe/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

func main() {

	app := fiber.New()
	// Custom CORS configuration
	// Ruta para la documentación Swagger
	app.Get("/swagger/*", swagger.HandlerDefault) // swagger.HandlerDefault es el handler predeterminado
	// default
	//http://localhost:81/auth/login

	// Enable CORS with specific settings
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	//router
	rol.NewRolRouter(app)
	ministerial.NewMinisterialRouter(app)
	user.NewUserRouter(app)
	church.NewChurchRouter(app)
	ministerialmember.NewMinisterialMemberRouter(app)

	// Start server
	log.Println("Server listening on port 3081")
	log.Fatal(app.Listen(":3081"))
}
