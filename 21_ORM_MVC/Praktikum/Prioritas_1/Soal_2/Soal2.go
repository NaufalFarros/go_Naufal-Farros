package main

import (
	"21_ORM_MVC/Praktikum/Prioritas_1/Soal_2/Database"
	"21_ORM_MVC/Praktikum/Prioritas_1/Soal_2/Routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	Database.InitDB()
	Database.InitialMigration()
}

func main() {
	Routes.UsersRoute()
}
