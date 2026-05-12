package submission

// mockRepository simula repository em testes unitários
type mockRepository struct{}

// Create simula criação de submissão
func (m *mockRepository) Create(submission *Submission) error {
	return nil
}

// FindAll simula listagem de submissões
func (m *mockRepository) FindAll() ([]Submission, error) {
	return []Submission{}, nil
}