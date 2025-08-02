package spaceshiphandler

import (
	"net/http"
	"strconv"

	spaceshipservice "imperial-fleet/services/spaceship"

	"github.com/gin-gonic/gin"
)

func GetSpaceshipByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	spaceship, err := spaceshipservice.GetSpaceshipByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "spaceship not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       spaceship.ID,
		"name":     spaceship.Name,
		"class":    spaceship.Class,
		"crew":     spaceship.Crew,
		"image":    spaceship.Image,
		"value":    spaceship.Value,
		"status":   spaceship.Status,
		"armament": spaceship.Armament,
	})
}
