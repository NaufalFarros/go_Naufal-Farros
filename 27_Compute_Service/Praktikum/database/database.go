package database

import (
	"log"
	"os"

	"github.com/NaufalFarros/golang-fiber/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func Connect() {
	dsn := "host=database-2.cgl7qm7eiaf0.ap-southeast-1.rds.amazonaws.com user=postgres password=GodhandLND356 dbname=farosRDS port=5432 sslmode=require TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migration")
	// auto migrate
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.Role{})

	Database = DbInstance{Db: db}
}
