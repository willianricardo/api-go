package entity

type Order struct {
	ID         UniqueID    `json:"id"`
	OrderDate  string      `json:"order_date"`
	CustomerID UniqueID    `json:"customer_id"`
	Customer   Customer    `json:"customer"`
	OrderItems []OrderItem `json:"order_items"`
}
