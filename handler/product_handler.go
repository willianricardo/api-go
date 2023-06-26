package handler

import (
	"api/model"
	"api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (handler *ProductHandler) GetProducts(c *gin.Context) {
	products, err := handler.productService.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (handler *ProductHandler) GetProductByID(c *gin.Context) {
	productID := c.Param("id")
	product, err := handler.productService.GetProductByID(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (handler *ProductHandler) CreateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
		return
	}

	if err := handler.productService.CreateProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.Status(http.StatusCreated)
}

func (handler *ProductHandler) UpdateProduct(c *gin.Context) {
	productID := c.Param("id")
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
		return
	}
	product.ID = productID

	if err := handler.productService.UpdateProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.Status(http.StatusOK)
}

func (handler *ProductHandler) DeleteProduct(c *gin.Context) {
	productID := c.Param("id")
	if err := handler.productService.DeleteProduct(productID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.Status(http.StatusOK)
}
