package routes

import (
	controllers "github.com/NaufalFarros/golang-fiber/Controllers"
	middleware "github.com/NaufalFarros/golang-fiber/Middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/register", controllers.CreateUsers)
	app.Post("/login", controllers.LoginUsers)
	app.Post("/logout", controllers.LogoutUsers)

	api := app.Group("/api")
	api.Use(middleware.Authentication)
	api.Get("/usersDetail", controllers.GetUsers)
	api.Post("/update", controllers.UpdateUser)
}
