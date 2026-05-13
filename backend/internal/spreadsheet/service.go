package spreadsheet

// ProcessSpreadsheet processa Excel
func ProcessSpreadsheet(
	filePath string,
) ([]SpreadsheetRow, error) {

	rows, err := ReadSpreadsheet(
		filePath,
	)

	if err != nil {
		return nil, err
	}

	submissions := ParseRows(rows)

	return submissions, nil
}