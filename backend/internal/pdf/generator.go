package pdf

import (
	"fmt"

	"ProjetoIniciacaoCientifica/internal/scheduling"

	"github.com/jung-kurt/gofpdf"
)

// GenerateSchedulePDF gera PDF final
func GenerateSchedulePDF(
	sessions []scheduling.ScheduledSession,
	outputPath string,
) error {

	pdf := gofpdf.New(
		"P",
		"mm",
		"A4",
		"",
	)

	pdf.SetFont(
		"Arial",
		"",
		12,
	)

	for _, session := range sessions {

		pdf.AddPage()

		// =========================
		// TÍTULO
		// =========================

		pdf.SetFont(
			"Arial",
			"B",
			16,
		)

		title := fmt.Sprintf(
			"SALA %s",
			session.RoomAllocation.RoomName,
		)

		pdf.Cell(40, 10, title)

		pdf.Ln(15)

		// =========================
		// HORÁRIO
		// =========================

		pdf.SetFont(
			"Arial",
			"",
			12,
		)

		timeText := fmt.Sprintf(
			"%s às %s",
			session.StartTime.Format("15:04"),
			session.EndTime.Format("15:04"),
		)

		pdf.Cell(40, 10, timeText)

		pdf.Ln(10)

		// =========================
		// ÁREA
		// =========================

		areaText := fmt.Sprintf(
			"Área: %s",
			session.RoomAllocation.KnowledgeArea,
		)

		pdf.Cell(40, 10, areaText)

		pdf.Ln(10)

		// =========================
		// AVALIADORES
		// =========================

		pdf.SetFont(
			"Arial",
			"B",
			12,
		)

		pdf.Cell(40, 10, "AVALIADORES")

		pdf.Ln(10)

		pdf.SetFont(
			"Arial",
			"",
			12,
		)

		for _, evaluator := range session.Evaluators {

			pdf.Cell(
				40,
				10,
				"- "+evaluator.Name,
			)

			pdf.Ln(8)
		}

		pdf.Ln(5)

		// =========================
		// TRABALHOS
		// =========================

		pdf.SetFont(
			"Arial",
			"B",
			12,
		)

		pdf.Cell(40, 10, "TRABALHOS")

		pdf.Ln(10)

		pdf.SetFont(
			"Arial",
			"",
			12,
		)

		for index, sub := range session.RoomAllocation.Submissions {

			line := fmt.Sprintf(
				"%d. %s",
				index+1,
				sub.Title,
			)

			pdf.MultiCell(
				0,
				8,
				line,
				"",
				"L",
				false,
			)
		}
	}

	return pdf.OutputFileAndClose(
		outputPath,
	)
}