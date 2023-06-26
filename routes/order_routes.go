package routes

import (
	"api/handler"

	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(r *gin.Engine, orderHandler *handler.OrderHandler) {
	router := r.Group("/orders")
	router.GET("", orderHandler.GetOrders)
	router.GET("/:id", orderHandler.GetOrderByID)
	router.POST("", orderHandler.CreateOrder)
	router.PUT("/:id", orderHandler.UpdateOrder)
	router.DELETE("/:id", orderHandler.DeleteOrder)
}
