package helpers

import "algoritma-apriori/model/domain"

func ExtractItemsets(txs []domain.Transaction) [][]string {
	var sets [][]string
	for _, tx := range txs {
		var ids []string
		for _, item := range tx.Items {
			ids = append(ids, item.MenuId)
		}
		sets = append(sets, ids)
	}
	return sets
}
