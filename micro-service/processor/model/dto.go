package model

type Category int16

const (
	Grocery Category = iota
	Food
	Medicine
	Hardware
	Toy
)

type Store struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Address  string   `json:"address"`
	Tin  string   `json:"tin"`
	PhoneNo  string   `json:"phone_no"`
	Category Category `json:"category"`
}

type Item struct {
	ItemID   string  `json:"item_id"`
	Name     string  `json:"name"`
	Quantity string   `json:"quantity"`
	Price    string `json:"price"`
	StoreID  string  `json:"store_id"`
}
