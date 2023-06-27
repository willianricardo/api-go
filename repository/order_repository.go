package repository

import (
	"api/model"
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

func (repository *OrderRepository) GetOrders() ([]model.Order, error) {
	ordersMap := make(map[string]model.Order)
	orderItemsMap := make(map[string][]model.OrderItem)

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
		order := model.Order{}
		orderItem := model.OrderItem{}
		customer := model.Customer{}
		product := model.Product{}

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

		if _, ok := ordersMap[order.ID]; !ok {
			ordersMap[order.ID] = order
		}
		orderItemsMap[order.ID] = append(orderItemsMap[order.ID], orderItem)
	}

	orders := make([]model.Order, 0, len(ordersMap))

	for orderID, order := range ordersMap {
		order.OrderItems = orderItemsMap[orderID]
		orders = append(orders, order)
	}

	return orders, nil
}

func (repository *OrderRepository) GetOrderByID(orderID string) (model.Order, error) {
	var order model.Order

	query := `
		SELECT o.id, o.customer_id, o.order_date,
			   c.id, c.name,
			   oi.id, oi.order_id, oi.product_id, oi.quantity, oi.price,
			   p.id, p.name, p.price
		FROM orders o
		INNER JOIN customers c ON o.customer_id = c.id
		INNER JOIN order_items oi ON oi.order_id = o.id
		INNER JOIN products p ON oi.product_id = p.id
		WHERE o.id = ?
	`

	rows, err := repository.db.Query(query, orderID)
	if err != nil {
		return order, err
	}
	defer rows.Close()

	orderItems := make([]model.OrderItem, 0)
	var customer model.Customer
	foundRows := false

	for rows.Next() {
		orderItem := model.OrderItem{}
		product := model.Product{}

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

func (repository *OrderRepository) CreateOrder(order model.Order) error {
	tx, err := repository.db.Begin()
	if err != nil {
		return err
	}

	// Insert order
	_, err = tx.Exec("INSERT INTO orders (id, customer_id, order_date) VALUES (?, ?, ?)", order.ID, order.CustomerID, order.OrderDate)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert order items
	for _, orderItem := range order.OrderItems {
		_, err = tx.Exec("INSERT INTO order_items (id, order_id, product_id, quantity, price) VALUES (?, ?, ?, ?, ?)", orderItem.ID, order.ID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
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

func (repository *OrderRepository) UpdateOrder(order model.Order) error {
	tx, err := repository.db.Begin()
	if err != nil {
		return err
	}

	// Update order
	_, err = tx.Exec("UPDATE orders SET customer_id = ?, order_date = ? WHERE id = ?", order.CustomerID, order.OrderDate, order.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Delete existing order items
	_, err = tx.Exec("DELETE FROM order_items WHERE order_id = ?", order.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert updated order items
	for _, orderItem := range order.OrderItems {
		_, err = tx.Exec("INSERT INTO order_items (id, order_id, product_id, quantity, price) VALUES (?, ?, ?, ?, ?)", orderItem.ID, order.ID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
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

	// Delete order items
	_, err = tx.Exec("DELETE FROM order_items WHERE order_id = ?", orderID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Delete order
	_, err = tx.Exec("DELETE FROM orders WHERE id = ?", orderID)
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
