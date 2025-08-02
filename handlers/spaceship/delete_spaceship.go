package spaceshiphandler

import (
	"net/http"
	"strconv"

	spaceshipservice "imperial-fleet/services/spaceship"

	"github.com/gin-gonic/gin"
)

func DeleteSpaceship(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := spaceshipservice.DeleteSpaceship(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "spaceship not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
