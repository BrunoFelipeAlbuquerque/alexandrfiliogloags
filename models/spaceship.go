package models

import "gorm.io/gorm"

type SpaceshipStatus string

const (
	StatusOperational SpaceshipStatus = "Operational"
	StatusDamaged     SpaceshipStatus = "Damaged"
	StatusRetired     SpaceshipStatus = "Retired"
)

type Spaceship struct {
	gorm.Model
	ID       uint            `json:"id" gorm:"primaryKey"`
	Name     string          `json:"name" gorm:"not null"`
	Class    string          `json:"class" gorm:"not null"`
	Crew     int             `json:"crew" gorm:"not null;check:crew >= 0"`
	Image    string          `json:"image"`
	Value    float64         `json:"value" gorm:"not null;check:value >= 0"`
	Status   SpaceshipStatus `json:"status" gorm:"type:spaceship_status;not null"`
	Armament []Armament      `json:"armament" gorm:"foreignKey:SpaceshipID"`
}

func IsValidStatus(s string) bool {
	switch SpaceshipStatus(s) {
	case StatusOperational, StatusDamaged, StatusRetired:
		return true
	default:
		return false
	}
}
