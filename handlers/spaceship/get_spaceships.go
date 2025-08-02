package spaceshiphandler

import (
	"net/http"

	"imperial-fleet/dto"
	spaceshipservice "imperial-fleet/services/spaceship"

	"github.com/gin-gonic/gin"
)

func GetSpaceships(c *gin.Context) {
	var filter dto.GetSpaceshipFilterDTO
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query parameters"})
		return
	}

	spaceships, total, err := spaceshipservice.GetSpaceships(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var result []gin.H
	for _, s := range spaceships {
		result = append(result, gin.H{
			"id":     s.ID,
			"name":   s.Name,
			"status": s.Status,
		})
	}

	if filter.Limit <= 0 {
		filter.Limit = 10
	}

	totalPages := (total + int64(filter.Limit) - 1) / int64(filter.Limit)

	c.JSON(http.StatusOK, gin.H{
		"data":       result,
		"total":      total,
		"page":       filter.Page,
		"per_page":   filter.Limit,
		"total_page": totalPages,
	})
}
