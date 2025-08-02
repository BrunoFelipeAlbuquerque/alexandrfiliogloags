package spaceshipservice

import (
	"errors"

	"imperial-fleet/config"
	"imperial-fleet/dto"
	"imperial-fleet/models"
	spaceshiprepositories "imperial-fleet/repositories/spaceship"
)

func PutSpaceship(id uint, input dto.PutSpaceshipDTO) error {
	if !models.IsValidStatus(input.Status) {
		return errors.New("invalid status value")
	}

	spaceship, err := spaceshiprepositories.GetByID(config.DB, id)
	if err != nil {
		return err
	}

	spaceship.Name = input.Name
	spaceship.Class = input.Class
	spaceship.Crew = input.Crew
	spaceship.Image = input.Image
	spaceship.Value = input.Value
	spaceship.Status = models.SpaceshipStatus(input.Status)
	spaceship.Armament = nil

	for _, arm := range input.Armament {
		spaceship.Armament = append(spaceship.Armament, models.Armament{
			Title: arm.Title,
			Qty:   arm.Qty,
		})
	}

	return spaceshiprepositories.PutSpaceship(config.DB, spaceship)
}
