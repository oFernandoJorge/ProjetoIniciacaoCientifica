package spreadsheet

import "testing"

func TestParseRows(t *testing.T) {

	rows := [][]string{
		{
			"Title",
			"PresenterName",
			"Course",
			"KnowledgeArea",
			"Modality",
			"Campus",
			"AdvisorName",
			"PresentationType",
		},
		{
			"Sistema de Gestão",
			"Fernando",
			"ADS",
			"Tecnologia",
			"Pesquisa",
			"Fortaleza",
			"Professor João",
			"ORAL",
		},
	}

	submissions := ParseRows(rows)

	if len(submissions) != 1 {

		t.Errorf(
			"expected 1 submission, got %d",
			len(submissions),
		)
	}

	if submissions[0].Title != "Sistema de Gestão" {

		t.Errorf(
			"unexpected title",
		)
	}
}