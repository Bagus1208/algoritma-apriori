package domain

type PopularItem struct {
	MenuId  string  `firestore:"menu_id"`
	Support float64 `firestore:"support"`
}
