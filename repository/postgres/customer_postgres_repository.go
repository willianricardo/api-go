package repository

import (
	"api/entity"
	"database/sql"
)

type CustomerPostgresRepository struct {
	db *sql.DB
}

func NewCustomerPostgresRepository(db *sql.DB) *CustomerPostgresRepository {
	return &CustomerPostgresRepository{
		db: db,
	}
}

func (repository *CustomerPostgresRepository) GetCustomers() ([]entity.Customer, error) {
	customers := make([]entity.Customer, 0)
	rows, err := repository.db.Query("SELECT id, name FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer entity.Customer
		err := rows.Scan(&customer.ID, &customer.Name)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (repository *CustomerPostgresRepository) GetCustomerByID(id string) (entity.Customer, error) {
	rows, err := repository.db.Query("SELECT id, name FROM customers WHERE id = $1", id)
	if err != nil {
		return entity.Customer{}, err
	}

	if rows.Next() {
		var customer entity.Customer
		err := rows.Scan(&customer.ID, &customer.Name)
		if err != nil {
			return entity.Customer{}, err
		}
		return customer, nil
	}

	return entity.Customer{}, nil
}

func (repository *CustomerPostgresRepository) CreateCustomer(customer *entity.Customer) error {
	_, err := repository.db.Exec("INSERT INTO customers (id, name) VALUES ($1, $2)", customer.ID, customer.Name)
	if err != nil {
		return err
	}
	return nil
}

func (repository *CustomerPostgresRepository) UpdateCustomer(customer *entity.Customer) error {
	_, err := repository.db.Exec("UPDATE customers SET name = $1 WHERE id = $2", customer.Name, customer.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repository *CustomerPostgresRepository) DeleteCustomer(id string) error {
	_, err := repository.db.Exec("DELETE FROM customers WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
