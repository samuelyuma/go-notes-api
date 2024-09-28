package config

import (
	"fmt"
	"os"

	"go-notes-api/src/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	errorENV := godotenv.Load()

	if errorENV != nil {
		panic("Failed to load env file!")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	database := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	db, errorDB := gorm.Open(postgres.Open(database), &gorm.Config{})

	err := db.AutoMigrate(&models.NoteColumn{})
	if err != nil {
		panic(fmt.Sprintf("Failed to migrate the database: %v", err))
	}

	if errorDB != nil {
		panic("Failed to connect postgres database!")
	} else {
		fmt.Println("Successfully connect to postgres database!")
	}

	return db
}

func DisconnectDB(db *gorm.DB) {
	dbPostgres, err := db.DB()

	if err != nil {
		panic("Failed to disconnect with postgres database")
	} else {
		fmt.Println("Successfully disconnect to postgres database!")
	}

	dbPostgres.Close()
}