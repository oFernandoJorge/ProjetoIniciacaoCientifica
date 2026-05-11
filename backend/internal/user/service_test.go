package user

import "testing"

func TestCreateUser(t *testing.T) {

	mockRepo := &mockRepository{}
	service := NewService(mockRepo)

	user := &User{

		Name: "Fernando",
		Email: "fernando@email.com",
		Password: "123456",
		Role: "admin",
	}

	err := service.Create(user)

	if err != nil {
		t.Errorf("excepted nil, got %v", err)
	}
}

func TestCreateUserWithoutName(t *testing.T) {
	mockRepo := &mockRepository{}
	service := NewService(mockRepo)

	user := &User{
		Name: "",
		Email: "fernando@email.com",
		Password: "123456",
		Role: "admin",
	}

	err := service.Create(user)

	if err == nil || err.Error() != "O nome do usuário é obrigatório" {
		t.Errorf("excepted error 'O nome do usuário é obrigatório', got %v", err)
	}
}

func TestCreateUserInvalidRole(t *testing.T) {
	mockRepo := &mockRepository{}
	service := NewService(mockRepo)

	user := &User{
		Name: "Fernando",
		Email: "fernando@email.com",
		Password: "123456",
		Role: "teste",
	}

	err := service.Create(user)

	if err == nil || err.Error() != "O papel do usuário é inválido" {
		t.Errorf("excepted error 'O papel do usuário é inválido', got %v", err)
	}
}

func TestCreateUserWithoutPassword(t *testing.T) {
	mockRepo := &mockRepository{}
	service := NewService(mockRepo)	

	user := &User{
		Name: "Fernando",
		Email: "fernando@email.com",
		Password: "",
		Role: "admin",
	}

	err := service.Create(user)

	if err == nil || err.Error() != "A senha do usuário é obrigatória" {
		t.Errorf("excepted error 'A senha do usuário é obrigatória', got %v", err)
	}
}