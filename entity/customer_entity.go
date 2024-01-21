package entity

import "errors"

type Customer struct {
	ID   UniqueID `json:"id"`
	Name string   `json:"name"`
}

func NewCustomer(id string, name string) (*Customer, error) {
	customerID := generateCustomerID(id)

	customer := Customer{
		ID:   customerID,
		Name: name,
	}

	err := customer.validate()
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (p *Customer) validate() error {
	if p.Name == "" {
		return errors.New("name is required")
	}

	return nil
}

func generateCustomerID(id string) UniqueID {
	if id == "" {
		return NewUniqueID()
	}
	return UniqueIDFromValue(id)
}
