package models

type Armament struct {
	ID          uint   `json:"-" gorm:"primaryKey"`
	SpaceshipID uint   `json:"-" gorm:"index"`
	Title       string `json:"title" gorm:"not null"`
	Qty         int    `json:"qty" gorm:"not null;check:qty > 0"`
}
