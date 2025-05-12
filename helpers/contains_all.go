package helpers

func ContainsAll(set map[string]bool, items []string) bool {
	for _, it := range items {
		if !set[it] {
			return false
		}
	}
	return true
}
