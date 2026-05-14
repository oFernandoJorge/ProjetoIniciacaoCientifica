package scheduling

import (
	"ProjetoIniciacaoCientifica/internal/evaluator"
)

// AllocateEvaluators distribui avaliadores
func AllocateEvaluators(
	roomAllocations []RoomAllocation,
	evaluators []evaluator.Evaluator,
) []EvaluatorAllocation {

	var allocations []EvaluatorAllocation

	evaluatorUsage := make(map[uint]int)

	for _, roomAllocation := range roomAllocations {

		var selectedEvaluators []evaluator.Evaluator

		for _, eval := range evaluators {

			// Verifica área
			if eval.KnowledgeArea != roomAllocation.KnowledgeArea {
				continue
			}

			// Verifica modalidade
			if eval.AcceptedPresentationType != "BOTH" &&
				eval.AcceptedPresentationType != roomAllocation.PresentationType {

				continue
			}

			// Verifica limite
			if evaluatorUsage[eval.ID] >= eval.MaxPresentations {
				continue
			}

			selectedEvaluators = append(
				selectedEvaluators,
				eval,
			)

			evaluatorUsage[eval.ID]++

			// Máximo 2 avaliadores por sessão
			if len(selectedEvaluators) == 2 {
				break
			}
		}

		allocation := EvaluatorAllocation{
			RoomAllocation: roomAllocation,

			Evaluators: selectedEvaluators,
		}

		allocations = append(
			allocations,
			allocation,
		)
	}

	return allocations
}