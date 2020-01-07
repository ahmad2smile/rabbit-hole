package dedupeList

func RemoveDupe(s []string) []string {
	m := make(map[string]bool)

	for _, item := range s {
		if _, ok := m[item]; !ok {
			m[item] = true
		}
	}

	var result []string

	for key, _ := range m {
		result = append(result, key)
	}

	return result
}
