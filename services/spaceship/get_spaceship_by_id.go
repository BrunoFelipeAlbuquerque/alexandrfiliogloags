package spaceshipservice

import (
	"imperial-fleet/config"
	"imperial-fleet/models"
	spaceshiprepositories "imperial-fleet/repositories/spaceship"
)

func GetSpaceshipByID(id uint) (*models.Spaceship, error) {
	return spaceshiprepositories.GetSpaceshipByID(config.DB, id)
}
