package helpers

import (
	"algoritma-apriori/model/domain"
	"sort"
	"strings"
)

func Apriori(itemsets [][]string, minSupport float64) map[string]domain.ItemsetSupport {
	freq := make(map[string]domain.ItemsetSupport)
	totalTx := float64(len(itemsets))

	// 1-itemset
	count1 := make(map[string]int)
	for _, items := range itemsets {
		for _, id := range items {
			count1[id]++
		}
	}
	var L [][]string
	for id, cnt := range count1 {
		support := float64(cnt) / totalTx
		if support >= minSupport {
			freq[id] = domain.ItemsetSupport{Itemset: []string{id}, Support: support}
			L = append(L, []string{id})
		}
	}

	k := 2
	for len(L) > 0 {
		// generate candidates Ck
		uniqueItems := Unique(Flatten(L))
		var candidates [][]string
		for _, comb := range Combinations(uniqueItems, k) {
			sort.Strings(comb)
			candidates = append(candidates, comb)
		}
		// count support
		countC := make(map[string]int)
		for _, items := range itemsets {
			txSet := make(map[string]bool)
			for _, itm := range items {
				txSet[itm] = true
			}
			for _, cand := range candidates {
				if ContainsAll(txSet, cand) {
					key := strings.Join(cand, ",")
					countC[key]++
				}
			}
		}
		// filter by minSupport
		L = nil
		for key, cnt := range countC {
			support := float64(cnt) / totalTx
			if support >= minSupport {
				items := strings.Split(key, ",")
				freq[key] = domain.ItemsetSupport{Itemset: items, Support: support}
				L = append(L, items)
			}
		}
		k++
	}
	return freq
}
