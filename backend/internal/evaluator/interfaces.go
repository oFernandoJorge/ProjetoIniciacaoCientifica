package evaluator

// Repository define comportamentos do evaluator
type Repository interface {
	Create(evaluator *Evaluator) error
	FindAll() ([]Evaluator, error)
}