package routes

import (
	"go-notes-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	mainRoute := router.Group("/api")
	{
		mainRoute.GET("/", controllers.HealthCheck)

		notesRoute := mainRoute.Group("/notes")
		{
			notesRoute.GET("/", controllers.GetNotes)
			notesRoute.GET("/:id", controllers.GetNote)
			notesRoute.POST("/", controllers.CreateNote)
			notesRoute.PATCH("/:id", controllers.UpdateNote)
			notesRoute.DELETE("/:id", controllers.DeleteNote)
		}
	}
}