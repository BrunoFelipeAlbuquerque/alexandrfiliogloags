package spaceshipservice

import (
	"imperial-fleet/config"
	spaceshiprepositories "imperial-fleet/repositories/spaceship"
)

func DeleteSpaceship(id uint) error {
	return spaceshiprepositories.DeleteSpaceship(config.DB, id)
}
