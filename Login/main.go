package main

import (
	"fiber-jwt/database"
	"fiber-jwt/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"log"
)

func main() {
	app := fiber.New(
		fiber.Config{Views: html.New("./view", ".html")},
	)

	app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
