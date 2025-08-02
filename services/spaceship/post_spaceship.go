package spaceshipservice

import (
	"errors"

	"imperial-fleet/config"
	"imperial-fleet/dto"
	"imperial-fleet/models"
	spaceshiprepositories "imperial-fleet/repositories/spaceship"
)

func PostSpaceship(input dto.PostSpaceshipDTO) error {
	if !models.IsValidStatus(input.Status) {
		return errors.New("invalid status value")
	}

	spaceship := models.Spaceship{
		Name:   input.Name,
		Class:  input.Class,
		Crew:   input.Crew,
		Image:  input.Image,
		Value:  input.Value,
		Status: models.SpaceshipStatus(input.Status),
	}

	for _, arm := range input.Armament {
		spaceship.Armament = append(spaceship.Armament, models.Armament{
			Title: arm.Title,
			Qty:   arm.Qty,
		})
	}

	return spaceshiprepositories.PostSpaceship(config.DB, &spaceship)
}
