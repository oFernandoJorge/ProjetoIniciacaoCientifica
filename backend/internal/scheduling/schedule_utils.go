package scheduling

import "time"

// calculateDuration calcula duração
func calculateDuration(
	presentationType string,
	totalSubmissions int,
) time.Duration {

	var minutesPerPresentation int

	switch presentationType {

	case "ORAL":
		minutesPerPresentation = 20

	case "E-POSTER":
		minutesPerPresentation = 5

	default:
		minutesPerPresentation = 10
	}

	totalMinutes := (
		minutesPerPresentation *
			totalSubmissions)

	return time.Duration(
		totalMinutes,
	) * time.Minute
}

// determineShift define turno
func determineShift(
	t time.Time,
) string {

	hour := t.Hour()

	switch {

	case hour >= 7 && hour < 12:
		return "MANHÃ"

	case hour >= 12 && hour < 18:
		return "TARDE"

	default:
		return "NOITE"
	}
}