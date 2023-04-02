package Routes

import (
	"22_Middleware/Praktikum/Controllers"
	"22_Middleware/Praktikum/Middleware"

	"github.com/labstack/echo"
)

func UsersRoute(e *echo.Echo) {
	// create a new echo instance

	// JWT middleware
	jwtMiddleware := Middleware.JWTMiddleware{
		SecretKey: "NaufalFarros",
	}
	g := e.Group("/Auth")
	g.Use(jwtMiddleware.JWTAuth)
	// Route / to handler function
	g.GET("/users", Controllers.GetUsersController)
	g.GET("/users/:id", Controllers.GetUserController)
	e.POST("/users", Controllers.CreateUserController)
	g.DELETE("/users/:id", Controllers.DeleteUserController)
	g.PUT("/users/:id", Controllers.UpdateUserController)
	e.POST("/users/login", Controllers.LoginController)

	// start the server, and log if it fails
	// e.Logger.Fatal(e.Start(":8000"))
}
