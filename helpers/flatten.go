package helpers

func Flatten(list [][]string) []string {
	var out []string
	for _, l := range list {
		out = append(out, l...)
	}
	return out
}
