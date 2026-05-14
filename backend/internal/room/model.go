package room

import "gorm.io/gorm"

// Room representa sala de apresentação
type Room struct {
	gorm.Model

	Name             string
	Floor            int
	PresentationType string
	Capacity         int
	IsAvailable      bool
}