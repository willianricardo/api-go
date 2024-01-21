package repository

import "api/entity"

type ProductRepository interface {
	GetProducts() ([]entity.Product, error)
	GetProductByID(id string) (entity.Product, error)
	CreateProduct(product *entity.Product) error
	UpdateProduct(product *entity.Product) error
	DeleteProduct(id string) error
}
