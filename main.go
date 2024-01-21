package main

import (
	"api/database"
	"api/handler"
	repository "api/repository/postgres"
	"api/routes"
	"api/service"
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func main() {
	port := ":5371"

	db, err := database.InitializeDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	productRepository := repository.NewProductPostgresRepository(db)
	customerRepository := repository.NewCustomerPostgresRepository(db)
	orderRepository := repository.NewOrderPostgresRepository(db)

	productService := service.NewProductService(productRepository)
	customerService := service.NewCustomerService(customerRepository)
	orderService := service.NewOrderService(orderRepository)

	productHandler := handler.NewProductHandler(productService)
	customerHandler := handler.NewCustomerHandler(customerService)
	orderHandler := handler.NewOrderHandler(orderService)

	app := fiber.New(fiber.Config{
		Immutable:   true,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Static("/", "./public")

	routes.SetupProductRoutes(app, productHandler)
	routes.SetupCustomerRoutes(app, customerHandler)
	routes.SetupOrderRoutes(app, orderHandler)

	if err := app.Listen(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
