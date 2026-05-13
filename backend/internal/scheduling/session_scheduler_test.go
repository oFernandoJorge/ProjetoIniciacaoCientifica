package scheduling

import (
	"testing"
	"ProjetoIniciacaoCientifica/internal/submission"
)

func TestGenerateSchedules(t *testing.T) {

	allocations := []EvaluatorAllocation{
		{
			RoomAllocation: RoomAllocation{
				PresentationType: "ORAL",

				Submissions: []submission.Submission{
					{},
					{},
					{},
				},
			},
		},
	}

	sessions := GenerateSchedules(
		allocations,
	)

	if len(sessions) != 1 {

		t.Errorf(
			"expected 1 session, got %d",
			len(sessions),
		)
	}

	if sessions[0].Shift != "MANHÃ" {

		t.Errorf(
			"expected MANHÃ, got %s",
			sessions[0].Shift,
		)
	}
}