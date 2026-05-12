package user

type mockRepository struct {}

func (m *mockRepository) Create(user *User) error {
	return nil
}

func (m *mockRepository) FindAll() ([]User, error) {
	return []User{},nil
}

func (m *mockRepository) FindByID(id uint) (*User, error) {
	return &User{}, nil
}