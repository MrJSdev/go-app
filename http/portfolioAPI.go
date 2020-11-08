package http

import (
	"MrHour/entity"
	repo "MrHour/repo"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPortfolios to receive
func GetPortfolios(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var messages []entity.Portfolio
	messages, err := repo.FindPortfolios()
	if err != nil {
		fmt.Println("Error while finding all messages")
	}
	c.JSON(http.StatusOK, gin.H{"data": messages})
}

// SavePortfolio to save api
func SavePortfolio(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var data entity.Portfolio

	c.BindJSON(&data)
	fmt.Println(data)
	repo.SavePortfolio(data)
	c.JSON(http.StatusOK, gin.H{"success": true})
}
