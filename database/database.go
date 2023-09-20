package database

import (
	"Sviluppo/go/go-fiber/config"
	"Sviluppo/go/go-fiber/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	// p := config.Config("POSTGRES_PORT")
	// port, err := strconv.ParseUint(p, 10, 32)
	// if err != nil {
	// 	fmt.Println("Error parsing str to int")
	// }

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Rome", config.Config("POSTGRES_HOST"), config.Config("POSTGRES_USER"), config.Config("POSTGRES_PASSWORD"), config.Config("POSTGRES_DATABASE"), port)
	dsn := config.Config("POSTGRES_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to the database \n", err)
		os.Exit(2)
	}
	log.Println("Connected")

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&models.User{}, &models.Product{})
	fmt.Println("Database Migrated")
	DB = Dbinstance{
		Db: db,
	}

}
