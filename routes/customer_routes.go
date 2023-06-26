package routes

import (
	"api/handler"

	"github.com/gin-gonic/gin"
)

func SetupCustomerRoutes(r *gin.Engine, customerHandler *handler.CustomerHandler) {
	router := r.Group("/customers")
	router.GET("", customerHandler.GetCustomers)
	router.GET("/:id", customerHandler.GetCustomerByID)
	router.POST("", customerHandler.CreateCustomer)
	router.PUT("/:id", customerHandler.UpdateCustomer)
	router.DELETE("/:id", customerHandler.DeleteCustomer)
}
