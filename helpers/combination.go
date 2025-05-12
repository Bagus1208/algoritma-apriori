package helpers

func Combinations(items []string, k int) [][]string {
	var res [][]string
	var comb func(start int, curr []string)
	comb = func(start int, curr []string) {
		if len(curr) == k {
			tmp := make([]string, k)
			copy(tmp, curr)
			res = append(res, tmp)
			return
		}
		for i := start; i < len(items); i++ {
			comb(i+1, append(curr, items[i]))
		}
	}
	comb(0, []string{})
	return res
}
