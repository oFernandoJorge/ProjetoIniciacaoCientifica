package scheduling

import "time"

// GenerateSchedules gera horários automáticos
func GenerateSchedules(
	allocations []EvaluatorAllocation,
) []ScheduledSession {

	var sessions []ScheduledSession

	baseDate := time.Date(
		2026,
		6,
		10,
		8,
		0,
		0,
		0,
		time.Local,
	)

	currentTime := baseDate

	for _, allocation := range allocations {

		duration := calculateDuration(
			allocation.RoomAllocation.PresentationType,
			len(allocation.RoomAllocation.Submissions),
		)

		endTime := currentTime.Add(duration)

		shift := determineShift(currentTime)

		session := ScheduledSession{
			EvaluatorAllocation: allocation,

			StartTime: currentTime,

			EndTime: endTime,

			Shift: shift,
		}

		sessions = append(
			sessions,
			session,
		)

		// Próxima sessão começa 10 min depois
		currentTime = endTime.Add(
			10 * time.Minute,
		)
	}

	return sessions
}
