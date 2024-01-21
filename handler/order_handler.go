package handler

import (
	"api/entity"
	"api/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (handler *OrderHandler) GetOrders(c *fiber.Ctx) error {
	orders, err := handler.orderService.GetOrders()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(orders)
}

func (handler *OrderHandler) GetOrderByID(c *fiber.Ctx) error {
	orderID := c.Params("id")
	order, err := handler.orderService.GetOrderByID(orderID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if order.ID == "" {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Order not found"})
	}
	return c.Status(http.StatusOK).JSON(order)
}

type OrderItem struct {
	ID        string  `json:"id"`
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type CreateOrderRequest struct {
	ID         string       `json:"id"`
	CustomerID string       `json:"customer_id"`
	OrderDate  string       `json:"order_date"`
	OrderItems []*OrderItem `json:"order_items"`
}

func (handler *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var request CreateOrderRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid order data"})
	}
	orderItems := make([]*entity.OrderItem, 0)
	for _, item := range request.OrderItems {
		orderItem, err := entity.NewOrderItem(item.ID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		orderItems = append(orderItems, orderItem)
	}
	order, err := entity.NewOrder(request.ID, request.CustomerID, request.OrderDate, orderItems)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := handler.orderService.CreateOrder(order); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Order created"})
}

type UpdateOrderRequest struct {
	CustomerID string       `json:"customer_id"`
	OrderDate  string       `json:"order_date"`
	OrderItems []*OrderItem `json:"order_items"`
}

func (handler *OrderHandler) UpdateOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")
	var request UpdateOrderRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid order data"})
	}
	orderItems := make([]*entity.OrderItem, 0)
	for _, item := range request.OrderItems {
		orderItem, err := entity.NewOrderItem(item.ID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		orderItems = append(orderItems, orderItem)
	}
	order, err := entity.NewOrder(orderID, request.CustomerID, request.OrderDate, orderItems)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := handler.orderService.UpdateOrder(order); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Order updated"})
}

func (handler *OrderHandler) DeleteOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")
	err := handler.orderService.DeleteOrder(orderID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Order deleted"})
}
