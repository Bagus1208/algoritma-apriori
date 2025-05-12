package domain

type Item struct {
	MenuId       string `firestore:"menu_id"`
	MenuName     string `firestore:"menu_name"`
	MenuImageUrl string `firestore:"menu_image_url"`
	Quantity     int    `firestore:"quantity"`
	Subtotal     int    `firestore:"subtotal"`
}

type Transaction struct {
	Id    string `firestore:"id"`
	Items []Item `firestore:"items"`
	Total int    `firestore:"total"`
}
