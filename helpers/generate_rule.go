package helpers

import (
	"algoritma-apriori/model/domain"
	"sort"
	"strings"
)

func GenerateRules(freq map[string]domain.ItemsetSupport, minConfidence float64) []domain.Rule {
	var rules []domain.Rule
	for _, is := range freq {
		items := is.Itemset
		if len(items) < 2 {
			continue
		}
		n := len(items)
		for i := 1; i < (1<<n)-1; i++ {
			var antecedent, consequent []string
			for j := 0; j < n; j++ {
				if i&(1<<j) > 0 {
					antecedent = append(antecedent, items[j])
				} else {
					consequent = append(consequent, items[j])
				}
			}
			sort.Strings(antecedent)
			sort.Strings(consequent)
			aKey := strings.Join(antecedent, ",")
			cKey := strings.Join(consequent, ",")
			confidence := is.Support / freq[aKey].Support
			lift := confidence / freq[cKey].Support
			if confidence >= minConfidence {
				rules = append(rules, domain.Rule{Antecedent: antecedent, Consequent: consequent, Support: is.Support, Confidence: confidence, Lift: lift})
			}
		}
	}
	return rules
}
