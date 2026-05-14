package scheduling

// contains verifica se slice possui valor
func contains(values []string, target string) bool {

	for _, value := range values {

		if value == target {
			return true
		}
	}

	return false
}