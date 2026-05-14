package spreadsheet

import (
	"github.com/xuri/excelize/v2"
)

// ReadSpreadsheet lê arquivo Excel
func ReadSpreadsheet(
	filePath string,
) ([][]string, error) {

	file, err := excelize.OpenFile(
		filePath,
	)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	sheets := file.GetSheetList()

	rows, err := file.GetRows(
		sheets[0],
	)

	if err != nil {
		return nil, err
	}

	return rows, nil
}