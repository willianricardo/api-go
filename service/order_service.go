package service

import (
	"api/entity"
	"api/repository"
)

type OrderService interface {
	GetOrders() ([]entity.Order, error)
	GetOrderByID(orderID string) (entity.Order, error)
	CreateOrder(order *entity.Order) error
	UpdateOrder(order *entity.Order) error
	DeleteOrder(orderID string) error
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &orderService{
		orderRepository: orderRepository,
	}
}

func (service *orderService) GetOrders() ([]entity.Order, error) {
	return service.orderRepository.GetOrders()
}

func (service *orderService) GetOrderByID(orderID string) (entity.Order, error) {
	return service.orderRepository.GetOrderByID(orderID)
}

func (service *orderService) CreateOrder(order *entity.Order) error {
	return service.orderRepository.CreateOrder(order)
}

func (service *orderService) UpdateOrder(order *entity.Order) error {
	return service.orderRepository.UpdateOrder(order)
}

func (service *orderService) DeleteOrder(orderID string) error {
	return service.orderRepository.DeleteOrder(orderID)
}
