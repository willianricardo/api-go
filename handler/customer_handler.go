package handler

import (
	"api/model"
	"api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerService *service.CustomerService
}

func NewCustomerHandler(customerService *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

func (handler *CustomerHandler) GetCustomers(c *gin.Context) {
	customers, err := handler.customerService.GetCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve customers"})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (handler *CustomerHandler) GetCustomerByID(c *gin.Context) {
	customerID := c.Param("id")
	customer, err := handler.customerService.GetCustomerByID(customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve customer"})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (handler *CustomerHandler) CreateCustomer(c *gin.Context) {
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer data"})
		return
	}

	if err := handler.customerService.CreateCustomer(customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	c.Status(http.StatusCreated)
}

func (handler *CustomerHandler) UpdateCustomer(c *gin.Context) {
	customerID := c.Param("id")
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer data"})
		return
	}
	customer.ID = customerID

	if err := handler.customerService.UpdateCustomer(customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer"})
		return
	}

	c.Status(http.StatusOK)
}

func (handler *CustomerHandler) DeleteCustomer(c *gin.Context) {
	customerID := c.Param("id")
	if err := handler.customerService.DeleteCustomer(customerID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer"})
		return
	}

	c.Status(http.StatusOK)
}
