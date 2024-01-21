package repository

import "api/entity"

type CustomerRepository interface {
	GetCustomers() ([]entity.Customer, error)
	GetCustomerByID(id string) (entity.Customer, error)
	CreateCustomer(Customer *entity.Customer) error
	UpdateCustomer(Customer *entity.Customer) error
	DeleteCustomer(id string) error
}
