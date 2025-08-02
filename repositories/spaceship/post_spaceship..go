package spaceshiprepositories

import (
	"imperial-fleet/models"

	"gorm.io/gorm"
)

func PostSpaceship(db *gorm.DB, spaceship *models.Spaceship) error {
	return db.Create(spaceship).Error
}
