package scheduling

import (
	"testing"

	"ProjetoIniciacaoCientifica/internal/evaluator"
)

func TestAllocateEvaluators(t *testing.T) {

	roomAllocations := []RoomAllocation{
		{
			KnowledgeArea: "Tecnologia",

			PresentationType: "ORAL",
		},
	}

	evaluators := []evaluator.Evaluator{
		{
			KnowledgeArea: "Tecnologia",

			AcceptedPresentationType: "ORAL",

			MaxPresentations: 10,
		},
		{
			KnowledgeArea: "Tecnologia",

			AcceptedPresentationType: "BOTH",

			MaxPresentations: 10,
		},
	}

	allocations := AllocateEvaluators(
		roomAllocations,
		evaluators,
	)

	if len(allocations[0].Evaluators) != 2 {

		t.Errorf(
			"expected 2 evaluators, got %d",
			len(allocations[0].Evaluators),
		)
	}
}