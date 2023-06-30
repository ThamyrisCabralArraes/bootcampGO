package models

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"nome"`
	Type  string  `json:"tipo"`
	Count int     `json:"quantidade"`
	Price float64 `json:"preço"`
}
