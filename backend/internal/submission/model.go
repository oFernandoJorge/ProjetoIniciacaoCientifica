package submission

import "gorm.io/gorm"

// Submission representa trabalho submetido
// ao evento acadêmico
type Submission struct {
	gorm.Model
	Title            string
	PresenterName    string
	Course           string
	KnowledgeArea    string
	Modality         string
	Campus           string
	AdvisorName      string
	PresentationType string
}
