package service

import (
	"api/model"
	"api/repository"
)

type OrderService interface {
	GetOrders() ([]model.Order, error)
	GetOrderByID(orderID string) (model.Order, error)
	CreateOrder(order model.Order) error
	UpdateOrder(order model.Order) error
	DeleteOrder(orderID string) error
}

type orderService struct {
	orderRepository *repository.OrderRepository
}

func NewOrderService(orderRepository *repository.OrderRepository) OrderService {
	return &orderService{
		orderRepository: orderRepository,
	}
}

func (service *orderService) GetOrders() ([]model.Order, error) {
	return service.orderRepository.GetOrders()
}

func (service *orderService) GetOrderByID(orderID string) (model.Order, error) {
	return service.orderRepository.GetOrderByID(orderID)
}

func (service *orderService) CreateOrder(order model.Order) error {
	return service.orderRepository.CreateOrder(order)
}

func (service *orderService) UpdateOrder(order model.Order) error {
	return service.orderRepository.UpdateOrder(order)
}

func (service *orderService) DeleteOrder(orderID string) error {
	return service.orderRepository.DeleteOrder(orderID)
}
