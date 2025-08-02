package spaceshiprepositories

import (
	"imperial-fleet/models"

	"gorm.io/gorm"
)

func GetByID(db *gorm.DB, id uint) (*models.Spaceship, error) {
	var spaceship models.Spaceship
	if err := db.Preload("Armament").First(&spaceship, id).Error; err != nil {
		return nil, err
	}
	return &spaceship, nil
}

func PutSpaceship(db *gorm.DB, spaceship *models.Spaceship) error {
	if err := db.Where("spaceship_id = ?", spaceship.ID).Delete(&models.Armament{}).Error; err != nil {
		return err
	}
	return db.Save(spaceship).Error
}
