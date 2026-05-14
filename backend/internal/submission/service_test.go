package submission

import "testing"

// TestCreateSubmissionSuccess testa criação válida
func TestCreateSubmissionSuccess(t *testing.T) {

	mockRepo := &mockRepository{}

	service := NewService(mockRepo)

	submission := &Submission{
		Title:            "Sistema Inteligente",
		PresenterName:    "Fernando",
		Course:           "ADS",
		KnowledgeArea:    "Tecnologia",
		Modality:         "E-POSTER",
		Campus:           "Fortaleza",
		AdvisorName:      "Professor X",
		PresentationType: "Oral",
	}

	err := service.Create(submission)

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

// TestCreateSubmissionWithoutTitle testa título obrigatório
func TestCreateSubmissionWithoutTitle(t *testing.T) {

	mockRepo := &mockRepository{}

	service := NewService(mockRepo)

	submission := &Submission{
		Title:            "",
		PresenterName:    "Fernando",
		Course:           "ADS",
		KnowledgeArea:    "Tecnologia",
		Modality:         "E-POSTER",
		Campus:           "Fortaleza",
		AdvisorName:      "Professor X",
		PresentationType: "Oral",
	}

	err := service.Create(submission)

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

// TestCreateSubmissionWithoutPresenter testa apresentador obrigatório
func TestCreateSubmissionWithoutPresenter(t *testing.T) {

	mockRepo := &mockRepository{}

	service := NewService(mockRepo)

	submission := &Submission{
		Title:            "Sistema Inteligente",
		PresenterName:    "",
		Course:           "ADS",
		KnowledgeArea:    "Tecnologia",
		Modality:         "E-POSTER",
		Campus:           "Fortaleza",
		AdvisorName:      "Professor X",
		PresentationType: "Oral",
	}

	err := service.Create(submission)

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

// TestCreateSubmissionWithoutCourse testa curso obrigatório
func TestCreateSubmissionWithoutCourse(t *testing.T) {

	mockRepo := &mockRepository{}

	service := NewService(mockRepo)

	submission := &Submission{
		Title:            "Sistema Inteligente",
		PresenterName:    "Fernando",
		Course:           "",
		KnowledgeArea:    "Tecnologia",
		Modality:         "E-POSTER",
		Campus:           "Fortaleza",
		AdvisorName:      "Professor X",
		PresentationType: "Oral",
	}

	err := service.Create(submission)

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}