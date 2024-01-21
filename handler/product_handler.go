package handler

import (
	"api/entity"
	"api/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (handler *ProductHandler) GetProducts(c *fiber.Ctx) error {
	products, err := handler.productService.GetProducts()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve products"})
	}
	return c.Status(http.StatusOK).JSON(products)
}

func (handler *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	productID := c.Params("id")
	product, err := handler.productService.GetProductByID(productID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve product"})
	}
	if product.ID == "" {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.Status(http.StatusOK).JSON(product)
}

type CreateProductRequest struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (handler *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var request CreateProductRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product data"})
	}
	product, err := entity.NewProduct(request.ID, request.Name, request.Price)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := handler.productService.CreateProduct(product); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create product"})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Product created"})
}

type UpdateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (handler *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	productID := c.Params("id")

	var request UpdateProductRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product data"})
	}
	product, err := entity.NewProduct(productID, request.Name, request.Price)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := handler.productService.UpdateProduct(product); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update product"})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Product updated"})
}

func (handler *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	productID := c.Params("id")
	if err := handler.productService.DeleteProduct(productID); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete product"})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Product deleted"})
}
