package scheduling

import (
	"ProjetoIniciacaoCientifica/internal/submission"
)

// GroupSubmissionsByArea agrupa trabalhos
func GroupSubmissionsByArea(
	submissions []submission.Submission,
) []SubmissionGroup {

	groupsMap := make(map[string]*SubmissionGroup)

	for _, sub := range submissions {

		// Chave única
		key := sub.KnowledgeArea + "_" + sub.PresentationType

		// Cria grupo se não existir
		if _, exists := groupsMap[key]; !exists {

			groupsMap[key] = &SubmissionGroup{
				KnowledgeArea: sub.KnowledgeArea,

				PresentationType: sub.PresentationType,

				Courses: []string{},

				Submissions: []submission.Submission{},
			}
		}

		group := groupsMap[key]

		// Adiciona curso se ainda não existir
		if !contains(group.Courses, sub.Course) {

			// Máximo 3 cursos
			if len(group.Courses) < 3 {
				group.Courses = append(group.Courses, sub.Course)
			}
		}

		group.Submissions = append(
			group.Submissions,
			sub,
		)
	}

	var groups []SubmissionGroup

	for _, group := range groupsMap {
		groups = append(groups, *group)
	}

	return groups
}