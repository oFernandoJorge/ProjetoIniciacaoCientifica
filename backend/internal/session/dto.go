package session

import "time"

// CreateSessionDTO representa payload
type CreateSessionDTO struct {
	RoomID uint `json:"room_id"`

	KnowledgeArea string `json:"knowledge_area"`

	PresentationType string `json:"presentation_type"`

	StartTime time.Time `json:"start_time"`

	EndTime time.Time `json:"end_time"`

	MaxCapacity int `json:"max_capacity"`
}