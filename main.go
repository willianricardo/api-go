package main

import (
	"api/database"
	"api/handler"
	"api/repository"
	"api/routes"
	"api/service"
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func main() {
	port := ":8080"

	// Initialize the database
	db, err := database.InitializeDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
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
	app := fiber.New(fiber.Config{
		Immutable:   true,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Static("/", "./public")

	// Define the API routes
	routes.SetupProductRoutes(app, productHandler)
	routes.SetupCustomerRoutes(app, customerHandler)
	routes.SetupOrderRoutes(app, orderHandler)

	// Start the HTTP server
	if err := app.Listen(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
