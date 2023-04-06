package routes

import (
	"24_Clean_Hexagonal_Architecture/Middleware"
	"24_Clean_Hexagonal_Architecture/controller"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func UserRoutes(app *echo.Echo, db *gorm.DB) {

	jwtMiddleware := Middleware.JWTMiddleware{
		SecretKey: "NaufalFarros",
	}
	app.Use(jwtMiddleware.JWTAuth)

	app.GET("/users", controller.GetAllUsers(db))
	app.POST("/users", controller.CreateUser(db))

}
