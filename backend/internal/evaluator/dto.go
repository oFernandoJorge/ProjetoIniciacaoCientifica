package evaluator

// CreateEvaluatorDTO representa payload de criação
type CreateEvaluatorDTO struct {
	Name string `json:"name"`

	Email string `json:"email"`

	Course string `json:"course"`

	KnowledgeArea string `json:"knowledge_area"`

	AvailableMorning bool `json:"available_morning"`

	AvailableAfternoon bool `json:"available_afternoon"`

	AvailableNight bool `json:"available_night"`

	MaxPresentations int `json:"max_presentations"`

	AcceptedPresentationType string `json:"accepted_presentation_type"`
}