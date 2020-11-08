package http

import (
	"MrHour/entity"
	repo "MrHour/repo"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMessages to list the hours
func GetMessages(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var messages []entity.Message
	messages, err := repo.FindMessages()
	if err != nil {
		fmt.Println("Error while finding all messages")
	}
	c.JSON(http.StatusOK, gin.H{"data": messages})
}

// SaveMessage to save api
func SaveMessage(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var msg entity.Message
	c.BindJSON(&msg)
	repo.SendMessage(msg)
	c.JSON(http.StatusOK, gin.H{"success": true})
}
