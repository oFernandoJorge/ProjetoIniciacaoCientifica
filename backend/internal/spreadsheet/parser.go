package spreadsheet

// ParseRows converte linhas em structs
func ParseRows(
	rows [][]string,
) []SpreadsheetRow {

	var submissions []SpreadsheetRow

	for index, row := range rows {

		// Ignora cabeçalho
		if index == 0 {
			continue
		}

		// Evita linhas incompletas (necessário até o índice 24)
		if len(row) < 25 {
			continue
		}

		submission := SpreadsheetRow{
			Title:            row[1],
			PresenterName:    row[5],
			Course:           row[23],
			KnowledgeArea:    row[22],
			Modality:         row[10],
			Campus:           row[24],
			AdvisorName:      row[2],
			PresentationType: row[10],
		}

		submissions = append(
			submissions,
			submission,
		)
	}

	return submissions
}