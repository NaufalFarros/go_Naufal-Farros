package main

import (
	"24_Clean_Hexagonal_Architecture/config"
	"24_Clean_Hexagonal_Architecture/routes"

	"github.com/labstack/echo"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	server := echo.New()
	routes.UserRoutes(server, db)
	server.Logger.Fatal(server.Start(":8000"))
}
