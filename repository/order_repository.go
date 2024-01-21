package repository

import "api/entity"

type OrderRepository interface {
	GetOrders() ([]entity.Order, error)
	GetOrderByID(id string) (entity.Order, error)
	CreateOrder(Order *entity.Order) error
	UpdateOrder(Order *entity.Order) error
	DeleteOrder(id string) error
}
