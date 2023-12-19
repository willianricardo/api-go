package routes

import (
	"api/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupCustomerRoutes(c *fiber.App, customerHandler *handler.CustomerHandler) {
	router := c.Group("/customers")
	router.Get("", customerHandler.GetCustomers)
	router.Get("/:id", customerHandler.GetCustomerByID)
	router.Post("", customerHandler.CreateCustomer)
	router.Put("/:id", customerHandler.UpdateCustomer)
	router.Delete("/:id", customerHandler.DeleteCustomer)
}
