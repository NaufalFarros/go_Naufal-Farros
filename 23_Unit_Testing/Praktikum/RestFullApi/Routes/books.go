package Routes

import (
	"23_Unit_Testing/Praktikum/RestFullApi/Controllers"
	"23_Unit_Testing/Praktikum/RestFullApi/Middleware"

	"github.com/labstack/echo"
)

func BooksRoute(e *echo.Echo) {
	// create a new echo instance
	// JWT middleware
	jwtMiddleware := Middleware.JWTMiddleware{
		SecretKey: "NaufalFarros",
	}
	g := e.Group("/Auth")
	g.Use(jwtMiddleware.JWTAuth)

	// Route / to handler function
	g.GET("/books", Controllers.GetBookController)
	g.GET("/books/:id", Controllers.GetBooksController)
	g.POST("/books", Controllers.CreateBooksController)
	g.DELETE("/books/:id", Controllers.DeleteBooksController)
	g.PUT("/books/:id", Controllers.UpdateBooksController)

	// start the server, and log if it fails
	// e.Logger.Fatal(e.Start(":8000"))
}
