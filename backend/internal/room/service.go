package room

// Service contém regras de negócio
type Service struct {
	repo Repository
}

// NewService injeta repository
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// GenerateDefaultRooms gera automaticamente
// todas as salas do evento
func (s *Service) GenerateDefaultRooms() error {

	// Gera andares do 2° ao 5°
	for floor := 2; floor <= 5; floor++ {

		// E-POSTER
		for roomNumber := 1; roomNumber <= 10; roomNumber++ {

			room := Room{
				Name:             generateRoomName(floor, roomNumber),
				Floor:            floor,
				PresentationType: "E-POSTER",
				Capacity:         12,
				IsAvailable:      true,
			}

			if err := s.repo.Create(&room); err != nil {
				return err
			}
		}

		// APRESENTAÇÃO ORAL
		for roomNumber := 11; roomNumber <= 20; roomNumber++ {

			room := Room{
				Name:             generateRoomName(floor, roomNumber),
				Floor:            floor,
				PresentationType: "ORAL",
				Capacity:         6,
				IsAvailable:      true,
			}

			if err := s.repo.Create(&room); err != nil {
				return err
			}
		}
	}

	return nil
}