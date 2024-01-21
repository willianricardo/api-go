package entity

type Currency struct {
	ID   UniqueID `json:"id"`
	Code string   `json:"code"`
}
