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
	return c.Status(http.StatusOK).JSON(order)
}

func (handler *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order entity.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := handler.orderService.CreateOrder(order)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Order created"})
}

func (handler *OrderHandler) UpdateOrder(c *fiber.Ctx) error {
	orderID := entity.UniqueIDFromValue(c.Params("id"))
	var order entity.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	order.ID = orderID
	err := handler.orderService.UpdateOrder(order)
	if err != nil {
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
