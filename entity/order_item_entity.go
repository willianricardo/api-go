package entity

import "errors"

type OrderItem struct {
	ID        UniqueID `json:"id"`
	ProductID UniqueID `json:"product_id"`
	Quantity  int      `json:"quantity"`
	Price     float64  `json:"price"`
}

func NewOrderItem(id string, productID string, quantity int, price float64) (*OrderItem, error) {
	orderItemID := generateOrderItemID(id)

	orderItem := OrderItem{
		ID:        orderItemID,
		ProductID: UniqueIDFromValue(productID),
		Quantity:  quantity,
		Price:     price,
	}

	err := orderItem.validate()
	if err != nil {
		return nil, err
	}

	return &orderItem, nil
}

func (p *OrderItem) validate() error {
	if p.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	if p.Price < 0 {
		return errors.New("price must be greater than or equal to 0")
	}

	return nil
}

func generateOrderItemID(id string) UniqueID {
	if id == "" {
		return NewUniqueID()
	}
	return UniqueIDFromValue(id)
}
