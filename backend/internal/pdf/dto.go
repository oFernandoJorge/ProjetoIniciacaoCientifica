package pdf

// PresentationPDF representa estrutura do PDF
type PresentationPDF struct {
	Room      string
	Date      string
	Turn      string
	Course    string
	Area      string
	Schedules []string
}