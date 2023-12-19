package entity

type Customer struct {
	ID   UniqueID `json:"id"`
	Name string   `json:"name"`
}
