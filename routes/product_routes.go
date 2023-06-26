package routes

import (
	"api/handler"

	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(r *gin.Engine, productHandler *handler.ProductHandler) {
	router := r.Group("/products")
	router.GET("", productHandler.GetProducts)
	router.GET("/:id", productHandler.GetProductByID)
	router.POST("", productHandler.CreateProduct)
	router.PUT("/:id", productHandler.UpdateProduct)
	router.DELETE("/:id", productHandler.DeleteProduct)
}
