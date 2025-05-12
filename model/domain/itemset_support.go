package domain

type ItemsetSupport struct {
	Itemset []string `firestore:"itemset"`
	Support float64  `firestore:"support"`
}
