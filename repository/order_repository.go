package repository

import (
	"api/entity"
	"database/sql"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (repository *OrderRepository) GetOrders() ([]entity.Order, error) {
	ordersMap := make(map[string]entity.Order)
	orderItemsMap := make(map[string][]entity.OrderItem)

	rows, err := repository.db.Query(`
		SELECT o.id, o.customer_id, o.order_date,
		       c.id, c.name,
		       oi.id, oi.product_id, oi.quantity, oi.price,
		       p.id, p.name, p.price
		FROM orders o
		INNER JOIN customers c ON o.customer_id = c.id
		INNER JOIN order_items oi ON oi.order_id = o.id
		INNER JOIN products p ON oi.product_id = p.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		order := entity.Order{}
		orderItem := entity.OrderItem{}
		customer := entity.Customer{}
		product := entity.Product{}

		err := rows.Scan(
			&order.ID, &order.CustomerID, &order.OrderDate,
			&customer.ID, &customer.Name,
			&orderItem.ID, &orderItem.ProductID, &orderItem.Quantity, &orderItem.Price,
			&product.ID, &product.Name, &product.Price,
		)
		if err != nil {
			return nil, err
		}

		order.Customer = customer
		orderItem.Product = product

		if _, ok := ordersMap[string(order.ID)]; !ok {
			ordersMap[string(order.ID)] = order
		}
		orderItemsMap[string(order.ID)] = append(orderItemsMap[string(order.ID)], orderItem)
	}

	orders := make([]entity.Order, 0, len(ordersMap))

	for orderID, order := range ordersMap {
		order.OrderItems = orderItemsMap[orderID]
		orders = append(orders, order)
	}

	return orders, nil
}

func (repository *OrderRepository) GetOrderByID(orderID string) (entity.Order, error) {
	var order entity.Order

	query := `
		SELECT o.id, o.customer_id, o.order_date,
			   c.id, c.name,
			   oi.id, oi.order_id, oi.product_id, oi.quantity, oi.price,
			   p.id, p.name, p.price
		FROM orders o
		INNER JOIN customers c ON o.customer_id = c.id
		INNER JOIN order_items oi ON oi.order_id = o.id
		INNER JOIN products p ON oi.product_id = p.id
		WHERE o.id = $1
	`

	rows, err := repository.db.Query(query, orderID)
	if err != nil {
		return order, err
	}
	defer rows.Close()

	orderItems := make([]entity.OrderItem, 0)
	var customer entity.Customer
	foundRows := false

	for rows.Next() {
		orderItem := entity.OrderItem{}
		product := entity.Product{}

		err := rows.Scan(
			&order.ID, &order.CustomerID, &order.OrderDate,
			&customer.ID, &customer.Name,
			&orderItem.ID, &orderItem.OrderID, &orderItem.ProductID, &orderItem.Quantity, &orderItem.Price,
			&product.ID, &product.Name, &product.Price,
		)
		if err != nil {
			return order, err
		}

		order.Customer = customer
		orderItem.Product = product
		orderItems = append(orderItems, orderItem)
		foundRows = true
	}

	if !foundRows {
		return order, sql.ErrNoRows
	}

	order.OrderItems = orderItems

	return order, nil
}

func (repository *OrderRepository) CreateOrder(order entity.Order) error {
	tx, err := repository.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO orders (id, customer_id, order_date) VALUES ($1, $2, $3)", order.ID, order.CustomerID, order.OrderDate)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, orderItem := range order.OrderItems {
		_, err = tx.Exec("INSERT INTO order_items (id, order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4, $5)", orderItem.ID, order.ID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (repository *OrderRepository) UpdateOrder(order entity.Order) error {
	tx, err := repository.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE orders SET customer_id = $1, order_date = $2 WHERE id = $3", order.CustomerID, order.OrderDate, order.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM order_items WHERE order_id = $1", order.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, orderItem := range order.OrderItems {
		_, err = tx.Exec("INSERT INTO order_items (id, order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4, $5)", orderItem.ID, order.ID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (repository *OrderRepository) DeleteOrder(orderID string) error {
	tx, err := repository.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM order_items WHERE order_id = $1", orderID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM orders WHERE id = $1", orderID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
