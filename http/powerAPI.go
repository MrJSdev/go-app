package http

import (
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetPower to receive
func GetPower(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var data map[string]float64
	c.BindJSON(&data)
	power := math.Pow(data["x"], data["power"])
	fmt.Println(data["x"], "this is data received", data["power"])
	const layout = "Mon, Jan 2, 2006"
	day := time.Now()
	tomorrow := day.Add(24 * time.Hour)
	fmt.Println(day.Format(layout))
	fmt.Println(tomorrow.Format(layout))
	c.JSON(http.StatusOK, gin.H{"status": "success", "success": true, "result": power, "layout": layout, "power": power})
}
