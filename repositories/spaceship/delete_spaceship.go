package spaceshiprepositories

import (
	"imperial-fleet/models"

	"gorm.io/gorm"
)

func DeleteSpaceship(db *gorm.DB, id uint) error {
	var spaceship models.Spaceship
	if err := db.First(&spaceship, id).Error; err != nil {
		return err
	}
	return db.Delete(&spaceship).Error
}
