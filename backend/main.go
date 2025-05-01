package main

import (
	"github.com/aungmyatmoe11/Go-JWT-Auth/database"
	"github.com/aungmyatmoe11/Go-JWT-Auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// ! connect database
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // Specify your allowed origins
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
