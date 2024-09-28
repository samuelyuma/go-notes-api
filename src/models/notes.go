package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Note struct {
	Title       string         `json:"title" gorm:"type:text"`
	Tags        pq.StringArray `json:"tags" gorm:"type:text[]"`
	Description string         `json:"description" gorm:"type:text"`
}

type NoteColumn struct {
	gorm.Model
	Title       string         `json:"title" gorm:"type:text"`
	Tags        pq.StringArray `json:"tags" gorm:"type:text[]"`
	Description string         `json:"description" gorm:"type:text"`
}

func (NoteColumn) TableName() string {
    return "notes"
}