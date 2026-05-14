package submission

// Repository define comportamentos do módulo submission
type Repository interface {
	Create(submission *Submission) error
	FindAll() ([]Submission, error)
}