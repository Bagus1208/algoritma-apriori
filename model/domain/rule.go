package domain

type Rule struct {
	Antecedent []string `firestore:"antecedent"`
	Consequent []string `firestore:"consequent"`
	Support    float64  `firestore:"support"`
	Confidence float64  `firestore:"confidence"`
	Lift       float64  `firestore:"lift"`
}
