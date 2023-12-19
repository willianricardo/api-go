package entity

type OrderItem struct {
	ID        UniqueID `json:"id"`
	OrderID   UniqueID `json:"order_id"`
	ProductID UniqueID `json:"product_id"`
	Product   Product  `json:"product"`
	Quantity  int      `json:"quantity"`
	Price     float64  `json:"price"`
}
