package evaluator

import "gorm.io/gorm"

// Evaluator representa avaliador do evento
type Evaluator struct {
	gorm.Model

	Name string `gorm:"not null"`

	Email string `gorm:"unique;not null"`

	Course string `gorm:"not null"`

	KnowledgeArea string `gorm:"not null"`

	AvailableMorning bool

	AvailableAfternoon bool

	AvailableNight bool

	MaxPresentations int

	AcceptedPresentationType string
}