package main

import (
	"22_Middleware/Praktikum/Database"
	"22_Middleware/Praktikum/Models"
	"22_Middleware/Praktikum/Routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	Database.DB.AutoMigrate(&Models.User{}, &Models.Book{})

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339_nano}] ${status} ${method} ${uri} ${latency_human}\n",
	}))
	Routes.UsersRoute(e)
	Routes.BooksRoute(e)

	e.Logger.Fatal(e.Start(":8000"))
}
