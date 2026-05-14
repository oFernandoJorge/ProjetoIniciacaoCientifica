package pdf

import (
	"testing"
	"time"

	"ProjetoIniciacaoCientifica/internal/evaluator"
	"ProjetoIniciacaoCientifica/internal/scheduling"
	"ProjetoIniciacaoCientifica/internal/submission"
)

func TestGenerateSchedulePDF(t *testing.T) {

	sessions := []scheduling.ScheduledSession{
		{
			EvaluatorAllocation: scheduling.EvaluatorAllocation{
				RoomAllocation: scheduling.RoomAllocation{
					RoomName: "211",

					KnowledgeArea: "Tecnologia",

					Submissions: []submission.Submission{
						{
							Title: "Sistema de Gestão",
						},
					},
				},

				Evaluators: []evaluator.Evaluator{
					{
						Name: "Professor João",
					},
				},
			},

			StartTime: time.Now(),

			EndTime: time.Now().Add(
				2 * time.Hour,
			),
		},
	}

	err := GenerateSchedulePDF(
		sessions,
		"schedule_test.pdf",
	)

	if err != nil {

		t.Errorf(
			"expected nil, got %v",
			err,
		)
	}
}