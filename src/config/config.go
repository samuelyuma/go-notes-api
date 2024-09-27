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
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	database := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUser , dbPass, dbHost , dbName)
	db, errorDB := gorm.Open(postgres.Open(database), &gorm.Config{})

	db.AutoMigrate(&models.NoteColumn{})

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