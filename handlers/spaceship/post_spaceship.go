package spaceshiphandler

import (
	"net/http"

	"imperial-fleet/dto"
	spaceshipservice "imperial-fleet/services/spaceship"

	"github.com/gin-gonic/gin"
)

func PostSpaceship(c *gin.Context) {
	var input dto.PostSpaceshipDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := spaceshipservice.PostSpaceship(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
