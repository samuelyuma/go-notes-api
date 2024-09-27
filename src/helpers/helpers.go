package helpers

import (
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status	string      `json:"status"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponse(c *gin.Context, httpStatus int, status string, errorMsg string, message string, data interface{}) {
	response := ResponseData{
		Status:  status,
		Error:   errorMsg,
		Message: message,
		Data:    data,
	}

	c.JSON(httpStatus, response)
}