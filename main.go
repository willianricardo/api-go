package main

import (
	"api/database"
	"api/handler"
	"api/repository"
	"api/routes"
	"api/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db, err := database.InitializeDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the repository instances
	productRepository := repository.NewProductRepository(db)
	customerRepository := repository.NewCustomerRepository(db)
	orderRepository := repository.NewOrderRepository(db)

	// Create the service instances
	productService := service.NewProductService(productRepository)
	customerService := service.NewCustomerService(customerRepository)
	orderService := service.NewOrderService(orderRepository)

	// Create the handler instances
	productHandler := handler.NewProductHandler(productService)
	customerHandler := handler.NewCustomerHandler(customerService)
	orderHandler := handler.NewOrderHandler(orderService)

	// Initialize the Gin routes
	r := gin.Default()

	// Define the API routes
	routes.SetupProductRoutes(r, productHandler)
	routes.SetupCustomerRoutes(r, customerHandler)
	routes.SetupOrderRoutes(r, orderHandler)

	// Start the HTTP server
	r.Run(":8080")
}
