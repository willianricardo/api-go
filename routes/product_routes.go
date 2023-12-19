package routes

import (
	"api/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(c *fiber.App, productHandler *handler.ProductHandler) {
	router := c.Group("/products")
	router.Get("", productHandler.GetProducts)
	router.Get("/:id", productHandler.GetProductByID)
	router.Post("", productHandler.CreateProduct)
	router.Put("/:id", productHandler.UpdateProduct)
	router.Delete("/:id", productHandler.DeleteProduct)
}
