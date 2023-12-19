package entity

type Product struct {
	ID    UniqueID `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
}
