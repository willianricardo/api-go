package service

import (
	"api/model"
	"api/repository"
)

type ProductService struct {
	productRepository *repository.ProductRepository
}

func NewProductService(productRepository *repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (service *ProductService) GetProducts() ([]model.Product, error) {
	return service.productRepository.GetProducts()
}

func (service *ProductService) GetProductByID(id string) (model.Product, error) {
	return service.productRepository.GetProductByID(id)
}

func (service *ProductService) CreateProduct(product model.Product) error {
	return service.productRepository.CreateProduct(product)
}

func (service *ProductService) UpdateProduct(product model.Product) error {
	return service.productRepository.UpdateProduct(product)
}

func (service *ProductService) DeleteProduct(id string) error {
	return service.productRepository.DeleteProduct(id)
}
