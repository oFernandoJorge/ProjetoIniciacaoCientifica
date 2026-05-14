package pdf

import (
	"fmt"
	"io"
	"strings"

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

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func GeneratePresentationSchedulePDF(request GeneratePresentationPdfRequest, output io.Writer) error {
	pdf := gofpdf.New("L", "mm", "A4", "")
	tr := pdf.UnicodeTranslatorFromDescriptor("")

	pdf.SetHeaderFunc(func() {
		// 1. Faixa superior (Verde Escuro - #345a4d)
		pdf.SetFillColor(52, 90, 77)
		pdf.Rect(0, 0, 297, 35, "F")

		// Logos alinhadas horizontalmente
		pdf.SetTextColor(255, 255, 255)

		// Lado Esquerdo: CNX CONEXÃO
		pdf.SetXY(15, 8)
		pdf.SetFont("Times", "B", 40)
		pdf.Cell(0, 20, "CNX")

		pdf.SetXY(55, 12)
		pdf.SetFont("Times", "B", 18)
		pdf.Cell(0, 10, tr("CONEXÃO"))
		pdf.SetXY(55, 19)
		pdf.SetFont("Times", "", 10)
		pdf.Cell(0, 10, "UNIFAMETRO")

		// Lado Direito: Unifametro
		pdf.SetXY(210, 10)
		pdf.SetFont("Times", "B", 24)
		pdf.Cell(0, 12, "Unifametro")
		pdf.SetXY(210, 20)
		pdf.SetFont("Times", "I", 9)
		pdf.Cell(0, 10, tr("Formar para Transformar"))

		pdf.SetTextColor(0, 0, 0)
	})

	pdf.SetFooterFunc(func() {
		pdf.SetY(-22)
		pdf.SetFont("Times", "I", 8)
		pdf.SetFillColor(255, 255, 0) // Amarelo

		text1 := tr("*Caso o apresentador não possa estar presente no horário, outro autor do mesmo trabalho poderá apresentá-lo.")
		text2 := tr("(é obrigatório o apresentador estar inscrito e estar usando o crachá do evento).")

		w1 := pdf.GetStringWidth(text1) + 4
		pdf.SetX((297 - w1) / 2)
		pdf.CellFormat(w1, 5, text1, "", 1, "C", true, 0, "")

		w2 := pdf.GetStringWidth(text2) + 4
		pdf.SetX((297 - w2) / 2)
		pdf.CellFormat(w2, 5, text2, "", 1, "C", true, 0, "")
	})

	pdf.SetTitle("Organizador de Salas", false)
	pdf.SetMargins(10, 10, 10) 
	pdf.SetAutoPageBreak(true, 30)

	for i, item := range request.Items {
		pdf.AddPage()

		// Posição inicial do conteúdo (Abaixo da faixa verde de 35mm)
		pdf.SetY(40) 

		// 2. Título da Sessão
		pdf.SetFont("Times", "B", 11)
		sessionTitle := fmt.Sprintf("SESSÃO %02d – (%s)", i+1, item.PresentationType)
		pdf.CellFormat(277, 7, tr(sessionTitle), "", 1, "C", false, 0, "")

		// 3. Metadados (DATA, TURNO, SALA, ÁREA)
		pdf.SetFont("Times", "B", 9)
		metadata := fmt.Sprintf("DATA: %s TURNO: %s SALA: %s / ÁREA DE CONHECIMENTO: %s",
			request.Date, request.Turn, item.RoomName, item.KnowledgeArea)
		// MultiCell para evitar corte se a Área for muito grande
		pdf.MultiCell(277, 5, tr(metadata), "", "C", false)

		// 4. Curso
		if len(item.Courses) > 0 {
			pdf.SetFont("Times", "B", 9)
			coursesText := fmt.Sprintf("Curso(s): %s", strings.Join(item.Courses, " / "))
			// MultiCell é essencial aqui pois a lista de cursos pode ser enorme
			pdf.MultiCell(277, 5, tr(coursesText), "", "C", false)
		}

		pdf.Ln(4)

		// 5. Cabeçalho da Tabela
		pdf.SetFillColor(244, 164, 96)
		pdf.SetFont("Times", "B", 9)
		pdf.CellFormat(30, 8, tr("HORÁRIO"), "1", 0, "C", true, 0, "")
		pdf.CellFormat(170, 8, tr("Título"), "1", 0, "C", true, 0, "")
		pdf.CellFormat(77, 8, tr("Nome do apresentador"), "1", 1, "C", true, 0, "")

		// 6. Linhas da Tabela
		pdf.SetFont("Times", "B", 9)
		for _, row := range item.Submissions {
			// Verifica se a próxima linha cabe na página (Margem de segurança)
			if pdf.GetY() > 175 { 
				pdf.AddPage()
				pdf.SetY(40) // Reinicia posição após o header automático
				
				// Opcional: Repetir cabeçalho da tabela na nova página
				pdf.SetFillColor(244, 164, 96)
				pdf.CellFormat(30, 8, tr("HORÁRIO"), "1", 0, "C", true, 0, "")
				pdf.CellFormat(170, 8, tr("Título"), "1", 0, "C", true, 0, "")
				pdf.CellFormat(77, 8, tr("Nome do apresentador"), "1", 1, "C", true, 0, "")
			}
			writeStyledRow(pdf, row.Time, row.Title, row.PresenterName, tr)
		}
	}

	return pdf.Output(output)
}

func writeStyledRow(pdf *gofpdf.Fpdf, timeValue, title, presenter string, tr func(string) string) {
	timeValue = tr(timeValue)
	title = tr(title)
	presenter = tr(presenter)

	rowHeight := 6.0
	x, y := pdf.GetX(), pdf.GetY()

	// Calcula a altura necessária
	titleLines := pdf.SplitLines([]byte(title), 170)
	presenterLines := pdf.SplitLines([]byte(presenter), 77)

	maxLines := len(titleLines)
	if len(presenterLines) > maxLines {
		maxLines = len(presenterLines)
	}
	if maxLines < 1 { maxLines = 1 }

	cellHeight := float64(maxLines) * rowHeight
	if cellHeight < 10 { cellHeight = 10 }

	// Verifica se cabe na página antes de desenhar
	if y + cellHeight > 210 - 30 {
		pdf.AddPage()
		y = pdf.GetY()
		x = pdf.GetX()
	}

	// Horário (Borda e Texto)
	pdf.CellFormat(30, cellHeight, timeValue, "1", 0, "C", false, 0, "")

	// Título (Borda manual e Texto centralizado verticalmente)
	pdf.Rect(x+30, y, 170, cellHeight, "D")
	titleY := y + (cellHeight - float64(len(titleLines))*rowHeight)/2
	pdf.SetXY(x+30, titleY)
	pdf.MultiCell(170, rowHeight, title, "", "C", false)

	// Apresentador (Borda manual e Texto centralizado verticalmente)
	pdf.Rect(x+30+170, y, 77, cellHeight, "D")
	presenterY := y + (cellHeight - float64(len(presenterLines))*rowHeight)/2
	pdf.SetXY(x+30+170, presenterY)
	pdf.MultiCell(77, rowHeight, presenter, "", "C", false)

	pdf.SetY(y + cellHeight)
}


