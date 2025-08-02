package spaceshipservice

import (
	"imperial-fleet/config"
	"imperial-fleet/dto"
	"imperial-fleet/models"
	spaceshiprepositories "imperial-fleet/repositories/spaceship"
)

func GetSpaceships(filter dto.GetSpaceshipFilterDTO) ([]models.Spaceship, int64, error) {
	return spaceshiprepositories.GetSpaceships(config.DB, filter)
}
