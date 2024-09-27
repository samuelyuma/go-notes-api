package main

import (
	"go-notes-api/src/config"
	"go-notes-api/src/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)

	router := gin.Default()

	routes.Routes(router)

	router.Run(":8080")
}