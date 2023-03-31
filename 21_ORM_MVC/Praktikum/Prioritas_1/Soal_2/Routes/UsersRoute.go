package routes

import (
	"21_ORM_MVC/Praktikum/Prioritas_1/Soal_2/Controllers"

	"github.com/labstack/echo"
)

func UsersRoute() {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/users", Controllers.GetUsersController)
	e.GET("/users/:id", Controllers.GetUserController)
	e.POST("/users", Controllers.CreateUserController)
	e.DELETE("/users/:id", Controllers.DeleteUserController)
	e.PUT("/users/:id", Controllers.UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
