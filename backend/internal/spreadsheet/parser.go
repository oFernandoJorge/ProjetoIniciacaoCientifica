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

		// Evita linhas incompletas
		if len(row) < 8 {
			continue
		}

		submission := SpreadsheetRow{
			Title: row[0],

			PresenterName: row[1],

			Course: row[2],

			KnowledgeArea: row[3],

			Modality: row[4],

			Campus: row[5],

			AdvisorName: row[6],

			PresentationType: row[7],
		}

		submissions = append(
			submissions,
			submission,
		)
	}

	return submissions
}