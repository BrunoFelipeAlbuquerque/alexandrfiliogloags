package handlers

import (
	spaceshiphandler "imperial-fleet/handlers/spaceship"

	"github.com/gin-gonic/gin"
)

func RegisterSpaceshipRoutes(r *gin.Engine) {
	r.GET("/spaceships", spaceshiphandler.GetSpaceships)
	r.GET("/spaceships/:id", spaceshiphandler.GetSpaceshipByID)
	r.POST("/spaceships", spaceshiphandler.PostSpaceship)
	r.PUT("/spaceships/:id", spaceshiphandler.PutSpaceship)
	r.DELETE("/spaceships/:id", spaceshiphandler.DeleteSpaceship)
}
