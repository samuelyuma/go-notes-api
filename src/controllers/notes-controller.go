package controllers

import (
	"net/http"
	"time"

	"go-notes-api/src/config"
	"go-notes-api/src/helpers"
	"go-notes-api/src/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = config.ConnectDB()
}

func HealthCheck(c *gin.Context) {
	date := time.Now().Format("Monday, 02 January 2006 - 15:04 MST")
	helpers.SendResponse(c, http.StatusOK, "success", "", "Hello World!", date)
}

func CreateNote(c *gin.Context) {
    var note models.Note

	if err := c.ShouldBindJSON(&note); err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "failed", "Invalid input", "", nil)
		return
	}

	if note.Title == "" {
		helpers.SendResponse(c, http.StatusBadRequest, "failed", "Title is required", "", nil)
		return
	}

	if err := db.Create(&note).Error; err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "failed", "Failed to create note", "", nil)
		return
	}

	helpers.SendResponse(c, http.StatusCreated, "success", "", "Note created successfully", note)
}

func GetNotes (c *gin.Context) {
	var notes []models.Note

	if err := db.Find(&notes).Error; err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "failed", "Failed to get all notes", "", nil)
		return
	}

	helpers.SendResponse(c, http.StatusOK, "success", "", "Successfully get all notes", notes)
}

func GetNote (c *gin.Context) {
	id := c.Param("id")

	var note models.Note

	if err := db.First(&note, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.SendResponse(c, http.StatusNotFound, "failed", "Note not found", "", nil)
		} else {
			helpers.SendResponse(c, http.StatusInternalServerError, "failed", "Failed to get note", "", nil)
		}
		return
	}

	helpers.SendResponse(c, http.StatusOK, "success", "", "Successfully get note", note)
}

func UpdateNote (c *gin.Context) {
	id := c.Param("id")
	var data models.Note

	if err := c.ShouldBindJSON(&data); err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "failed", "Invalid input", "", nil)
		return
	}

	updates := map[string]interface{}{}
	if data.Title != "" {
		updates["title"] = data.Title
	}
	if data.Description != "" {
		updates["description"] = data.Description
	}
	if len(data.Tags) > 0 {
		updates["tags"] = data.Tags
	}

	if len(updates) == 0 {
		helpers.SendResponse(c, http.StatusBadRequest, "failed", "No fields to update", "", nil)
		return
	}

	result := db.Model(&models.Note{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "failed", "Failed to update note", "", nil)
		return
	}
	if result.RowsAffected == 0 {
		helpers.SendResponse(c, http.StatusNotFound, "failed", "Note not found", "", nil)
		return
	}

	helpers.SendResponse(c, http.StatusOK, "success", "", "Successfully update note", updates)
}

func DeleteNote (c *gin.Context) {
	id := c.Param("id")

	result := db.Delete(&models.Note{}, "id = ?", id)

    if result.Error != nil {
        helpers.SendResponse(c, http.StatusInternalServerError, "failed", result.Error.Error(), "", nil)
        return
    }

    if result.RowsAffected == 0 {
        helpers.SendResponse(c, http.StatusNotFound, "failed", "Note not found", "", nil)
        return
    }

	helpers.SendResponse(c, http.StatusOK, "success", "", "Successfully delete note", nil)
}