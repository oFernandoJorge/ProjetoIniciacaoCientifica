package room

import "fmt"

// generateRoomName gera nome da sala
func generateRoomName(floor int, roomNumber int) string {

	return fmt.Sprintf(
		"%d%02d",
		floor,
		roomNumber,
	)
}