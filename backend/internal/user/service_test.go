package user

import "testing"

func TestCreateUserSuccess(t *testing.T) {
	mockRepo := &mockRepository{}
	service := NewService(mockRepo)
	user := &User{
		Name:     "Fernando",
		Email:    "fernando@email.com",
		Password: "123456",
		Role:     "admin",
	}
	err := service.Create(user)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}
func TestCreateUserWithoutName(t *testing.T) {
	mockRepo := &mockRepository{}
	service := NewService(mockRepo)
	user := &User{
		Name:     "",
		Email:    "fernando@email.com",
		Password: "123456",
		Role:     "admin",
	}
	err := service.Create(user)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
