package http

import (
	"MrHour/entity"
	repo "MrHour/repo"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHours to list the hours
func GetHours(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var h []entity.Hour
	h, err := repo.FindAll()
	if err != nil {
		fmt.Println("Error while finding all hours")
	}
	c.JSON(http.StatusOK, gin.H{"data": h})
}

// SaveHour to save api
func SaveHour(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var h entity.Hour
	c.BindJSON(&h)
	// hour := entity.Hour{ID: 5, Note: "Todays work registration page", UpdatedAt: time.Now(), CreatedAt: time.Now(), Duration: "1 hour 30 minutes", Date: time.Now()}
	// repo.Save(hour)
	repo.Save(h)
}
