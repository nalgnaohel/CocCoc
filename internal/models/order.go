package models

type Order struct {
	ID    int    `json:"id"`
	Items []Item `json:"items"`
}

func NewOrder(id int, items []Item) *Order {
	return &Order{
		ID:    id,
		Items: items,
	}
}
