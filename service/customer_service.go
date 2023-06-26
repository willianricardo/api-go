package service

import (
	"api/model"
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

func (service *CustomerService) GetCustomers() ([]model.Customer, error) {
	return service.customerRepository.GetCustomers()
}

func (service *CustomerService) GetCustomerByID(id string) (model.Customer, error) {
	return service.customerRepository.GetCustomerByID(id)
}

func (service *CustomerService) CreateCustomer(customer model.Customer) error {
	return service.customerRepository.CreateCustomer(customer)
}

func (service *CustomerService) UpdateCustomer(customer model.Customer) error {
	return service.customerRepository.UpdateCustomer(customer)
}

func (service *CustomerService) DeleteCustomer(id string) error {
	return service.customerRepository.DeleteCustomer(id)
}
