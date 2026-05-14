package session

import (
	"time"

	"gorm.io/gorm"
)

// Session representa sessão de apresentação
type Session struct {
	gorm.Model

	RoomID uint

	KnowledgeArea string

	PresentationType string

	StartTime time.Time

	EndTime time.Time

	MaxCapacity int
}