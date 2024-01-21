package entity

import "errors"

type Order struct {
	ID         UniqueID     `json:"id"`
	OrderDate  string       `json:"order_date"`
	CustomerID UniqueID     `json:"customer_id"`
	OrderItems []*OrderItem `json:"order_items"`
}

func NewOrder(id string, customerID string, orderDate string, orderItems []*OrderItem) (*Order, error) {
	orderID := generateOrderID(id)

	order := Order{
		ID:         orderID,
		CustomerID: UniqueIDFromValue(customerID),
		OrderDate:  orderDate,
		OrderItems: orderItems,
	}

	err := order.validate()
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (p *Order) validate() error {
	if p.CustomerID == "" {
		return errors.New("customer id is required")
	}

	if p.OrderDate == "" {
		return errors.New("order date is required")
	}

	if len(p.OrderItems) == 0 {
		return errors.New("order items are required")
	}

	return nil
}

func generateOrderID(id string) UniqueID {
	if id == "" {
		return NewUniqueID()
	}
	return UniqueIDFromValue(id)
}
