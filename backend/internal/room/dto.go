package room

// RoomResponse representa retorno da API
type RoomResponse struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	Floor            int    `json:"floor"`
	PresentationType string `json:"presentation_type"`
	Capacity         int    `json:"capacity"`
	IsAvailable      bool   `json:"is_available"`
}