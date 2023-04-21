package main

import (
	"log"

	"github.com/NaufalFarros/golang-fiber/database"
	"github.com/NaufalFarros/golang-fiber/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
