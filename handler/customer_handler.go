package handler

import (
	"api/entity"
	"api/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	customerService *service.CustomerService
}

func NewCustomerHandler(customerService *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

func (handler *CustomerHandler) GetCustomers(c *fiber.Ctx) error {
	customers, err := handler.customerService.GetCustomers()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve customers"})
	}
	return c.Status(http.StatusOK).JSON(customers)
}

func (handler *CustomerHandler) GetCustomerByID(c *fiber.Ctx) error {
	customerID := c.Params("id")
	customer, err := handler.customerService.GetCustomerByID(customerID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve customer"})
	}
	if customer.ID == "" {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Customer not found"})
	}
	return c.Status(http.StatusOK).JSON(customer)
}

type CreateCustomerRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (handler *CustomerHandler) CreateCustomer(c *fiber.Ctx) error {
	var request CreateCustomerRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid customer data"})
	}
	customer, err := entity.NewCustomer(request.ID, request.Name)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := handler.customerService.CreateCustomer(customer); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create customer"})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Customer created"})
}

type UpdateCustomerRequest struct {
	Name string `json:"name"`
}

func (handler *CustomerHandler) UpdateCustomer(c *fiber.Ctx) error {
	customerID := c.Params("id")

	var request UpdateCustomerRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid customer data"})
	}
	customer, err := entity.NewCustomer(customerID, request.Name)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := handler.customerService.UpdateCustomer(customer); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update customer"})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Customer updated"})
}

func (handler *CustomerHandler) DeleteCustomer(c *fiber.Ctx) error {
	customerID := c.Params("id")
	if err := handler.customerService.DeleteCustomer(customerID); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete customer"})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Customer deleted"})
}
