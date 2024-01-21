package entity

import "errors"

type Product struct {
	ID    UniqueID `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
}

func NewProduct(id string, name string, price float64) (*Product, error) {
	productID := generateProductID(id)

	product := Product{
		ID:    productID,
		Name:  name,
		Price: price,
	}

	err := product.validate()
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *Product) validate() error {
	if p.Name == "" {
		return errors.New("name is required")
	}

	if p.Price < 0 {
		return errors.New("price must be greater than or equal to 0")
	}

	return nil
}

func generateProductID(id string) UniqueID {
	if id == "" {
		return NewUniqueID()
	}
	return UniqueIDFromValue(id)
}
