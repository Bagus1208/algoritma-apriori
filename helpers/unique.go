package helpers

import "sort"

func Unique(items []string) []string {
	m := make(map[string]bool)
	for _, it := range items {
		m[it] = true
	}
	var uniq []string
	for it := range m {
		uniq = append(uniq, it)
	}
	sort.Strings(uniq)
	return uniq
}
