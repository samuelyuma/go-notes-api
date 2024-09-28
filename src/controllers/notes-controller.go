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

var db *gorm.DB = config.ConnectDB()

func HealthCheck(c *gin.Context) {
	var date, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")

	var formattedDate = date.Format("Monday, 01 January 2024 - 15:04 MST")

	helpers.SendResponse(c, http.StatusOK, "success", "", "Hello World!", formattedDate)
}

func CreateNote(c *gin.Context) {
    var data models.Note

    if err := c.ShouldBindJSON(&data); err != nil {
        helpers.SendResponse(c, http.StatusBadRequest, "failed", err.Error(), "", nil)
        return
    }

    if data.Title == "" {
        helpers.SendResponse(c, http.StatusBadRequest, "failed", "Title is required", "", nil)
        return
    }

    note := models.Note{
        Title:       data.Title,
        Description: data.Description,
        Tags:        data.Tags,
    }

    if err := db.Create(&note).Error; err != nil {
        helpers.SendResponse(c, http.StatusBadRequest, "failed", err.Error(), "", nil)
        return
    }

    helpers.SendResponse(c, http.StatusCreated, "success", "", "Note created successfully", note)
}

func GetNotes (c *gin.Context) {
	var notes []models.Note

	if err := db.Find(&notes).Error; err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "failed", err.Error(), "", nil)
    	return
	}

	helpers.SendResponse(c, http.StatusOK, "success", "", "Successfully get all notes", notes)
}

func GetNote (c *gin.Context) {
	id := c.Param("id")

	var note models.Note

	if err := db.Where("id = ?", id).First(&note).Error; err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "failed", err.Error(), "", nil)
    	return
	}

	helpers.SendResponse(c, http.StatusOK, "success", "", "Successfully get note", note)
}

func UpdateNote (c *gin.Context) {
	id := c.Param("id")

	var data models.Note

	if err := c.ShouldBindJSON(&data); err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "failed", err.Error(), "", nil)
    	return
	}

	note := make(map[string]interface{})

	if data.Title != "" {
		note["title"] = data.Title
	}

	if data.Description != "" {
		note["description"] = data.Description
	}

	if len(data.Tags) > 0 {
		note["tags"] = data.Tags
	}

	if len(note) == 0 {
		helpers.SendResponse(c, http.StatusBadRequest, "failed", "No fields to update", "", nil)
		return
	}

	if err := db.Model(&models.Note{}).Where("id = ?", id).Updates(note).Error; err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "failed", err.Error(), "", nil)
    	return
	}

	helpers.SendResponse(c, http.StatusOK, "success", "", "Successfully update note", note)
}

func DeleteNote (c *gin.Context) {
	id := c.Param("id")

	var note models.Note

	if err := db.First(&note, "id = ?", id).Error; err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "failed", err.Error(), "", nil)
		return
	}

	if err := db.Where("id = ?", id).Delete(&note).Error; err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "failed", err.Error(), "", nil)
		return
	}

	helpers.SendResponse(c, http.StatusOK, "success", "", "Successfully delete note", nil)
}