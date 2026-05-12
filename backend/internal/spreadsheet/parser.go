package spreadsheet

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

// ParseSpreadsheet lê planilha Excel
func ParseSpreadsheet(path string) error {
	file, err := excelize.OpenFile(path)
	if err != nil {
		return err
	}
	sheets := file.GetSheetList()
	for _, sheet := range sheets {
		rows, err := file.GetRows(sheet)
		if err != nil {
			return err
		}
		fmt.Println("Sheet:", sheet)
		for _, row := range rows {
			fmt.Println(row)
		}
	}
	return nil
}
