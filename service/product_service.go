package service

import (
	"api/entity"
	"api/repository"
)

type ProductService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (service *ProductService) GetProducts() ([]entity.Product, error) {
	return service.productRepository.GetProducts()
}

func (service *ProductService) GetProductByID(id string) (entity.Product, error) {
	return service.productRepository.GetProductByID(id)
}

func (service *ProductService) CreateProduct(product *entity.Product) error {
	return service.productRepository.CreateProduct(product)
}

func (service *ProductService) UpdateProduct(product *entity.Product) error {
	return service.productRepository.UpdateProduct(product)
}

func (service *ProductService) DeleteProduct(id string) error {
	return service.productRepository.DeleteProduct(id)
}
