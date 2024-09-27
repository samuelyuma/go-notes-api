package routes

import (
	"go-notes-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", controllers.GetNotes)
	router.GET("/:id", controllers.GetNote)
	router.POST("/", controllers.CreateNote)
	router.PATCH("/:id", controllers.UpdateNote)
	router.DELETE("/:id", controllers.DeleteNote)
}