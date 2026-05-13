package scheduling

import (
	"ProjetoIniciacaoCientifica/internal/room"
)

// AllocateRooms distribui trabalhos
func AllocateRooms(
	groups []SubmissionGroup,
	rooms []room.Room,
) []RoomAllocation {

	var allocations []RoomAllocation

	roomIndex := 0

	for _, group := range groups {

		var compatibleRooms []room.Room

		// Filtra salas compatíveis
		for _, room := range rooms {

			if room.PresentationType == group.PresentationType {

				compatibleRooms = append(
					compatibleRooms,
					room,
				)
			}
		}

		if len(compatibleRooms) == 0 {
			continue
		}

		selectedRoom := compatibleRooms[roomIndex%len(compatibleRooms)]

		capacity := selectedRoom.Capacity

		for i := 0; i < len(group.Submissions); i += capacity {

			end := i + capacity

			if end > len(group.Submissions) {
				end = len(group.Submissions)
			}

			allocation := RoomAllocation{
				RoomID: selectedRoom.ID,

				RoomName: selectedRoom.Name,

				KnowledgeArea: group.KnowledgeArea,

				PresentationType: group.PresentationType,

				Submissions: group.Submissions[i:end],
			}

			allocations = append(
				allocations,
				allocation,
			)

			roomIndex++
		}
	}

	return allocations
}
