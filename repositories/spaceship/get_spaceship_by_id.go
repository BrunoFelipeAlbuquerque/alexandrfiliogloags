package spaceshiprepositories

import (
	"imperial-fleet/models"

	"gorm.io/gorm"
)

func GetSpaceshipByID(db *gorm.DB, id uint) (*models.Spaceship, error) {
	var spaceship models.Spaceship
	if err := db.Preload("Armament").First(&spaceship, id).Error; err != nil {
		return nil, err
	}
	return &spaceship, nil
}
