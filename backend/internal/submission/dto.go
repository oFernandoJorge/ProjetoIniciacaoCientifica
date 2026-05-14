package submission

// CreateSubmissionDTO representa os dados
// recebidos na criação de uma submissão
type CreateSubmissionDTO struct {
	Title            string `json:"title"`
	PresenterName    string `json:"presenter_name"`
	Course           string `json:"course"`
	KnowledgeArea    string `json:"knowledge_area"`
	Modality         string `json:"modality"`
	Campus           string `json:"campus"`
	AdvisorName      string `json:"advisor_name"`
	PresentationType string `json:"presentation_type"`
}

// UpdateSubmissionDTO representa atualização
// parcial ou total de submissão
type UpdateSubmissionDTO struct {
	Title            string `json:"title"`
	PresenterName    string `json:"presenter_name"`
	Course           string `json:"course"`
	KnowledgeArea    string `json:"knowledge_area"`
	Modality         string `json:"modality"`
	Campus           string `json:"campus"`
	AdvisorName      string `json:"advisor_name"`
	PresentationType string `json:"presentation_type"`
}

// SubmissionResponse representa resposta enviada
// pela API ao frontend
type SubmissionResponse struct {
	ID               uint   `json:"id"`
	Title            string `json:"title"`
	PresenterName    string `json:"presenter_name"`
	Course           string `json:"course"`
	KnowledgeArea    string `json:"knowledge_area"`
	Modality         string `json:"modality"`
	Campus           string `json:"campus"`
	AdvisorName      string `json:"advisor_name"`
	PresentationType string `json:"presentation_type"`
}

// SpreadsheetSubmissionDTO representa linha
// processada da planilha Excel
type SpreadsheetSubmissionDTO struct {
	Title            string
	PresenterName    string
	Course           string
	KnowledgeArea    string
	Modality         string
	Campus           string
	AdvisorName      string
	PresentationType string
}

// ScheduleSubmissionDTO representa submissão
// já processada pelo algoritmo de ensalamento
type ScheduleSubmissionDTO struct {
	SubmissionID  uint
	Title         string
	PresenterName string
	Course        string
	KnowledgeArea string

	Room      string
	Date      string
	StartTime string
	EndTime   string
}

// PDFSubmissionDTO representa estrutura final
// usada na geração do PDF institucional
type PDFSubmissionDTO struct {
	Room             string
	Date             string
	Turn             string
	Course           string
	KnowledgeArea    string
	StartTime        string
	EndTime          string
	Title            string
	PresenterName    string
	AdvisorName      string
	PresentationType string
}