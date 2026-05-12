package scheduling

import "ProjetoIniciacaoCientifica/internal/submission"
import "ProjetoIniciacaoCientifica/internal/evaluator"
import "time"

// SubmissionGroup representa agrupamento
type SubmissionGroup struct {
	KnowledgeArea string

	PresentationType string

	Courses []string

	Submissions []submission.Submission
}

type RoomAllocation struct {
	RoomID uint

	RoomName string

	KnowledgeArea string

	PresentationType string

	Submissions []submission.Submission
}

// EvaluatorAllocation representa distribuição de avaliadores
type EvaluatorAllocation struct {
	RoomAllocation RoomAllocation

	Evaluators []evaluator.Evaluator
}

// ScheduledSession representa sessão agendada
type ScheduledSession struct {
	EvaluatorAllocation

	StartTime time.Time

	EndTime time.Time

	Shift string
}