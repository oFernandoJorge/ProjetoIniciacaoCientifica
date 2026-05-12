package submission

import "errors"

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

// Create cria submissão
func (s *Service) Create(submission *Submission) error {

	if submission.Title == "" {
		return errors.New("título obrigatório")
	}

	if submission.PresenterName == "" {
		return errors.New("apresentador obrigatório")
	}

	if submission.Course == "" {
		return errors.New("curso obrigatório")
	}

	return s.repo.Create(submission)
}

// FindAll retorna submissões
func (s *Service) FindAll() ([]SubmissionResponse, error) {

	submissions, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	var response []SubmissionResponse

	for _, s := range submissions {

		response = append(response, SubmissionResponse{
			ID:               s.ID,
			Title:            s.Title,
			PresenterName:    s.PresenterName,
			Course:           s.Course,
			KnowledgeArea:    s.KnowledgeArea,
			Modality:         s.Modality,
			Campus:           s.Campus,
			AdvisorName:      s.AdvisorName,
			PresentationType: s.PresentationType,
		})
	}

	return response, nil
}