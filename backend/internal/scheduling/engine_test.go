package scheduling

import (
	"testing"

	"ProjetoIniciacaoCientifica/internal/submission"
)

func TestGroupSubmissionsByArea(t *testing.T) {

	submissions := []submission.Submission{
		{
			Title: "Projeto 1",

			Course: "ADS",

			KnowledgeArea: "Tecnologia",

			PresentationType: "ORAL",
		},
		{
			Title: "Projeto 2",

			Course: "CC",

			KnowledgeArea: "Tecnologia",

			PresentationType: "ORAL",
		},
		{
			Title: "Projeto 3",

			Course: "ES",

			KnowledgeArea: "Tecnologia",

			PresentationType: "ORAL",
		},
	}

	groups := GroupSubmissionsByArea(submissions)

	if len(groups) != 1 {
		t.Errorf("expected 1 group, got %d", len(groups))
	}

	if len(groups[0].Courses) != 3 {
		t.Errorf(
			"expected 3 courses, got %d",
			len(groups[0].Courses),
		)
	}
}