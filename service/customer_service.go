package service

import (
	"api/entity"
	"api/repository"
)

type CustomerService struct {
	customerRepository *repository.CustomerRepository
}

func NewCustomerService(customerRepository *repository.CustomerRepository) *CustomerService {
	return &CustomerService{
		customerRepository: customerRepository,
	}
}

func (service *CustomerService) GetCustomers() ([]entity.Customer, error) {
	return service.customerRepository.GetCustomers()
}

func (service *CustomerService) GetCustomerByID(id string) (entity.Customer, error) {
	return service.customerRepository.GetCustomerByID(id)
}

func (service *CustomerService) CreateCustomer(customer entity.Customer) error {
	return service.customerRepository.CreateCustomer(customer)
}

func (service *CustomerService) UpdateCustomer(customer entity.Customer) error {
	return service.customerRepository.UpdateCustomer(customer)
}

func (service *CustomerService) DeleteCustomer(id string) error {
	return service.customerRepository.DeleteCustomer(id)
}
