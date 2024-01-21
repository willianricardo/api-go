package entity

type Money struct {
	Value    float64  `json:"value"`
	Currency Currency `json:"currency"`
}
