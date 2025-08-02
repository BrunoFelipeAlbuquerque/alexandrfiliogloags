package spaceshiphandler

import (
	"net/http"
	"strconv"

	"imperial-fleet/dto"
	spaceshipservice "imperial-fleet/services/spaceship"

	"github.com/gin-gonic/gin"
)

func PutSpaceship(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var input dto.PutSpaceshipDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := spaceshipservice.PutSpaceship(uint(id), input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
